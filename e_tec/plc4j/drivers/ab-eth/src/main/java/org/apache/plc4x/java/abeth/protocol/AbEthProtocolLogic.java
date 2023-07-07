/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.abeth.protocol;

import org.apache.plc4x.java.abeth.configuration.AbEthConfiguration;
import org.apache.plc4x.java.abeth.tag.AbEthTag;
import org.apache.plc4x.java.abeth.readwrite.*;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcResponse;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.messages.DefaultPlcReadResponse;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.apache.plc4x.java.spi.values.PlcINT;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicInteger;

public class AbEthProtocolLogic extends Plc4xProtocolBase<CIPEncapsulationPacket> implements HasConfiguration<AbEthConfiguration> {

    private static final Logger logger = LoggerFactory.getLogger(AbEthProtocolLogic.class);
    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);

    private static final List<Short> emptySenderContext = Arrays.asList((short) 0x00 ,(short) 0x00 ,(short) 0x00,
        (short) 0x00,(short) 0x00,(short) 0x00, (short) 0x00,(short) 0x00);

    private AbEthConfiguration configuration;

    private final AtomicInteger transactionCounterGenerator = new AtomicInteger(10);
    private RequestTransactionManager tm;
    private long sessionHandle;

    @Override
    public void setConfiguration(AbEthConfiguration configuration) {
        this.configuration = configuration;
        // Set the transaction manager to allow only one message at a time.
        this.tm = new RequestTransactionManager(1);
    }

    @Override
    public void onConnect(ConversationContext<CIPEncapsulationPacket> context) {
        logger.debug("Sending COTP Connection Request");
        CIPEncapsulationConnectionRequest connectionRequest =
            new CIPEncapsulationConnectionRequest(0L, 0L, emptySenderContext, 0L);
        context.sendRequest(connectionRequest)
            .expectResponse(CIPEncapsulationPacket.class, REQUEST_TIMEOUT)
            .check(CIPEncapsulationConnectionResponse.class:: isInstance)
            .unwrap(CIPEncapsulationConnectionResponse.class::cast)
            .handle(cipEncapsulationConnectionResponse -> {
                sessionHandle = cipEncapsulationConnectionResponse.getSessionHandle();
                // Send an event that connection setup is complete.
                context.fireConnected();
            });
    }

    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        //Warning ... we are senging one request per tag ... the result has to be merged back together ...
        for (String tagName : readRequest.getTagNames()) {
            PlcTag tag = readRequest.getTag(tagName);
            if (!(tag instanceof AbEthTag)) {
                logger.error("The tag should have been of type AbEthTag");
            }
            AbEthTag abEthTag = null;
            if(tag instanceof AbEthTag){
                abEthTag = (AbEthTag) tag;
            }
            
            
            final int transactionCounter = transactionCounterGenerator.incrementAndGet();
            // If we've reached the max value for a 16 bit transaction identifier, reset back to 1
            if(transactionCounterGenerator.get() == 0xFFFF) {
                transactionCounterGenerator.set(1);
            }
            if (abEthTag != null) {
                DF1RequestProtectedTypedLogicalRead logicalRead = new DF1RequestProtectedTypedLogicalRead(
                        abEthTag.getByteSize(), abEthTag.getFileNumber(), abEthTag.getFileType().getTypeCode(),
                        abEthTag.getElementNumber(), (short) 0);
                // Use the logicalRead object
                DF1RequestMessage requestMessage = new DF1CommandRequestMessage(
                (short) configuration.getStation(), (short) 5, (short) 0,
                transactionCounter, logicalRead);
            CIPEncapsulationReadRequest read = new CIPEncapsulationReadRequest(
                sessionHandle, 0, emptySenderContext, 0, requestMessage);

            
// origin/sender: constant = 5
            CompletableFuture<PlcReadResponse> future = new CompletableFuture<>();
            RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
            transaction.submit(() -> context.sendRequest(read)
                .expectResponse(CIPEncapsulationPacket.class, REQUEST_TIMEOUT)
                .onTimeout(future::completeExceptionally)
                .onError((p, e) -> future.completeExceptionally(e))
                .check(CIPEncapsulationReadResponse.class::isInstance)
                .unwrap(CIPEncapsulationReadResponse.class::cast)
                .check(p -> p.getResponse().getTransactionCounter() == transactionCounter)
                .handle(p -> transaction.endRequest()));

                    //Not sure how to merge things back together ...

                    // Finish the request-transaction.
                    
            }  

            // This aborts reading other tags after sending the first tags request ... refactor.
        }
        //Should return an aggregated future ....
        
        return null;
    }

    @Override
    public void close(ConversationContext<CIPEncapsulationPacket> context) {
            throw new UnsupportedOperationException();
    }

    public PlcResponse decodeReadResponse(
        CIPEncapsulationReadResponse plcReadResponse, PlcReadRequest plcReadRequest, PlcValue plcValue) {
        Map<String, ResponseItem<PlcValue>> values = new HashMap<>();
        for (String tagName : plcReadRequest.getTagNames()) {
            AbEthTag tag = (AbEthTag) plcReadRequest.getTag(tagName);
            PlcResponseCode responseCode = decodeResponseCode(plcReadResponse.getResponse().getStatus());
            nuovoCod1(responseCode, tagName, tag);
            
            
            ResponseItem<PlcValue> result = new ResponseItem<>(responseCode, plcValue);
            values.put(tagName, result);
        }

        // Double check if it's really a InternalPlcReadRequest ...
        return new DefaultPlcReadResponse(plcReadRequest, values);
    }

    public void nuovoCod1(PlcResponseCode responseCode, String tagName, AbEthTag tag){
        if (responseCode == PlcResponseCode.OK) {
            try{
                //commento
            }
            catch (Exception e) {
                logger.warn("Some other error occurred casting tag {}, TagInformation: {}",tagName, tag,e);
            }
        }
    }

    public PlcValue nuovoCod2(CIPEncapsulationReadResponse plcReadResponse, PlcValue plcValue, String tagName, AbEthTag tag){
        List<Short> data;
        switch (tag.getFileType()) {
            
            case INTEGER: // output as single bytes
                if(plcReadResponse.getResponse() instanceof DF1CommandResponseMessageProtectedTypedLogicalRead) {
                    DF1CommandResponseMessageProtectedTypedLogicalRead df1PTLR = (DF1CommandResponseMessageProtectedTypedLogicalRead) plcReadResponse.getResponse();
                    data = df1PTLR.getData();
                    if(data.size() == 1) {
                        plcValue = new PlcINT(data.get(0));
                    }
                }
                break;
            case WORD:
                if(plcReadResponse.getResponse() instanceof DF1CommandResponseMessageProtectedTypedLogicalRead) {
                    DF1CommandResponseMessageProtectedTypedLogicalRead df1PTLR = (DF1CommandResponseMessageProtectedTypedLogicalRead) plcReadResponse.getResponse();
                    data = df1PTLR.getData();
                    if(data.size() == 1) {
                        plcValue = new PlcINT(data.get(0));
                    }
                }
                break;
            
            case DWORD:
                if(plcReadResponse.getResponse() instanceof DF1CommandResponseMessageProtectedTypedLogicalRead) {
                    DF1CommandResponseMessageProtectedTypedLogicalRead df1PTLR = (DF1CommandResponseMessageProtectedTypedLogicalRead) plcReadResponse.getResponse();
                    data = df1PTLR.getData();
                    if(data.size() == 1) {
                        plcValue = new PlcINT(data.get(0));
                    }
                }
                break;
            case SINGLEBIT:
                if(plcReadResponse.getResponse() instanceof DF1CommandResponseMessageProtectedTypedLogicalRead) {
                    DF1CommandResponseMessageProtectedTypedLogicalRead df1PTLR = (DF1CommandResponseMessageProtectedTypedLogicalRead) plcReadResponse.getResponse();
                    data = df1PTLR.getData();
                    if(data.size() == 1) {
                        plcValue = new PlcINT(data.get(0));
                    }
                }
                break;
            default:
                logger.warn("Problem during decoding of tag {}: Decoding of file type not implemented; " +
                    "TagInformation: {}", tagName, tag);
        }
        return plcValue;
    }

    private PlcResponseCode decodeResponseCode(short status) {
        if(status == 0) {
            return PlcResponseCode.OK;
        }
        return PlcResponseCode.NOT_FOUND;
    }

}
