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
package org.apache.plc4x.simulator.server.s7.protocol;

import io.netty.channel.*;

import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.simulator.model.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.BitSet;
import java.util.List;

public class S7Step7ServerAdapter extends ChannelInboundHandlerAdapter {

    private static final Logger LOGGER = LoggerFactory.getLogger(S7Step7ServerAdapter.class);

    private Context context;

    private State state;

    // COTP parameters
    private static final int LOCALREFERENCE = 42;
    
    
    private static final int LOCALTSAPID = 1;
    private int remoteTsapId = -1;
    private static final COTPTpduSize maxTpduSize = COTPTpduSize.SIZE_256;
    
    // S7 parameters
    // Set this to 1 as we don't want to handle stuff in parallel
    private static final int MAXAMQCALLER = 1;
    
    // Set this to 1 as we don't want to handle stuff in parallel
    private static final int MAXAMQCALLEE = 1;
    
    private static final int MAXPDULENGTH = 240;
    

    public S7Step7ServerAdapter(Context context) {
        this.context = context;
        state = State.INITIAL;
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        COTPTpduSize tpduSize = null;
        int remoteReference = -1;
        COTPProtocolClass protocolClass;
        if (msg instanceof TPKTPacket) {
            TPKTPacket packet = (TPKTPacket) msg;
            final COTPPacket cotpPacket = packet.getPayload();
            switch (state) {
                case INITIAL: {
                    if (!(cotpPacket instanceof COTPPacketConnectionRequest)) {
                        LOGGER.error("Expecting COTP Connection-Request");
                        return;
                    }

                    COTPTpduSize proposedTpduSize = null;
                    COTPPacketConnectionRequest cotpConnectionRequest = (COTPPacketConnectionRequest) cotpPacket;
                    for (COTPParameter parameter : cotpConnectionRequest.getParameters()) {
                        proposedTpduSize = cod5(parameter);
                    }

                    remoteReference = cotpConnectionRequest.getSourceReference();
                    protocolClass = cotpConnectionRequest.getProtocolClass();
                    try{
                        if (proposedTpduSize != null) {
                            tpduSize = (proposedTpduSize.getSizeInBytes() > maxTpduSize.getSizeInBytes()) ? maxTpduSize : proposedTpduSize;
                        } else {
                            LOGGER.error("error: proposedTpduSize is null");
                        }
                        
                    }catch(NullPointerException e){
                        LOGGER.error(String.format("error: %s", e));
                    }

                    // Prepare a response and send it back to the remote.
                    List<COTPParameter> parameters = new ArrayList<>();
                    parameters.add(new COTPParameterCalledTsap(remoteTsapId));
                    parameters.add(new COTPParameterCallingTsap(LOCALTSAPID));
                    try{
                        parameters.add(new COTPParameterTpduSize(tpduSize));
                    }catch(NullPointerException e){
                        LOGGER.error(String.format("error: %s", e));
                    }
                    COTPPacketConnectionResponse response = new COTPPacketConnectionResponse(
                        parameters, null, remoteReference, LOCALREFERENCE, protocolClass
                    );
                    ctx.writeAndFlush(new TPKTPacket(response));

                    state = State.COTP_CONNECTED;
                    break;
                }
                case COTP_CONNECTED: {
                    cod1(cotpPacket, ctx);
                    break;
                }
                case S7_CONNECTED: {
                    cod7(cotpPacket, ctx);
                    break;
                }
                default:
                    throw new IllegalStateException("Unexpected value: " + state);
            }
        }
    }

    private void cod1(COTPPacket cotpPacket, ChannelHandlerContext ctx){
        if (!(cotpPacket instanceof COTPPacketData)) {
            LOGGER.error("Expecting COTP Data packet");
            return;
        }

        COTPPacketData packetData = (COTPPacketData) cotpPacket;
        final short cotpTpduRef = packetData.getTpduRef();
        final S7Message payload = packetData.getPayload();
        if (!(payload instanceof S7MessageRequest)) {
            LOGGER.error("Expecting S7 Message Request");
            return;
        }
        S7MessageRequest s7MessageRequest = (S7MessageRequest) payload;
        final int s7TpduReference = s7MessageRequest.getTpduReference();
        final S7Parameter s7Parameter = s7MessageRequest.getParameter();
        if (!(s7Parameter instanceof S7ParameterSetupCommunication)) {
            LOGGER.error("Expecting S7 Message Request containing a S7 Setup Communication Parameter");
            return;
        }
        S7ParameterSetupCommunication s7ParameterSetupCommunication =
            (S7ParameterSetupCommunication) s7Parameter;

        int amqCaller = Math.min(s7ParameterSetupCommunication.getMaxAmqCaller(), MAXAMQCALLER);
        int amqCallee = Math.min(s7ParameterSetupCommunication.getMaxAmqCallee(), MAXAMQCALLEE);
        int pduLength = Math.min(s7ParameterSetupCommunication.getPduLength(), MAXPDULENGTH);

        S7ParameterSetupCommunication s7ParameterSetupCommunicationResponse =
            new S7ParameterSetupCommunication(amqCaller, amqCallee, pduLength);
        S7MessageResponseData s7MessageResponse = new S7MessageResponseData(
            s7TpduReference, s7ParameterSetupCommunicationResponse, null, (short) 0, (short) 0);
        ctx.writeAndFlush(new TPKTPacket(new COTPPacketData(null, s7MessageResponse, true, cotpTpduRef)));

        state = State.S7_CONNECTED;
    }

    private void cod2(S7PayloadUserDataItem userDataPayloadItem, S7ParameterUserDataItemCPUFunctions function, int s7TpduReference, short cotpTpduRef, ChannelHandlerContext ctx) throws IllegalAccessException{
        if (userDataPayloadItem instanceof S7PayloadUserDataItemCpuFunctionReadSzlRequest) {
            S7PayloadUserDataItemCpuFunctionReadSzlRequest readSzlRequestPayload =
                (S7PayloadUserDataItemCpuFunctionReadSzlRequest) userDataPayloadItem;

            final SzlId szlId = readSzlRequestPayload.getSzlId();
            // This is a request to list the type of device
            if ((szlId.getTypeClass() == SzlModuleTypeClass.CPU) &&
                (szlId.getSublistList() == SzlSublist.MODULE_IDENTIFICATION)) {

                S7ParameterUserDataItemCPUFunctions readSzlResponseParameter =
                    new S7ParameterUserDataItemCPUFunctions((short) 0x12,
                        (byte) 0x08, function.getCpuFunctionGroup(),
                        function.getCpuSubfunction(), (short) 1,
                        (short) 0, (short) 0, 0);

                // This is the product number of a S7-1200
                List<SzlDataTreeItem> items = new ArrayList<>();
                items.add(new SzlDataTreeItem((short) 0x0001,
                    "6ES7 212-1BD30-0XB0 ".getBytes(), 0x2020, 0x0001, 0x2020));


                List<S7ParameterUserDataItem> responseParameterItems = new ArrayList<>();
                responseParameterItems.add(readSzlResponseParameter);
                S7ParameterUserData responseParameterUserData =
                    new S7ParameterUserData(responseParameterItems);

                List<S7PayloadUserDataItem> responsePayloadItems = new ArrayList<>();
                responsePayloadItems.add(userDataPayloadItem);
                S7PayloadUserData responsePayloadUserData =
                    new S7PayloadUserData(responsePayloadItems);

                S7Message s7ResponseMessage = new S7MessageUserData(s7TpduReference,
                    responseParameterUserData, responsePayloadUserData);
                ctx.writeAndFlush(new TPKTPacket(new COTPPacketData(null, s7ResponseMessage, true, cotpTpduRef)));
            } else {
                throw new IllegalAccessException("Not able to respond to the given request Read SZL with SZL type class " +
                    szlId.getTypeClass().name() + " and SZL sublist " + szlId.getSublistList().name());
            }

        }
    }

    private void cod3(COTPPacket cotpPacket, short cotpTpduRef, ChannelHandlerContext ctx){
        if (cotpPacket.getPayload() instanceof S7MessageRequest) {
            S7MessageRequest request = (S7MessageRequest) cotpPacket.getPayload();
            if (request.getParameter() instanceof S7ParameterReadVarRequest) {
                S7ParameterReadVarRequest readVarRequestParameter =
                    (S7ParameterReadVarRequest) request.getParameter();
                List<S7VarRequestParameterItem> items = readVarRequestParameter.getItems();
                List<S7VarPayloadDataItem> payloadItems = new ArrayList<>();
                for (S7VarRequestParameterItem item : items) {
                    if (item instanceof S7VarRequestParameterItemAddress) {
                        S7VarRequestParameterItemAddress address =
                            (S7VarRequestParameterItemAddress) item;
                        final S7Address address1 = address.getAddress();
                        if (address1 instanceof S7AddressAny) {
                            cod4(address1, payloadItems);
                        }
                    }
                }
                S7ParameterReadVarResponse readVarResponseParameter = new S7ParameterReadVarResponse((short) items.size());
                S7PayloadReadVarResponse readVarResponsePayload = new S7PayloadReadVarResponse(payloadItems);
                S7MessageResponseData response = new S7MessageResponseData(request.getTpduReference(),
                    readVarResponseParameter, readVarResponsePayload, (short) 0x00, (short) 0x00);
                ctx.writeAndFlush(new TPKTPacket(new COTPPacketData(null, response, true, cotpTpduRef)));
            }
        }
    }

    private void cod4(S7Address address1, List<S7VarPayloadDataItem> payloadItems){
        S7AddressAny addressAny = (S7AddressAny) address1;
                            switch (addressAny.getArea()) {
                                case DATA_BLOCKS: {
                                    final int dataBlockNumber = addressAny.getDbNumber();
                                    if (dataBlockNumber != 1) {
                                        // Return unknown object.
                                    }
                                    final int numberOfElements = addressAny.getNumberOfElements();
                                    if (numberOfElements != 1) {
                                        // Return invalid address.
                                    }
                                    final int byteAddress = addressAny.getByteAddress();
                                    if (byteAddress != 0) {
                                        // Return invalid address.
                                    }
                                    switch (addressAny.getTransportSize()) {
                                        case BOOL:
                                            payloadItems.add(new S7VarPayloadDataItem(DataTransportErrorCode.OK, DataTransportSize.BIT, new byte[]{1}));
                                            break;
                                        case INT:
                                        case UINT: {
                                            short shortValue = 42; 
                                            byte[] data = new byte[2];
                                            data[1] = (byte) (shortValue & 0xff);
                                            data[0] = (byte) ((shortValue >> 8) & 0xff);
                                            payloadItems.add(new S7VarPayloadDataItem(DataTransportErrorCode.OK, DataTransportSize.BYTE_WORD_DWORD, data));
                                            break;
                                        }
                                        case BYTE:
                                        case CHAR:
                                        case DATE:
                                        case DATE_AND_TIME:
                                        case DINT:
                                        case DT:
                                        case DWORD:
                                        case LINT:
                                        case LREAL:
                                        case LTIME:
                                        case LWORD:
                                        case REAL:
                                        case SINT:
                                        case STRING:
                                        case TIME:
                                        case TIME_OF_DAY:
                                        case TOD:
                                        case UDINT:
                                        case ULINT:
                                        case USINT:
                                        case WCHAR:
                                        case WORD:
                                        case WSTRING:
                                        default:
                                    }
                                    break;
                                }
                                case INPUTS:
                                case OUTPUTS: {
                                    final int ioNumber = (addressAny.getByteAddress() * 8) + addressAny.getBitAddress();
                                    final int numElements = (addressAny.getTransportSize() == TransportSize.BOOL) ?
                                        addressAny.getNumberOfElements() : addressAny.getTransportSize().getSizeInBytes() * 8;
                                    final BitSet bitSet = toBitSet(context.getDigitalInputs(), ioNumber, numElements);
                                    final byte[] data = Arrays.copyOf(bitSet.toByteArray(), (numElements + 7) / 8);
                                    payloadItems.add(new S7VarPayloadDataItem(DataTransportErrorCode.OK, DataTransportSize.BYTE_WORD_DWORD, data));
                                    break;
                                }
                                case COUNTERS:
                                case DIRECT_PERIPHERAL_ACCESS:
                                case FLAGS_MARKERS:
                                case INSTANCE_DATA_BLOCKS:
                                case LOCAL_DATA:
                                case TIMERS:
                                default:
                            }
    }

    private COTPTpduSize cod5(COTPParameter parameter) throws IllegalAccessException{
        COTPTpduSize proposedTpduSize2 = null;
        if (parameter instanceof COTPParameterCalledTsap) {
            // this is actually ignored as it doesn't contain any information.
        } else if (parameter instanceof COTPParameterCallingTsap) {
            COTPParameterCallingTsap callingTsapParameter = (COTPParameterCallingTsap) parameter;
            remoteTsapId = callingTsapParameter.getTsapId();
        } else {
            throw new IllegalAccessException(String.format("Unexpected COTP Connection-Request Parameter %s",
                parameter.getClass().getName()));
        }
        if(parameter instanceof COTPParameterTpduSize){
            COTPParameterTpduSize tpduSizeParameter = (COTPParameterTpduSize) parameter;
            proposedTpduSize2 = tpduSizeParameter.getTpduSize();
        }
        return proposedTpduSize2;
    }

    private void cod6(S7Message payload, short cotpTpduRef, ChannelHandlerContext ctx) throws IllegalAccessException{
        S7MessageUserData s7MessageUserData = (S7MessageUserData) payload;
                        final int s7TpduReference = s7MessageUserData.getTpduReference();
                        final S7Parameter s7Parameter = s7MessageUserData.getParameter();
                        if (s7Parameter instanceof S7ParameterUserData) {
                            S7ParameterUserData userDataParameter = (S7ParameterUserData) s7Parameter;
                            for (S7ParameterUserDataItem item : userDataParameter.getItems()) {
                                if (item instanceof S7ParameterUserDataItemCPUFunctions) {
                                    S7ParameterUserDataItemCPUFunctions function =
                                        (S7ParameterUserDataItemCPUFunctions) item;
                                    final S7PayloadUserData userDataPayload =
                                        (S7PayloadUserData) s7MessageUserData.getPayload();

                                    for (S7PayloadUserDataItem userDataPayloadItem : userDataPayload.getItems()) {
                                        cod2(userDataPayloadItem, function, s7TpduReference, cotpTpduRef, ctx);
                                    }
                                }
                            }
                        }
    }

    private void cod7(COTPPacket cotpPacket, ChannelHandlerContext ctx) throws IllegalAccessException{
        if (!(cotpPacket instanceof COTPPacketData)) {
            LOGGER.error("Expecting COTP Data packet");
            return;
        }

        COTPPacketData packetData = (COTPPacketData) cotpPacket;
        final short cotpTpduRef = packetData.getTpduRef();
        final S7Message payload = packetData.getPayload();
        if (payload instanceof S7MessageUserData) {
            cod6(payload, cotpTpduRef, ctx);
        } else {
            cod3(cotpPacket, cotpTpduRef, ctx);
        }
    }

    private enum State {
        INITIAL,
        COTP_CONNECTED,
        S7_CONNECTED
    }

    private BitSet toBitSet(List<Boolean> booleans, int startIndex, int numElements) {
        BitSet bitSet = new BitSet(booleans.size());
        for (int i = 0; i < Math.min(booleans.size() - startIndex, numElements); i++) {
            bitSet.set(i, booleans.get(i + startIndex));
        }
        return bitSet;
    }

}
