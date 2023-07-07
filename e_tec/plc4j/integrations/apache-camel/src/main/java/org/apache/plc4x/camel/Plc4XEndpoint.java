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
package org.apache.plc4x.camel;

import org.apache.camel.*;
import org.apache.camel.support.DefaultEndpoint;
import org.apache.camel.spi.Metadata;
import org.apache.camel.spi.UriEndpoint;
import org.apache.camel.spi.UriParam;
import org.apache.camel.spi.UriPath;
import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.utils.connectionpool.PooledPlcDriverManager;

import java.util.Map;
import java.util.Objects;

@UriEndpoint(scheme = "plc4x", title = "PLC4X", syntax = "plc4x:driver", label = "plc4x")
public class Plc4XEndpoint extends DefaultEndpoint {
    private static final String S1="plc4x:/?/?";

    @UriPath
    @Metadata(required = true)
    private String driver;

    @UriParam
    private Map<String, Object> tags;

    @UriParam
    private String trigger;

    @UriParam
    private int period;

    public int getPeriod() {
        return period;
    }

    public void setPeriod(int period) {
        this.period = period;
    }

    private PlcDriverManager plcDriverManager;
    private PlcConnection connection;
    private String uri;

    public String getUri() {
        return uri;
    }

    public String getTrigger() {
        return trigger;
    }

    public void setTrigger(String trigger) {
        this.trigger = trigger;
        plcDriverManager = new PooledPlcDriverManager();
        String plc4xURI = uri.replaceFirst(S1, "");
        // is this mutation really intentional
        uri = plc4xURI;
        try {
            connection = plcDriverManager.getConnection(plc4xURI);
        } catch (PlcConnectionException e) {
            throw new PlcRuntimeException(e);
        }
    }

    public Plc4XEndpoint(String endpointUri, Component component) throws PlcConnectionException {
        super(endpointUri, component);
        this.plcDriverManager = new PlcDriverManager();
        //Here we establish the connection in the endpoint, as it is created once during the context
        // to avoid disconnecting and reconnecting for every request
        this.uri = endpointUri.replaceFirst(S1, "");
        this.connection = plcDriverManager.getConnection(this.uri);
    }

    public PlcConnection getConnection() {
        return connection;
    }

    @Override
    public void setProperties(Object bean, Map<String, Object> parameters) {
      // document why this method is empty
    }

    @Override
    public Producer createProducer() throws Exception {
        //Checking if connection is still up and reconnecting if not
        if (!connection.isConnected()) {
            connection = plcDriverManager.getConnection(uri.replaceFirst(S1, ""));
        }
        return new Plc4XProducer(this);
    }

    @Override
    public Consumer createConsumer(Processor processor) throws Exception {
        //Checking if connection is still up and reconnecting if not
        if (!connection.isConnected()) {
            connection = plcDriverManager.getConnection(uri.replaceFirst(S1, ""));
        }
        return new Plc4XConsumer(this, processor);
    }

    @Override
    public boolean isSingleton() {
        return true;
    }

    public PlcDriverManager getPlcDriverManager() {
        return plcDriverManager;
    }

    public String getDriver() {
        return driver;
    }

    public void setDriver(String driver) {
        this.driver = driver;
    }

    public Map<String, Object> getTags() {
        return tags;
    }

    public void setTags(Map<String, Object> tags) {
        this.tags = tags;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) {
            return true;
        }
        if (!(o instanceof Plc4XEndpoint)) {
            return false;
        }
        if (!super.equals(o)) {
            return false;
        }
        Plc4XEndpoint that = (Plc4XEndpoint) o;
        return Objects.equals(getDriver(), that.getDriver()) &&
            Objects.equals(getTags(), that.getTags()) &&
            Objects.equals(getPlcDriverManager(), that.getPlcDriverManager());
    }

    @Override
    public int hashCode() {
        return Objects.hash(super.hashCode(), getDriver(), getTags(), getPlcDriverManager());
    }

    @Override
    public void doStop() throws Exception {
        //Shutting down the connection when leaving the Context
        if (connection != null && connection.isConnected()) {
            connection.close();
        }
    }

}
