/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.simulator.server.bacnet.protocol;

import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.io.IOException;

import org.apache.plc4x.java.bacnetip.readwrite.*;
import org.apache.plc4x.java.bacnetip.readwrite.utils.StaticHelper;
import org.apache.plc4x.simulator.model.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public class BacnetServerAdapter extends ChannelInboundHandlerAdapter {

    private static final Logger LOGGER = LoggerFactory.getLogger(BacnetServerAdapter.class);


    // : make configurable
    private static int deviceInstance = 4711;

    // : make configurable
    

    public BacnetServerAdapter(Context context) {
        LOGGER.info("Creating adapter with context {}", context);
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        if(LOGGER.isInfoEnabled()){
         LOGGER.info("Got request");
         LOGGER.info(msg.toString());
        }
        if (!(msg instanceof BVLC)) {
            return;
        }
        BVLC bvlc = (BVLC) msg;
        
        BVLCOriginalUnicastNPDU bvlcOriginalUnicastNPDU = (BVLCOriginalUnicastNPDU) bvlc;
        // : get messageTypeTag
        APDU apdu = bvlcOriginalUnicastNPDU.getNpdu().getApdu();
        if (apdu instanceof APDUUnconfirmedRequest) {
            APDUUnconfirmedRequest apduUnconfirmedRequest = (APDUUnconfirmedRequest) apdu;
            BACnetUnconfirmedServiceRequest serviceRequest = apduUnconfirmedRequest.getServiceRequest();
            if (serviceRequest instanceof BACnetUnconfirmedServiceRequestWhoIs) {
                
                BVLCOriginalUnicastNPDU response = new BVLCOriginalUnicastNPDU(
                    new NPDU(
                        (short) 1,
                        new NPDUControl(false, false, false, false, NPDUNetworkPriority.NORMAL_MESSAGE),
                        null,
                        null,
                        null,
                        null,
                        null,
                        null,
                        null,
                        null,
                        new APDUUnconfirmedRequest(
                            new BACnetUnconfirmedServiceRequestIAm(
                                StaticHelper.createBACnetApplicationTagObjectIdentifier(BACnetObjectType.DEVICE.getValue(), deviceInstance),
                                StaticHelper.createBACnetApplicationTagUnsignedInteger(1024),
                                StaticHelper.creatBACnetSegmentationTagged(BACnetSegmentation.NO_SEGMENTATION),
                                StaticHelper.createBACnetVendorIdApplicationTagged(BACnetVendorId.MAPPED.getVendorId()),
                                0
                            ),
                            0
                        ),
                        0
                    ),
                    0
                );
                
                ctx.writeAndFlush(response).addListener((ChannelFutureListener) f -> {
                    if (!f.isSuccess()) {
                        String format = "An error occurre!" + f;
                        LOGGER.info(format);
                    }
                });
            } else {
                throw new IOException(apdu.getClass() + " not set supported");
            }
        }
    }

}
