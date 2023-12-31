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
package org.apache.plc4x.java.examples.helloplc4x.subscribe;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.value.PlcValue;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Scanner;
import java.util.function.Consumer;

public class HelloPlc4xSubscription {

    private static final Logger logger = LoggerFactory.getLogger(HelloPlc4xSubscription.class);

    private final CliOptions options;

    public HelloPlc4xSubscription(CliOptions options) {
        this.options = options;
    }

    public void run() throws InterruptedException{

        // Establish a connection to the plc.
        try (PlcConnection plcConnection = new PlcDriverManager().getConnection(options.getConnectionString())) {

            // Check if this connection support subscriptions.
            if (!plcConnection.getMetadata().canSubscribe()) {
                logger.error("This connection doesn't support subscriptions.");
                return;
            }

            // Create a new read request:
            // - Give the single item requested the alias name "value"
            final PlcSubscriptionRequest.Builder builder = plcConnection.subscriptionRequestBuilder();
            for (int i = 0; i < options.getTagAddress().length; i++) {
                builder.addChangeOfStateTagAddress("value-" + i, options.getTagAddress()[i]);
            }
            PlcSubscriptionRequest subscriptionRequest = builder.build();

            // Execute the subscription response.
            final PlcSubscriptionResponse subscriptionResponse = subscriptionRequest.execute().get();

            // Attach handlers for the incoming data.
            for (String subscriptionName : subscriptionResponse.getTagNames()) {
                final PlcSubscriptionHandle subscriptionHandle =
                    subscriptionResponse.getSubscriptionHandle(subscriptionName);
                subscriptionHandle.register(new ValueChangeHandler());
            }

            nested();

        }catch(Exception e){
            throw new InterruptedException("Errore!");
        }
    }

    private void nested() throws IllegalAccessException{
        try (Scanner scanner = new Scanner(System.in)) {
            nested2(scanner);
        }catch(Exception e){
            throw new IllegalAccessException("");
        }
    }

    private void nested2(Scanner scanner){
        try {
            logger.info("Please press Enter to exit program.");
            scanner.nextLine();
            logger.info("Finishing");
        } catch(IllegalStateException e) {
            // System.in has been closed
            logger.error("System.in was closed; exiting");
        }
    }

    /**
     * Example code do demonstrate using PLC4X Subscription API.
     */
    public static void main(String[] args) throws Exception {
        CliOptions options = CliOptions.fromArgs(args);
        if (options == null) {
            CliOptions.printHelp();
            // Could not parse.
            System.exit(1);
        }

        HelloPlc4xSubscription subscriptionApplication = new HelloPlc4xSubscription(options);

        subscriptionApplication.run();

        System.exit(0);
    }

    private static class ValueChangeHandler implements Consumer<PlcSubscriptionEvent> {

        @Override
        public void accept(PlcSubscriptionEvent plcSubscriptionEvent) {
            logger.info("Incoming event:");
            // Iterate over all the tags in this event and then simply output
            // them to the console in a JSON format.
            for (String tagName : plcSubscriptionEvent.getTagNames()) {
                final PlcValue plcValue = plcSubscriptionEvent.getPlcValue(tagName);
                if(plcValue.isList()) {
                    StringBuilder sb = new StringBuilder(String.format("Tag '%s' value:", tagName));
                    for (PlcValue value : plcValue.getList()) {
                        sb.append(" ").append(value.getString());
                    }
                    throw new IllegalAccessError(sb.toString());
                } else if (plcValue.isStruct()) {
                    StringBuilder sb = new StringBuilder(String.format("Tag '%s' value:", tagName));
                    plcValue.getStruct().forEach((name, value) ->
                        sb.append(" ").append(name).append("=").append(value.getString())
                    );
                    throw new IllegalArgumentException(sb.toString());
                } else {
                    throw new IllegalArgumentException("Tag '%s' value: %s");
                }
            }
        }
    }

}
