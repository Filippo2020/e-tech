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
package org.apache.plc4x.java.knxnetip.protocol;

import io.netty.channel.socket.DatagramChannel;
import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.knxnetip.context.KnxNetIpDriverContext;
import org.apache.plc4x.java.knxnetip.ets.model.EtsModel;
import org.apache.plc4x.java.knxnetip.ets.model.GroupAddress;
import org.apache.plc4x.java.knxnetip.tag.KnxNetIpTag;
import org.apache.plc4x.java.knxnetip.model.KnxNetIpSubscriptionHandle;
import org.apache.plc4x.java.knxnetip.readwrite.KnxGroupAddress;
import org.apache.plc4x.java.knxnetip.readwrite.KnxGroupAddress2Level;
import org.apache.plc4x.java.knxnetip.readwrite.KnxGroupAddress3Level;
import org.apache.plc4x.java.knxnetip.readwrite.KnxGroupAddressFreeLevel;
import org.apache.plc4x.java.knxnetip.readwrite.KnxDatapoint;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.knxnetip.readwrite.*;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.*;
import org.apache.plc4x.java.spi.messages.*;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.model.DefaultPlcConsumerRegistration;
import org.apache.plc4x.java.spi.model.DefaultPlcSubscriptionTag;
import org.apache.plc4x.java.spi.values.PlcSTRING;
import org.apache.plc4x.java.spi.values.PlcStruct;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetSocketAddress;
import java.time.Duration;
import java.time.Instant;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.atomic.AtomicInteger;
import java.util.function.Consumer;

public class KnxNetIpProtocolLogic extends Plc4xProtocolBase<KnxNetIpMessage> implements PlcSubscriber {

    private static final Logger LOGGER = LoggerFactory.getLogger(KnxNetIpProtocolLogic.class);
    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);

    private KnxNetIpDriverContext knxNetIpDriverContext;
    private Timer connectionStateTimer;
    private static final AtomicInteger sequenceCounter = new AtomicInteger(0);
    

    private final Map<DefaultPlcConsumerRegistration, Consumer<PlcSubscriptionEvent>> consumers = new ConcurrentHashMap<>();

    @Override
    public void setDriverContext(DriverContext driverContext) {
        super.setDriverContext(driverContext);
        this.knxNetIpDriverContext = (KnxNetIpDriverContext) driverContext;

        // Initialize Transaction Manager.
        // Until the number of concurrent requests is successfully negotiated we set it to a
        // maximum of only one request being able to be sent at a time. During the login process
        // No concurrent requests can be sent anyway. It will be updated when receiving the
        // S7ParameterSetupCommunication response.
       
    }

    @Override
    public void onConnect(ConversationContext<KnxNetIpMessage> context) {
        // Only the UDP transport supports login.
        if (!context.isPassive()) {
            LOGGER.info("KNX Driver running in ACTIVE mode.");
            knxNetIpDriverContext.setPassiveMode(false);

            DatagramChannel channel = (DatagramChannel) context.getChannel();
            final InetSocketAddress localSocketAddress = channel.localAddress();
            knxNetIpDriverContext.setLocalIPAddress(new IPAddress(localSocketAddress.getAddress().getAddress()));
            knxNetIpDriverContext.setLocalPort(localSocketAddress.getPort());

            // First send out a search request
            // REMARK: This might be optional ... usually we would send a search request to ip 224.0.23.12
            // Any KNX Gateway will respond with a search response. We're currently directly sending to the
            // known gateway address, so it's sort of pointless, but at least only one device will respond.
            LOGGER.info("Sending KNXnet/IP Search Request.");
            SearchRequest searchRequest = new SearchRequest(
                new HPAIDiscoveryEndpoint(HostProtocolCode.IPV4_UDP,
                    knxNetIpDriverContext.getLocalIPAddress(), knxNetIpDriverContext.getLocalPort()));
            context.sendRequest(searchRequest)
                .expectResponse(KnxNetIpMessage.class, Duration.ofMillis(1000))
                .check(SearchResponse.class::isInstance)
                .unwrap(p -> (SearchResponse) p)
                .handle(searchResponse -> {
                    LOGGER.info("Got KNXnet/IP Search Response.");
                    // Check if this device supports tunneling services.
                    final ServiceId tunnelingService = searchResponse.getDibSuppSvcFamilies().getServiceIds().stream().filter(serviceId -> serviceId instanceof KnxNetIpTunneling).findFirst().orElse(null);

                    // If this device supports this type of service, tell the driver, we found a suitable device.
                    if (tunnelingService != null) {
                        // Extract the required information form the search request.
                        knxNetIpDriverContext.setGatewayAddress(searchResponse.getDibDeviceInfo().getKnxAddress());
                        knxNetIpDriverContext.setGatewayName(new String(searchResponse.getDibDeviceInfo().getDeviceFriendlyName()).trim());

                        LOGGER.info(String.format("Found KNXnet/IP Gateway '%s' with KNX address '%d.%d.%d'",
                            knxNetIpDriverContext.getGatewayName(),
                            knxNetIpDriverContext.getGatewayAddress().getMainGroup(),
                            knxNetIpDriverContext.getGatewayAddress().getMiddleGroup(),
                            knxNetIpDriverContext.getGatewayAddress().getSubGroup()));

                        // Next send a connection request to the gateway.
                        ConnectionRequest connectionRequest = new ConnectionRequest(
                            new HPAIDiscoveryEndpoint(HostProtocolCode.IPV4_UDP,
                                knxNetIpDriverContext.getLocalIPAddress(), knxNetIpDriverContext.getLocalPort()),
                            new HPAIDataEndpoint(HostProtocolCode.IPV4_UDP,
                                knxNetIpDriverContext.getLocalIPAddress(), knxNetIpDriverContext.getLocalPort()),
                            new ConnectionRequestInformationTunnelConnection(
                                knxNetIpDriverContext.getTunnelConnectionType()));
                        LOGGER.info("Sending KNXnet/IP Connection Request.");
                        context.sendRequest(connectionRequest)
                            .expectResponse(KnxNetIpMessage.class, Duration.ofMillis(1000))
                            .check(ConnectionResponse.class::isInstance)
                            .unwrap(p -> (ConnectionResponse) p)
                            .handle(connectionResponse -> {
                                // Remember the communication channel id.
                                knxNetIpDriverContext.setCommunicationChannelId(
                                    connectionResponse.getCommunicationChannelId());

                                LOGGER.info(String.format("Received KNXnet/IP Connection Response (Connection Id %s)",
                                    knxNetIpDriverContext.getCommunicationChannelId()));

                                
                            });
                    } else {
                        // This device doesn't support tunneling ... do some error handling.
                        LOGGER.error("Not connected to KNCnet/IP Gateway. The device doesn't support Tunneling.");
                        //Actively disconnect
                    }
                });
        }
        // This usually when we're running a passive mode river.
        else {
            LOGGER.info("KNX Driver running in PASSIVE mode.");
            knxNetIpDriverContext.setPassiveMode(true);

            // No login required, just confirm that we're connected.
            context.fireConnected();
        }
    }

    @Override
    public void onDisconnect(ConversationContext<KnxNetIpMessage> context) {
        // Cancel the timer for sending connection state requests.
        connectionStateTimer.cancel();

        DisconnectRequest disconnectRequest = new DisconnectRequest(knxNetIpDriverContext.getCommunicationChannelId(),
            new HPAIControlEndpoint(HostProtocolCode.IPV4_UDP,
                knxNetIpDriverContext.getLocalIPAddress(), knxNetIpDriverContext.getLocalPort()));
        context.sendRequest(disconnectRequest)
            .expectResponse(KnxNetIpMessage.class, Duration.ofMillis(1000))
            .check(DisconnectResponse.class::isInstance)
            .unwrap(p -> (DisconnectResponse) p)
            .handle(disconnectResponse -> {
                // In general we should probably check if the disconnect was successful, but in
                // the end we couldn't do much if the disconnect would fail.
                final String gatewayName = knxNetIpDriverContext.getGatewayName();
                final KnxAddress gatewayAddress = knxNetIpDriverContext.getGatewayAddress();
                LOGGER.info(String.format("Disconnected from KNX Gateway '%s' with KNX address '%d.%d.%d'", gatewayName,
                    gatewayAddress.getMainGroup(), gatewayAddress.getMiddleGroup(), gatewayAddress.getSubGroup()));

                // Send an event that connection disconnect is complete.
                context.fireDisconnected();
                LOGGER.debug("Disconnected event fired from KNX protocol");
            });
    }

    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();
        DefaultPlcWriteRequest request = (DefaultPlcWriteRequest) writeRequest;

        // As the KNX driver is using the SingleTagOptimizer, each request here will have
        // only one item.
        final Optional<String> first = request.getTagNames().stream().findFirst();
        if (first.isPresent()) {
            String tagName = first.get();
            final KnxNetIpTag tag = (KnxNetIpTag) request.getTag(tagName);
            byte[] destinationAddress = toKnxAddressData(tag);
            if (sequenceCounter.get() == Short.MAX_VALUE) {
                sequenceCounter.set(0);
            }

            final EtsModel etsModel = knxNetIpDriverContext.getEtsModel();
            if (etsModel != null) {
                final String destinationAddressString = etsModel.parseGroupAddress(destinationAddress);
                final GroupAddress groupAddress = etsModel.getGroupAddresses().get(destinationAddressString);
                if ((groupAddress == null) || (groupAddress.getType() == null)) {
                    future.completeExceptionally(new PlcRuntimeException(
                        "ETS5 model didn't specify group address '" + destinationAddressString +
                            "' or didn't define a type for it."));
                    return future;
                }

                // Use the data in the ets model to correctly check and serialize the PlcValue
                
            }
            
        }
        return future;
    }

    @Override
    protected void decode(ConversationContext<KnxNetIpMessage> context, KnxNetIpMessage msg) {
        try{
        // Handle a normal tunneling request, which is delivering KNX data.
        if (msg instanceof TunnelingRequest) {
            TunnelingRequest tunnelingRequest = (TunnelingRequest) msg;
            final short curCommunicationChannelId =
                tunnelingRequest.getTunnelingRequestDataBlock().getCommunicationChannelId();

            // Only if the communication channel id match, do anything with the request.
            // In case of a passive-mode driver we'll simply accept all communication ids.
            if (knxNetIpDriverContext.isPassiveMode() ||
                (curCommunicationChannelId == knxNetIpDriverContext.getCommunicationChannelId())) {
                // Data packets received from a link layer tunneling connection.
                if (tunnelingRequest.getCemi() instanceof LDataInd) {
                    LDataInd dataInd = (LDataInd) tunnelingRequest.getCemi();
                    final LDataFrame lDataFrame = dataInd.getDataFrame();
                    if (lDataFrame instanceof LDataExtended) {
                        LDataExtended lDataFrameDataExt = (LDataExtended) lDataFrame;
                        Apdu apdu = lDataFrameDataExt.getApdu();
                        if (apdu instanceof ApduDataContainer) {
                            ApduDataContainer apduDataContainer = (ApduDataContainer) apdu;
                            ApduData dataApdu = apduDataContainer.getDataApdu();
                            if (dataApdu instanceof ApduDataGroupValueWrite) {
                                ApduDataGroupValueWrite groupWrite = (ApduDataGroupValueWrite) dataApdu;
                                processCemiData(lDataFrameDataExt.getSourceAddress(), lDataFrameDataExt.getDestinationAddress(),
                                    groupWrite.getDataFirstByte(), groupWrite.getData());
                            }
                        }
                    }
                }
                // Data packets received from a busmonitor tunneling connection.
                

                // Confirm receipt of the request.
                final short sequenceCounter2 = tunnelingRequest.getTunnelingRequestDataBlock().getSequenceCounter();
                TunnelingResponse tunnelingResponse = new TunnelingResponse(
                    new TunnelingResponseDataBlock(knxNetIpDriverContext.getCommunicationChannelId(), sequenceCounter2,
                        Status.NO_ERROR));
                context.sendToWire(tunnelingResponse);
            }
        } else if (msg instanceof TunnelingResponse) {
            // This is just handling of all the Ack messages that might come in for any read- or
            // write requests. Usually this is just OK (I haven't managed to fake a message to cause
            // something else than an OK)
        }
    }catch(Exception e){
        throw new IllegalArgumentException(e);
    }
    }

    protected void processCemiData(KnxAddress sourceAddress, byte[] destinationGroupAddress,
                                   byte firstByte, byte[] restBytes) throws ParseException {
        // The first byte is actually just 6 bit long, but we'll treat it as a full one.
        // So here we create a byte array containing the first and all the following bytes.
        byte[] payload = new byte[1 + restBytes.length];
        payload[0] = firstByte;
        System.arraycopy(restBytes, 0, payload, 1, restBytes.length);

        // Decode the group address depending on the project settings.
        ReadBuffer addressBuffer = new ReadBufferByteBased(destinationGroupAddress);
        final KnxGroupAddress knxGroupAddress =
            KnxGroupAddress.staticParse(addressBuffer, knxNetIpDriverContext.getGroupAddressType());
        final String destinationAddress = toString(knxGroupAddress);

        // If there is an ETS model provided, continue decoding the payload.
        if (knxNetIpDriverContext.getEtsModel() != null) {
            final EtsModel etsModel = knxNetIpDriverContext.getEtsModel();
            final GroupAddress groupAddress = etsModel.getGroupAddresses().get(destinationAddress);
            final String areaName = etsModel.getTopologyName(destinationAddress.substring(
                0, destinationAddress.indexOf('/')));
            final String lineName = etsModel.getTopologyName(destinationAddress.substring(
                0, destinationAddress.indexOf('/', destinationAddress.indexOf('/') + 1)));

            if ((groupAddress != null) && (groupAddress.getType() != null)) {
                String a1 ="Message from: '%s' to: '%s'" +
                    toString(sourceAddress) + destinationAddress;
                LOGGER.debug(a1);

                // Parse the payload depending on the type of the group-address.
                ReadBuffer rawDataReader = new ReadBufferByteBased(payload);
                final PlcValue value = KnxDatapoint.staticParse(rawDataReader,
                    groupAddress.getType());

                // Assemble the plc4x return data-structure.
                Map<String, PlcValue> dataPointMap = new HashMap<>();
                dataPointMap.put("sourceAddress", new PlcSTRING(toString(sourceAddress)));
                dataPointMap.put("targetAddress", new PlcSTRING(groupAddress.getGroupAddress()));
                if (groupAddress.getFunction() != null) {
                    dataPointMap.put("location", new PlcSTRING(groupAddress.getFunction().getSpaceName()));
                    dataPointMap.put("function", new PlcSTRING(groupAddress.getFunction().getName()));
                } else {
                    dataPointMap.put("location", null);
                    dataPointMap.put("function", null);
                }
                if (areaName != null) {
                    dataPointMap.put("area", new PlcSTRING(areaName));
                }
                if (lineName != null) {
                    dataPointMap.put("line", new PlcSTRING(lineName));
                }
                dataPointMap.put("description", new PlcSTRING(groupAddress.getName()));
                dataPointMap.put("unitOfMeasurement", new PlcSTRING(groupAddress.getType().getName()));
                dataPointMap.put("value", value);
                final PlcStruct dataPoint = new PlcStruct(dataPointMap);

                // Send the data-structure.
                publishEvent(groupAddress, dataPoint);
            } else {
                    String a2 = "Message from: '%s' to unknown group address: '%s'%n payload: '%s'" +
                        toString(sourceAddress) + destinationAddress + Hex.encodeHexString(payload);
                    LOGGER.warn(a2);
            }
        }
        // Else just output the raw payload.
        else {
               String a3 = "Raw Message: '%s' to: '%s'%n payload: '%s'" +
                KnxNetIpProtocolLogic.toString(sourceAddress) + destinationAddress +
                Hex.encodeHexString(payload);
                LOGGER.info(a3);
        }
    }

    @Override
    public void close(ConversationContext<KnxNetIpMessage> context) {
        // Implement Closing on Protocol Level
    }

    @Override
    public CompletableFuture<PlcSubscriptionResponse> subscribe(PlcSubscriptionRequest subscriptionRequest) {
        Map<String, ResponseItem<PlcSubscriptionHandle>> values = new HashMap<>();
        for (String tagName : subscriptionRequest.getTagNames()) {
            final DefaultPlcSubscriptionTag tag = (DefaultPlcSubscriptionTag) subscriptionRequest.getTag(tagName);
            if (!(tag.getTag() instanceof KnxNetIpTag)) {
                values.put(tagName, new ResponseItem<>(PlcResponseCode.INVALID_ADDRESS, null));
            } else {
                values.put(tagName, new ResponseItem<>(PlcResponseCode.OK,
                    new KnxNetIpSubscriptionHandle(this, (KnxNetIpTag) tag.getTag())));
            }
        }
        return CompletableFuture.completedFuture(
            new DefaultPlcSubscriptionResponse(subscriptionRequest, values));
    }

    @Override
    public PlcConsumerRegistration register(Consumer<PlcSubscriptionEvent> consumer, Collection<PlcSubscriptionHandle> collection) {
        final DefaultPlcConsumerRegistration consumerRegistration =
            new DefaultPlcConsumerRegistration(this, consumer, collection.toArray(new PlcSubscriptionHandle[0]));
        consumers.put(consumerRegistration, consumer);
        return consumerRegistration;
    }

    @Override
    public void unregister(PlcConsumerRegistration plcConsumerRegistration) {
        DefaultPlcConsumerRegistration consumerRegistration = (DefaultPlcConsumerRegistration) plcConsumerRegistration;
        consumers.remove(consumerRegistration);
    }

    protected void publishEvent(GroupAddress groupAddress, PlcValue plcValue) {
        // Create a subscription event from the input.
        // Check this ... this is sort of not really right ...
        final PlcSubscriptionEvent event = new DefaultPlcSubscriptionEvent(Instant.now(),
            Collections.singletonMap("knxData", new ResponseItem<>(PlcResponseCode.OK, plcValue)));

        // Try sending the subscription event to all listeners.
        for (Map.Entry<DefaultPlcConsumerRegistration, Consumer<PlcSubscriptionEvent>> entry : consumers.entrySet()) {
            final DefaultPlcConsumerRegistration registration = entry.getKey();
            final Consumer<PlcSubscriptionEvent> consumer = entry.getValue();
            // Only if the current data point matches the subscription, publish the event to it.
            for (PlcSubscriptionHandle handle : registration.getSubscriptionHandles()) {
                if (handle instanceof KnxNetIpSubscriptionHandle) {
                    KnxNetIpSubscriptionHandle subscriptionHandle = (KnxNetIpSubscriptionHandle) handle;
                    // Check if the subscription matches this current event.
                    if (subscriptionHandle.getTag().matchesGroupAddress(groupAddress)) {
                        consumer.accept(event);
                    }
                }
            }
        }
    }

    protected byte[] toKnxAddressData(KnxNetIpTag tag) {
        WriteBufferByteBased address = new WriteBufferByteBased(2);
        try {
            switch (knxNetIpDriverContext.getGroupAddressType()) {
                case 3:
                    address.writeUnsignedShort(5, Short.parseShort(tag.getMainGroup()));
                    address.writeUnsignedByte(3, Byte.parseByte(tag.getMiddleGroup()));
                    address.writeUnsignedShort(8, Short.parseShort(tag.getSubGroup()));
                    break;
                case 2:
                    address.writeUnsignedShort(5, Short.parseShort(tag.getMainGroup()));
                    address.writeUnsignedShort(11, Short.parseShort(tag.getSubGroup()));
                    break;
                case 1:
                    address.writeUnsignedShort(16, Short.parseShort(tag.getSubGroup()));
                    break;
                default:
            }
        } catch (Exception e) {
            throw new PlcRuntimeException("Error converting tag into knx address data.", e);
        }
        return address.getData();
    }

    protected static String toString(KnxAddress knxAddress) {
        return knxAddress.getMainGroup() + "." + knxAddress.getMiddleGroup() + "." + knxAddress.getSubGroup();
    }

    protected static String toString(KnxGroupAddress groupAddress) {
        if (groupAddress instanceof KnxGroupAddress3Level) {
            KnxGroupAddress3Level level3 = (KnxGroupAddress3Level) groupAddress;
            return level3.getMainGroup() + "/" + level3.getMiddleGroup() + "/" + level3.getSubGroup();
        } else if (groupAddress instanceof KnxGroupAddress2Level) {
            KnxGroupAddress2Level level2 = (KnxGroupAddress2Level) groupAddress;
            return level2.getMainGroup() + "/" + level2.getSubGroup();
        } else if (groupAddress instanceof KnxGroupAddressFreeLevel) {
            KnxGroupAddressFreeLevel free = (KnxGroupAddressFreeLevel) groupAddress;
            return free.getSubGroup() + "";
        }
        throw new PlcRuntimeException("Unsupported Group Address Type " + groupAddress.getClass().getName());
    }

}
