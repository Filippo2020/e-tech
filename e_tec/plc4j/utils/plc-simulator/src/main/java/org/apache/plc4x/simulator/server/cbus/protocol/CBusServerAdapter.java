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
package org.apache.plc4x.simulator.server.cbus.protocol;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;
import org.apache.plc4x.java.cbus.readwrite.*;
import org.apache.plc4x.simulator.model.Context;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedList;
import java.util.List;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.ThreadLocalRandom;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.locks.Lock;
import java.util.concurrent.locks.ReentrantLock;

public class CBusServerAdapter extends ChannelInboundHandlerAdapter {

    private static final List<Byte> AVAILABLE_UNITS = Arrays.asList((byte) 3, (byte) 23, (byte) 48);

    private static final Logger LOGGER = LoggerFactory.getLogger(CBusServerAdapter.class);

    private static final RequestContext requestContext = new RequestContext(false);

    private static boolean connect;
    private static boolean smart;
    private static boolean idmon;
    private static boolean exstat;
    private static boolean monitor;
    private static boolean monall;
    private static boolean pun;
    private static boolean pcn;
    private static boolean srchk;

    private static CBusOptions cBusOptions;

    private static final  Lock outputLock = new ReentrantLock();

    private static ScheduledFuture<?> salMonitorFuture;

    private static ScheduledFuture<?> mmiMonitorFuture;
    private static String info1 = "Sending echo";
    
    private static byte monitorApplicationAddress1;
    private static byte monitorApplicationAddress2;


    public CBusServerAdapter(Context context) {
        LOGGER.info("Creating adapter with context {}", context);
    }

    private static void buildCBusOptions() {
        LOGGER.info("Updating options {}", cBusOptions);
        cBusOptions = new CBusOptions(connect, smart, idmon, exstat, monitor, monall, pun, pcn, srchk);
        LOGGER.info("Updated options {}", cBusOptions);
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        if (salMonitorFuture != null)
            salMonitorFuture.cancel(false);
        if (mmiMonitorFuture != null)
            mmiMonitorFuture.cancel(false);
        super.channelInactive(ctx);
    }

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        if (!(msg instanceof CBusMessage)) {
            return;
        }
        try {
            outputLock.lock();
            syncChannelRead(ctx, msg);
        } finally {
            ctx.flush();
            outputLock.unlock();
        }
    }

    private void syncChannelRead(ChannelHandlerContext ctx, Object msg) throws InterruptedException {
        LOGGER.debug("Working with cBusOptions\n{}", cBusOptions);
        // Serial is slow
        TimeUnit.MILLISECONDS.sleep(100);
        if (!smart && !connect) {
            // In this mode every message will be echoed
            LOGGER.info(info1);
            ctx.write(msg);
        }
        CBusMessage packet = (CBusMessage) msg;
        if (packet instanceof CBusMessageToClient) {
            LOGGER.info("Message to client not supported\n{}", packet);
            return;
        }
        CBusMessageToServer cBusMessageToServer = (CBusMessageToServer) packet;
        Request request = cBusMessageToServer.getRequest();
        if (request instanceof RequestEmpty || request instanceof RequestNull) {
            LOGGER.debug("Ignoring\n{}", request);
        } else if (request instanceof RequestDirectCommandAccess) {
            handleDirectCommandAccess(ctx, (RequestDirectCommandAccess) request);
        } else if (request instanceof RequestCommand) {
            handleRequestCommand(ctx, (RequestCommand) request);
        } else if (request instanceof RequestObsolete) {
            RequestObsolete requestObsolete = (RequestObsolete) request;
            LOGGER.info("Handling RequestObsolete\n{}", requestObsolete);
            // : handle this
        } else if (request instanceof RequestReset) {
            if (smart || connect) {
                // On reset, we need to send the echo if we had not sent it above
                LOGGER.info(info1);
                ctx.write(msg);
            }
            LOGGER.info("Handling RequestReset\n{}", request);
            handleReset();
        } else if (request instanceof RequestSmartConnectShortcut) {
            handleSmartConnect((RequestSmartConnectShortcut) request);
        }
    }


    public void handleSmartConnect(RequestSmartConnectShortcut request) {
        //da implementare 
    }

    private void handleDirectCommandAccess(ChannelHandlerContext ctx, RequestDirectCommandAccess requestDirectCommandAccess) {
        CALData calData = requestDirectCommandAccess.getCalData();
        LOGGER.info("Handling RequestDirectCommandAccess\n{}", requestDirectCommandAccess);
        handleCalData(ctx, calData, requestDirectCommandAccess.getAlpha());
    }

    private void handleRequestCommand(ChannelHandlerContext ctx, RequestCommand requestCommand) {
        LOGGER.info("Handling RequestCommand\n{}", requestCommand);
        CBusCommand cbusCommand = requestCommand.getCbusCommand();
        LOGGER.info("Handling CBusCommand\n{}", cbusCommand);
        if (cbusCommand instanceof CBusCommandPointToPoint) {
            CBusCommandPointToPoint cBusCommandPointToPoint = (CBusCommandPointToPoint) cbusCommand;
            CBusPointToPointCommand command = cBusCommandPointToPoint.getCommand();
            UnitAddress unitAddress = null;
            if (command instanceof CBusPointToPointCommandIndirect) {
                CBusPointToPointCommandIndirect cBusPointToPointCommandIndirect = (CBusPointToPointCommandIndirect) command;
                //handle bridgeAddress
                //handle networkRoute
                unitAddress = cBusPointToPointCommandIndirect.getUnitAddress();
            }
            if (command instanceof CBusPointToPointCommandDirect) {
                CBusPointToPointCommandDirect cBusPointToPointCommandDirect = (CBusPointToPointCommandDirect) command;
                unitAddress = cBusPointToPointCommandDirect.getUnitAddress();
            }
            if (unitAddress == null) {
                throw new IllegalStateException("Unit address should be set at this point");
            }
            boolean knownUnit = AVAILABLE_UNITS.contains(unitAddress.getAddress());
            if (!knownUnit) {
                LOGGER.warn("{} not a known unit", unitAddress);
                ReplyOrConfirmation replyOrConfirmation = new ServerErrorReply((byte) 0x0, cBusOptions, requestContext);
                CBusMessageToClient cBusMessageToClient = new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
                ctx.writeAndFlush(cBusMessageToClient);
                return;
            }
            CALData calData = command.getCalData();
            handleCalData(ctx, calData, requestCommand.getAlpha());
            return;
        } else if (cbusCommand instanceof CBusCommandPointToMultiPoint) {
            CBusCommandPointToMultiPoint cBusCommandPointToMultiPoint = (CBusCommandPointToMultiPoint) cbusCommand;
            CBusPointToMultiPointCommand command = cBusCommandPointToMultiPoint.getCommand();
            if (command instanceof CBusPointToMultiPointCommandStatus) {
                CBusPointToMultiPointCommandStatus cBusPointToMultiPointCommandStatus = (CBusPointToMultiPointCommandStatus) command;
                StatusRequest statusRequest = cBusPointToMultiPointCommandStatus.getStatusRequest();
                if (statusRequest instanceof StatusRequestBinaryState) {
                    StatusRequestBinaryState statusRequestBinaryState = (StatusRequestBinaryState) statusRequest;
                    LOGGER.info("Handling StatusRequestBinaryState\n{}", statusRequestBinaryState);
                    handleStatusRequestBinary(ctx, requestCommand, statusRequestBinaryState.getApplication());
                }
                if (statusRequest instanceof StatusRequestBinaryStateDeprecated) {
                    StatusRequestBinaryStateDeprecated statusRequestBinaryStateDeprecated = (StatusRequestBinaryStateDeprecated) statusRequest;
                    LOGGER.info("Handling StatusRequestBinaryStateDeprecated\n{}", statusRequestBinaryStateDeprecated);
                    handleStatusRequestBinary(ctx, requestCommand, statusRequestBinaryStateDeprecated.getApplication());
                    return;
                }
                if (statusRequest instanceof StatusRequestLevel) {
                    StatusRequestLevel statusRequestLevel = (StatusRequestLevel) statusRequest;
                    handleStatusRequestLevel(ctx, requestCommand, statusRequestLevel);
                    return;
                }
                throw new IllegalStateException();
            }
            if (command instanceof CBusPointToMultiPointCommandNormal) {
                CBusPointToMultiPointCommandNormal cBusPointToMultiPointCommandNormal = (CBusPointToMultiPointCommandNormal) command;
                LOGGER.info("Handling CBusPointToMultiPointCommandNormal\n{}", cBusPointToMultiPointCommandNormal);
                //handle this
            }
            //handle this
            return;
        } else if (cbusCommand instanceof CBusCommandPointToPointToMultiPoint) {
            CBusCommandPointToPointToMultiPoint cBusCommandPointToPointToMultiPoint = (CBusCommandPointToPointToMultiPoint) cbusCommand;
            LOGGER.info("Handling CBusCommandPointToPointToMultiPoint\n{}", cBusCommandPointToPointToMultiPoint);
            //handle this
            return;
        } else if (cbusCommand instanceof CBusCommandDeviceManagement) {
            CBusCommandDeviceManagement cBusCommandDeviceManagement = (CBusCommandDeviceManagement) cbusCommand;
            LOGGER.info("Handling CBusCommandDeviceManagement\n{}", cBusCommandDeviceManagement);
            //handle this
            return;
        }

        Alpha alpha = requestCommand.getAlpha();
        if (alpha != null) {
            Confirmation confirmation = new Confirmation(alpha, null, ConfirmationType.NOT_TRANSMITTED_CORRUPTION);
            ReplyOrConfirmationConfirmation replyOrConfirmationConfirmation = new ReplyOrConfirmationConfirmation(alpha.getCharacter(), confirmation, null, cBusOptions, requestContext);
            CBusMessage response = new CBusMessageToClient(replyOrConfirmationConfirmation, requestContext, cBusOptions);
            LOGGER.info("Send response\n{}", response);
            ctx.writeAndFlush(response);
        }
    }

    private static void handleStatusRequestLevel(ChannelHandlerContext ctx, RequestCommand requestCommand, StatusRequestLevel statusRequestLevel) {
        StatusCoding coding = StatusCoding.LEVEL_BY_THIS_SERIAL_INTERFACE;
        //map actual values from simulator
        byte blockStart = statusRequestLevel.getStartingGroupAddressLabel();
        List<LevelInformation> levelInformations = Collections.singletonList(new LevelInformationNormal(0x5555, LevelInformationNibblePair.Value_F, LevelInformationNibblePair.Value_F));
        CALData calData = new CALDataStatusExtended(CALCommandTypeContainer.CALCommandReply_4Bytes, null, coding, statusRequestLevel.getApplication(), blockStart, null, levelInformations, requestContext);
        CALReply calReply;
        if (exstat) {
            calReply = new CALReplyLong((byte) 0x0, calData, (byte) 0x0, new UnitAddress((byte) 0x04), null, new SerialInterfaceAddress((byte) 0x02), (byte) 0x0, null, cBusOptions, requestContext);
        } else {
            calReply = new CALReplyShort((byte) 0x0, calData, cBusOptions, requestContext);
        }
        CBusMessage response = createCBusMessageForReply(requestCommand.getAlpha(), calReply, cBusOptions);
        LOGGER.info("Send level status response\n{}", response);
        ctx.writeAndFlush(response);
    }

    private void handleStatusRequestBinary(ChannelHandlerContext ctx, RequestCommand requestCommand, ApplicationIdContainer application) {
        if (application == ApplicationIdContainer.NETWORK_CONTROL) {
            LOGGER.info("Handling installation MMI Request");
            sendInstallationMMIResponse(ctx, requestCommand.getAlpha());
            return;
        }
        
        List<StatusByte> statusBytes = new LinkedList<>();
        // : map actual values from simulator
        for (int i = 0; i < 22; i++) {
            statusBytes.add(new StatusByte(GAVState.ON, GAVState.ERROR, GAVState.OFF, GAVState.DOES_NOT_EXIST));
        }

        LOGGER.info("Send binary status response");
        sendStatusBytes(ctx, "First parts {}", application, (byte) 0x0, statusBytes, requestCommand.getAlpha(), cBusOptions);
    }

    private static void handleCalData(ChannelHandlerContext ctx, CALData calData, Alpha alpha) {
        if (calData instanceof CALDataGetStatus) {
            // : implement me
        } else if (calData instanceof CALDataWrite) {
            CALDataWrite calDataWrite = (CALDataWrite) calData;
            Runnable acknowledger = () -> {
                CALDataAcknowledge calDataAcknowledge = new CALDataAcknowledge(CALCommandTypeContainer.CALCommandAcknowledge, null, calDataWrite.getParamNo(), (short) 0x0, requestContext);
                CALReplyShort calReply = new CALReplyShort((byte) 0x0, calDataAcknowledge, cBusOptions, requestContext);
                EncodedReplyCALReply encodedReply = new EncodedReplyCALReply((byte) 0x0, calReply, cBusOptions, requestContext);
                ReplyEncodedReply replyEncodedReply = new ReplyEncodedReply((byte) 0x0, encodedReply, null, cBusOptions, requestContext);
                ReplyOrConfirmation replyOrConfirmation = new ReplyOrConfirmationReply((byte) 0x0, replyEncodedReply, new ResponseTermination(), cBusOptions, requestContext);
                if (alpha != null) {
                    replyOrConfirmation = new ReplyOrConfirmationConfirmation((byte) 0x0, new Confirmation(alpha, null, ConfirmationType.CONFIRMATION_SUCCESSFUL), replyOrConfirmation, cBusOptions, requestContext);
                }
                CBusMessageToClient cBusMessageToClient = new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
                LOGGER.info("Sending ack\n{}", cBusMessageToClient);
                ctx.writeAndFlush(cBusMessageToClient);
            };
            switch (calDataWrite.getParamNo().getParameterType()) {
                case APPLICATION_ADDRESS_1:
                    ApplicationAddress1 applicationAddress1 = ((ParameterValueApplicationAddress1) calDataWrite.getParameterValue()).getValue();
                    monitorApplicationAddress1 = applicationAddress1.getAddress();
                    acknowledger.run();
                    return;
                case APPLICATION_ADDRESS_2:
                    ApplicationAddress2 applicationAddress2 = ((ParameterValueApplicationAddress2) calDataWrite.getParameterValue()).getValue();
                    monitorApplicationAddress2 = applicationAddress2.getAddress();
                    acknowledger.run();
                    return;
                case INTERFACE_OPTIONS_1:
                    InterfaceOptions1 interfaceOptions1 = ((ParameterValueInterfaceOptions1) calDataWrite.getParameterValue()).getValue();
                    idmon = interfaceOptions1.getIdmon();
                    monitor = interfaceOptions1.getMonitor();
                    smart = interfaceOptions1.getSmart();
                    srchk = interfaceOptions1.getSrchk();
                    //add support for xonxoff
                    connect = interfaceOptions1.getConnect();
                    buildCBusOptions();
                    acknowledger.run();
                    return;
                case INTERFACE_OPTIONS_2:
                    pezzoCod1(acknowledger);
                    return;
                case INTERFACE_OPTIONS_3:
                    InterfaceOptions3 interfaceOptions3Value = ((ParameterValueInterfaceOptions3) calDataWrite.getParameterValue()).getValue();
                    exstat = interfaceOptions3Value.getExstat();
                    
                    pun = interfaceOptions3Value.getPun();
                    // : add support for localsal
                    pcn = interfaceOptions3Value.getPcn();
                    buildCBusOptions();
                    acknowledger.run();
                    return;
                case BAUD_RATE_SELECTOR:
                    pezzoCod1(acknowledger);
                    return;
                case INTERFACE_OPTIONS_1_POWER_UP_SETTINGS:
                    InterfaceOptions1 interfaceOptions1PowerUpSettings = ((ParameterValueInterfaceOptions1PowerUpSettings) calDataWrite.getParameterValue()).getValue().getInterfaceOptions1();
                    idmon = interfaceOptions1PowerUpSettings.getIdmon();
                    monitor = interfaceOptions1PowerUpSettings.getMonitor();
                    smart = interfaceOptions1PowerUpSettings.getSmart();
                    srchk = interfaceOptions1PowerUpSettings.getSrchk();
                    //add support for xonxoff
                    connect = interfaceOptions1PowerUpSettings.getConnect();
                    buildCBusOptions();
                    acknowledger.run();
                    return;
                case CUSTOM_MANUFACTURER:
                    //handle other param typed
                    acknowledger.run();
                    return;
                case SERIAL_NUMBER:
                    //handle other param typed
                    acknowledger.run();
                    return;
                case CUSTOM_TYPE:
                    //handle other param typed
                    acknowledger.run();
                    return;
                default:
                    throw new IllegalStateException("Unmapped type");
            }
        } else {
            throw new IllegalStateException("Unmapped type: " + calData.getClass());
        }
    }

    private static  Runnable pezzoCod1(Runnable acknowledger){
        buildCBusOptions();
        acknowledger.run();
        return acknowledger;
    }

    private static CBusMessage createCBusMessageForReply(Alpha alpha, CALReply calReply, CBusOptions cBusOptions) {
        EncodedReply encodedReply = new EncodedReplyCALReply((byte) 0x0, calReply, CBusServerAdapter.cBusOptions, CBusServerAdapter.requestContext);
        ReplyEncodedReply replyEncodedReply = new ReplyEncodedReply((byte) 0xC0, encodedReply, null, CBusServerAdapter.cBusOptions, CBusServerAdapter.requestContext);
        ReplyOrConfirmation replyOrConfirmation = new ReplyOrConfirmationReply((byte) 0xFF, replyEncodedReply, new ResponseTermination(), CBusServerAdapter.cBusOptions, CBusServerAdapter.requestContext);
        if (alpha != null) {
            Confirmation confirmation = new Confirmation(alpha, null, ConfirmationType.CONFIRMATION_SUCCESSFUL);
            replyOrConfirmation = new ReplyOrConfirmationConfirmation(alpha.getCharacter(), confirmation, replyOrConfirmation, CBusServerAdapter.cBusOptions, CBusServerAdapter.requestContext);
        }
        return new CBusMessageToClient(replyOrConfirmation, requestContext, cBusOptions);
    }

    private static void sendInstallationMMIResponse(ChannelHandlerContext ctx, Alpha alpha) {
        LOGGER.info("Send installation MMIs");
        sendMMIs(ctx, ApplicationIdContainer.NETWORK_CONTROL, alpha, cBusOptions);
    }

   


    private static void sendMMIsMetodo(ChannelHandlerContext ctx, ApplicationIdContainer application, Alpha alpha, CBusOptions cBusOptions, byte blockStart){
    
        List<StatusByte> unitStatusBytes = new LinkedList<>();
        LOGGER.debug("Produced {}, status bytes which equates to {} status", unitStatusBytes.size(), unitStatusBytes.size() * 4);
        sendStatusBytes(ctx, "Sending second part {}", application, blockStart, unitStatusBytes, alpha, cBusOptions);
    }

    


    private static void sendMMIs(ChannelHandlerContext ctx, ApplicationIdContainer application, Alpha alpha, CBusOptions cBusOptions) {
        byte blockStart;
        blockStart = 0x0;
        sendMMIsMetodo( ctx, application, alpha, cBusOptions, blockStart);
        blockStart = 88;
        sendMMIsMetodo( ctx, application, null, cBusOptions, blockStart);
        blockStart = (byte) 176;
        sendMMIsMetodo( ctx, application, null, cBusOptions, blockStart);        
    }

    private static void sendStatusBytes(ChannelHandlerContext ctx, String logMessage, ApplicationIdContainer application, byte blockStart, List<StatusByte> unitStatusBytes, Alpha alpha, CBusOptions cBusOptions) {
        int numberOfStatusBytes = unitStatusBytes.size();
        CALReply calReply = null;
        CALCommandTypeContainer commandTypeContainer = null;
        if (cBusOptions.getExstat()) {
            for (CALCommandTypeContainer calCommandTypeContainerElement : CALCommandTypeContainer.values()) {
                if (calCommandTypeContainerElement.getCommandType() != CALCommandType.STATUS_EXTENDED) {
                    continue;
                }
                if (calCommandTypeContainerElement.getNumBytes() + 3 == numberOfStatusBytes) {
                    commandTypeContainer = calCommandTypeContainerElement;
                }
            }
            CALData calData = new CALDataStatusExtended(commandTypeContainer, null, StatusCoding.BINARY_BY_THIS_SERIAL_INTERFACE, application, blockStart, unitStatusBytes, null, requestContext);
            // : do we use a random unit or a fixed or do we need it as parameter
            int randomElementIndex = 1;
            if(AVAILABLE_UNITS != null){
                randomElementIndex = ThreadLocalRandom.current().nextInt(AVAILABLE_UNITS.size()) % AVAILABLE_UNITS.size();
                byte randomUnit = AVAILABLE_UNITS.get(randomElementIndex);
                calReply = new CALReplyLong((byte) 0x86, calData, 0x00, new UnitAddress(randomUnit), null, new SerialInterfaceAddress((byte) 0x02), (byte) 0x00, null, cBusOptions, requestContext);
            }
            
        } else {
            CALData calData = new CALDataStatus(commandTypeContainer, null, application, (byte) 0x00, unitStatusBytes, requestContext);
            calReply = new CALReplyShort((byte) 0x0, calData, cBusOptions, requestContext);
        }
        CBusMessage response = createCBusMessageForReply(alpha, calReply, cBusOptions);
        LOGGER.debug(logMessage, response);
        ctx.writeAndFlush(response);
    }

    private static void handleReset() {
        connect = false;
        smart = false;
        idmon = false;
        exstat = false;
        monitor = false;
        monall = false;
        pun = false;
        pcn = false;
        srchk = false;
        stopSALMonitor();
        stopMMIMonitor();
    }

    private static void stopSALMonitor() {
        if (salMonitorFuture == null) {
            return;
        }
        LOGGER.info("Stopping SAL monitor");
        salMonitorFuture.cancel(false);
        salMonitorFuture = null;
    }

    private static void stopMMIMonitor() {
        if (mmiMonitorFuture == null) {
            return;
        }
        LOGGER.info("Stopping monitor");
        mmiMonitorFuture.cancel(false);
        mmiMonitorFuture = null;
    }

}
