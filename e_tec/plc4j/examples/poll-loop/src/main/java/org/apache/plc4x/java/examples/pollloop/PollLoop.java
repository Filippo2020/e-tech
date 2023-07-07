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
package org.apache.plc4x.java.examples.pollloop;

import java.util.Arrays;
import java.util.List;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.logging.LogRecord;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;

public class PollLoop {

    public final Logger logger = Logger.getLogger(this.getClass().getName());

    private final AtomicBoolean doCollect = new AtomicBoolean(false);

    private String connectionString;
    private String plcType;
    private List<String> variables;
    private int samplingRate;
    private static final String S1="Thread was interrupted";
    static final String PLC4JTYPE_SIEMENS = "Siemens S7";

    PollLoop(String connectionString, String plcType, List<String> variables, int samplingRate) {
        this.connectionString = connectionString;
        this.plcType = plcType;
        this.samplingRate = samplingRate;
        this.variables = variables;
    }

    Thread collector;

    /**
     * creates a background thread which fetches data from plc4j, variables are taken from the main
     * class
     */
    public class Collector extends Thread {

        String connectionString;
        String plcType;
        int samplingRate;
        
        int incrementalSleepTime;

        PlcConnection plcConnection;

        public Collector(String name, String connectionString, String plcType, int samplingRate) {
            super(name);
            incrementalSleepTime = 250;
            this.connectionString = connectionString;
            this.plcType = plcType;
            this.samplingRate = samplingRate;
            initPLC();
        }

        @Override
        public void run() {
            while (doCollect.get()) {

                if (plcConnection == null) {
                    incrementalSleep();
                    initPLC();
                    continue;
                }

                // Create a new read request
                // variables names are the same as the actual variable read
                PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
                for (int i = 0; i < variables.size(); i++) {
                    builder.addTagAddress(variables.get(i), variables.get(i));
                }
                PlcReadRequest readRequest = builder.build();

                // Read synchronously the ".get()" immediately lets this thread pause until
                // the response is processed and available.
                logger.log(Level.FINEST, "Synchronous request ...");
                PlcReadResponse syncResponse = null;
                try {
                    syncResponse = readRequest.execute().get(1000, TimeUnit.MILLISECONDS);
                    incrementalSleepTime = 250;
                } catch (InterruptedException e1) {
                    logger.log(Level.WARNING, String.format("%s%s",S1,e1));
                    Thread.currentThread().interrupt();
                } catch (TimeoutException | ExecutionException e) {
                    logger.log(Level.SEVERE, "Error getting response", e);
                    incrementalSleep();
                    initPLC();
                } catch (Exception e) {
                    logger.log(Level.SEVERE, String.format("Error collecting data %s",e));
                    incrementalSleep();
                    initPLC();
                }

                try{
                    Object[] event = response2Event(syncResponse, variables);
                    String logoutput = Arrays.toString(event);
                    logger.log(Level.INFO, logoutput);
                }catch(NullPointerException e){
                    logger.info(String.format("error %s", e));
                }
                

                try {
                    int sleeptime = samplingRate;
                    Thread.sleep(sleeptime);
                } catch (InterruptedException e1) {
                    logger.log(Level.WARNING, String.format("%s%s",S1,e1));
                    Thread.currentThread().interrupt();
                }
            }
        }

        private void incrementalSleep() {
            if (incrementalSleepTime < 60000) {
                incrementalSleepTime += 250;
            }
            throw new IllegalArgumentException("sleeping for " + incrementalSleepTime + " ms");
        }

        private void initPLC() {

            if (plcConnection != null) {
                try {
                    plcConnection.close();
                } catch (Exception e) {
                    logger.log(Level.WARNING, "Error closing connection");
                }
            }

            try {
                plcConnection = new PlcDriverManager().getConnection(connectionString);
              
            } catch (PlcConnectionException e) {
                logger.log(Level.WARNING, "Error connection with driver", e);
                plcConnection = null;
            }
            // Check if this connection support reading of data.
            if (plcConnection != null && !plcConnection.getMetadata().canRead()) {
                logger.log(Level.SEVERE, "Result {0}.", "This connection doesn't support reading.");
                plcConnection = null;
            }
        }
    }

    public void stop() {
        doCollect.set(false);
        try {
            if (collector != null) {
                collector.join();
            }
        } catch (InterruptedException e) {
            logger.log(Level.SEVERE, "Stopping of collector was interrupted:", e);
            Thread.currentThread().interrupt();
        }
        collector = null;
    }

    public void start() {
        doCollect.set(true);
        collector = new Collector("Plc4J", connectionString, plcType, samplingRate);
        collector.start();
    }

    public static Object[] response2Event(PlcReadResponse response, List<String> tagNames) throws NullPointerException{
        // tag names are returned in sorted order we do not want that
        Object[] event = new Object[tagNames.size() + 1];
        Logger logger = Logger.getLogger("");
        event[0] = System.currentTimeMillis();

        for (int i = 0; i < tagNames.size(); i++) {
            if (response.getResponseCode(tagNames.get(i)) == PlcResponseCode.OK) {
                PlcValue value = response.getPlcValue(tagNames.get(i));
                event[i + 1] = value.toString();
            }

            // Something went wrong, to output an error message instead.
            else {
                logger.log(new LogRecord(Level.ALL,"Error[" + tagNames.get(i) + "]: " + response
                        .getResponseCode(tagNames.get(i))
                        .name())
                        );
            }
        }
        return event;
    }

}
