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
package org.apache.plc4x.java.opcuaserver.backend;

import java.util.Arrays;

import org.apache.plc4x.java.api.types.PlcValueType;
import org.eclipse.milo.opcua.sdk.server.AbstractLifecycle;
import org.eclipse.milo.opcua.sdk.server.api.DataItem;
import org.eclipse.milo.opcua.sdk.server.nodes.filters.AttributeFilterContext;
import org.eclipse.milo.opcua.stack.core.Identifiers;
import org.eclipse.milo.opcua.stack.core.types.builtin.DataValue;
import org.eclipse.milo.opcua.stack.core.types.builtin.NodeId;
import org.eclipse.milo.opcua.stack.core.types.builtin.StatusCode;
import org.eclipse.milo.opcua.stack.core.types.builtin.Variant;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadRequest.Builder;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;
import org.apache.plc4x.java.utils.connectionpool.*;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;

import org.apache.plc4x.java.api.model.PlcTag;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.Map;
import java.util.HashMap;


public class Plc4xCommunication extends AbstractLifecycle {

    private PlcDriverManager driverManager;
    private final Logger logger = LoggerFactory.getLogger(getClass());
    private static final Integer DEFAULT_TIMEOUT = 1000000;
    private static final Integer DEFAULT_RETRY_BACKOFF = 5000;
    private final DataValue bADrESPONSE = new DataValue(new Variant(null), StatusCode.BAD);

    private Map<String, Long> failedConnectionList = new HashMap<>();

    Map<NodeId, DataItem> monitoredList = new HashMap<>();

    public Plc4xCommunication () {
        
      //document why this constructor is empty
    }

    @Override
    protected void onStartup() {
        driverManager = new PooledPlcDriverManager();
    }

    @Override
    protected void onShutdown() {
        //Do Nothing
    }

    public PlcDriverManager getDriverManager() {
        return driverManager;
    }

    public void setDriverManager(PlcDriverManager driverManager) {
        this.driverManager =  driverManager;
    }

    public PlcTag getTag(String tag, String connectionString) throws PlcConnectionException {
        return driverManager.getDriverForUrl(connectionString).prepareTag(tag);
    }

    public void addTag(DataItem item) {
        String format ="Adding item to OPC UA monitored list %s" + item.getReadValueId();
        logger.info(format);
        monitoredList.put(item.getReadValueId().getNodeId(), item);
    }

    public void removeTag(DataItem item) {
        String format = "Removing item from OPC UA monitored list %s" + item.getReadValueId();
        logger.info(format);
        monitoredList.remove(item.getReadValueId().getNodeId());
    }

    public static NodeId getNodeId(PlcValueType plcValueType) {
        switch (plcValueType) {
            case BOOL:
                return Identifiers.Boolean;
            case BYTE:
            case USINT:
                return Identifiers.Byte;
            case SINT:
                return Identifiers.SByte;
            case INT:
                return Identifiers.Int16;
            case WORD:
            case UINT:
                return Identifiers.UInt16;
            case DINT:
                return Identifiers.Int32;
            case DWORD:
            case UDINT:
                return Identifiers.UInt32;
            case LINT:
                return Identifiers.Int64;
            case ULINT:
            case LWORD:
                return Identifiers.UInt64;
            case REAL:
                return Identifiers.Float;
            case LREAL:
                return Identifiers.Double;
            case CHAR:
            case WCHAR:
            case STRING:
            case WSTRING:
                return Identifiers.String;
            default:
                return Identifiers.BaseDataType;
        }
    }

    public DataValue getValue(AttributeFilterContext.GetAttributeContext ctx, String tag, String connectionString) throws Exception {
        PlcConnection connection = prcd();
        try {

            //Check if we just polled the connection and it failed. Wait for the backoff counter to expire before we try again.
            if (failedConnectionList.containsKey(connectionString)) {
                if (System.currentTimeMillis() > failedConnectionList.get(connectionString) + DEFAULT_RETRY_BACKOFF) {
                    failedConnectionList.remove(connectionString);
                } else {
                    throw new IllegalAccessException("Waiting for back off timer - " + ((failedConnectionList.get(connectionString) + DEFAULT_RETRY_BACKOFF) - System.currentTimeMillis()) + " ms left");
                }
            }

            //Try to connect to PLC
            nested3(connectionString);

            long timeout = DEFAULT_TIMEOUT;
            if (monitoredList.containsKey(ctx.getNode().getNodeId())) {
                timeout = (long) monitoredList.get(ctx.getNode().getNodeId()).getSamplingInterval() * 1000;
            }

            // Create a new read request:
            // - Give the single item requested an alias name
            PlcReadRequest.Builder builder = connection.readRequestBuilder();
            builder.addTagAddress("value-1", tag);
            PlcReadRequest readRequest = builder.build();

            nested(readRequest, timeout, connection);
            DataValue resp = bADrESPONSE;

            nested4(connection, connectionString);

            return resp;
        } catch (Exception e) {
            logger.debug(String.valueOf("General error reading value " + e.getStackTrace()[0].toString()));
                connection.close();
            return bADrESPONSE;
        }
    }

    public DataValue nested(PlcReadRequest readRequest, long timeout, PlcConnection connection) throws Exception{
        try {
            readRequest.execute().get(timeout, TimeUnit.MICROSECONDS);
        } catch (InterruptedException e) {
            logger.debug(String.valueOf(String.format("%s Occurred while reading value, using timeout of %s",e, (timeout / 1000 + "ms"))));
            Thread.currentThread().interrupt();
            try {
                connection.close();
            } catch (Exception exception) {
                throw new IllegalArgumentException("Closing connection failed with error - " + exception);
            }
            return bADrESPONSE;
        }
        return bADrESPONSE;
    }

    public void nested2(PlcConnection connection){
        try {
            connection.close();
        } catch (Exception e) {
            throw new IllegalArgumentException("Closing connection failed with error - " + e);
        }
    }

    public DataValue nested3(String connectionString) throws IllegalAccessException{
        try {
            driverManager.getConnection(connectionString);
            throw new IllegalAccessException(connectionString + " Connected");
        } catch (PlcConnectionException e) {
            String format = "Failed to connect to device, error raised - " + e;
            logger.info(format);
            failedConnectionList.put(connectionString, System.currentTimeMillis());
            return bADrESPONSE;
        }
    }

    public void nested4(PlcConnection connection, String connectionString){
        try {
            connection.close();
        } catch (Exception e) {
            failedConnectionList.put(connectionString, System.currentTimeMillis());
            logger.warn(String.format("Closing connection failed with error %s",e));
        }
    }

    public void setValue(String tag, String value, String connectionString) {
        PlcConnection connection = prcd();
        try {
          connection = driverManager.getConnection(connectionString);
          if (!connection.isConnected()) {
              logger.debug("getConnection() returned a connection that isn't connected");
              connection.connect();
          }
        } catch (PlcConnectionException e) {
          logger.warn(String.format("Failed %s", e));
        }

        if (!connection.getMetadata().canWrite()) {
            logger.error("This connection doesn't support writing.");
            try {
              connection.close();
            } catch (Exception e) {
              logger.warn(String.format("Closing connection failed with error %s",e));
            }
            return;
        }

        // Create a new read request:
        // - Give the single item requested an alias name
        final PlcWriteRequest.Builder builder = connection.writeRequestBuilder();

        //If an array value is passed instead of a single value then convert to a String array
        if ((value.charAt(0) == '[') && (value.charAt(value.length() - 1) == ']')) {
            String[] values = value.substring(1,value.length() - 1).split(",");
            try {
                throw new IllegalAccessException("Adding Tag " + Arrays.toString(values));
            } catch (IllegalAccessException e) {
                String palle = "error: " + e;
                logger.info(palle);
            }
            builder.addTagAddress(tag, tag, values);
        } else {
            builder.addTagAddress(tag, tag, value);
        }

        PlcWriteRequest writeRequest = builder.build();

        try {
          writeRequest.execute().get();
        } catch (InterruptedException | ExecutionException e) {
          logger.warn(String.format("Failed %s",e));
          Thread.currentThread().interrupt();
        }

        try {
          connection.close();
        } catch (Exception e) {
          logger.warn(String.format("Closing Connection Failed with error %s",e));
        }
    }

    public PlcConnection prcd(){
        return new PlcConnection() {

            @Override
            public void connect() throws PlcConnectionException {
                throw new UnsupportedOperationException("Unimplemented method 'connect'");
            }

            @Override
            public boolean isConnected() {
                throw new UnsupportedOperationException("Unimplemented method 'isConnected'");
            }

            @Override
            public void close() throws Exception {
                throw new UnsupportedOperationException("Unimplemented method 'close'");
            }

            @Override
            public PlcConnectionMetadata getMetadata() {
                throw new UnsupportedOperationException("Unimplemented method 'getMetadata'");
            }

            @Override
            public CompletableFuture<Void> ping() {
                throw new UnsupportedOperationException("Unimplemented method 'ping'");
            }

            @Override
            public Builder readRequestBuilder() {
                throw new UnsupportedOperationException("Unimplemented method 'readRequestBuilder'");
            }

            @Override
            public org.apache.plc4x.java.api.messages.PlcWriteRequest.Builder writeRequestBuilder() {
                throw new UnsupportedOperationException("Unimplemented method 'writeRequestBuilder'");
            }

            @Override
            public org.apache.plc4x.java.api.messages.PlcSubscriptionRequest.Builder subscriptionRequestBuilder() {
                throw new UnsupportedOperationException("Unimplemented method 'subscriptionRequestBuilder'");
            }

            @Override
            public org.apache.plc4x.java.api.messages.PlcUnsubscriptionRequest.Builder unsubscriptionRequestBuilder() {
                throw new UnsupportedOperationException("Unimplemented method 'unsubscriptionRequestBuilder'");
            }

            @Override
            public org.apache.plc4x.java.api.messages.PlcBrowseRequest.Builder browseRequestBuilder() {
                throw new UnsupportedOperationException("Unimplemented method 'browseRequestBuilder'");
            }
            
        };
    }
}
