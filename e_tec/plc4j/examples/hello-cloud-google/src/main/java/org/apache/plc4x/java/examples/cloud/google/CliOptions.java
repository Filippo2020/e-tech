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
package org.apache.plc4x.java.examples.cloud.google;

import org.apache.commons.cli.*;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/** Command line options for the MQTT example. */
public class CliOptions {

    private static final Logger logger = LoggerFactory.getLogger(CliOptions.class);

    private static Options options;

    private final String projectId;
    private final String registryId;
    private final String deviceId;
    private final String privateKeyFile;
    private final String algorithm;
    private final String cloudRegion;
    private final int tokenExpMins;
    private final String mqttBridgeHostname;
    private final short mqttBridgePort;
    private final String messageType;
    private static final String S1="token-exp-minutes";
    private static final String S2="mqtt-bridge-port";
    private static final String S3="cloud-region";
    private static final String S4="mqtt-bridge-hostname";
    private static final String S5="message-type";
    public static CliOptions fromArgs(String[] args) {
        options = new Options();
        // Required arguments
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt("project-id")
                .hasArg()
                .desc("GCP cloud project name.")
                .required()
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt("registry-id")
                .hasArg()
                .desc("Cloud IoT Core registry id.")
                .required()
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt("device-id")
                .hasArg()
                .desc("Cloud IoT Core device id.")
                .required()
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt("private-key-file")
                .hasArg()
                .desc("Path to private key file.")
                .required()
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt("algorithm")
                .hasArg()
                .desc("Encryption algorithm to use to generate the JWT. Either 'RS256' or 'ES256'.")
                .required()
                .build());

        // Optional arguments.
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt(S3)
                .hasArg()
                .desc("GCP cloud region.")
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt(S4)
                .hasArg()
                .desc("MQTT bridge hostname.")
                .build());
        options.addOption(
            Option.builder()
                .type(Number.class)
                .longOpt(S1)
                .hasArg()
                .desc("Minutes to JWT token refresh (token expiration time).")
                .build());
        options.addOption(
            Option.builder()
                .type(Number.class)
                .longOpt(S2)
                .hasArg()
                .desc("MQTT bridge port.")
                .build());
        options.addOption(
            Option.builder()
                .type(String.class)
                .longOpt(S5)
                .hasArg()
                .desc("Indicates whether the message is a telemetry event or a device state message")
                .build());

        CommandLineParser parser = new DefaultParser();
        CommandLine commandLine;
        try {
            commandLine = parser.parse(options, args);

            String projectId = commandLine.getOptionValue("project-id");
            String registryId = commandLine.getOptionValue("registry-id");
            String deviceId = commandLine.getOptionValue("device-id");
            String privateKeyFile = commandLine.getOptionValue("private-key-file");
            String algorithm = commandLine.getOptionValue("algorithm");
            String cloudRegion = "europe-west1";
            if (commandLine.hasOption(S3)) {
                cloudRegion = commandLine.getOptionValue(S3);
            }
            int tokenExpMins = 20;
            if (commandLine.hasOption(S1)) {
                tokenExpMins =
                    ((Number) commandLine.getParsedOptionValue(S1)).intValue();
            }
            String mqttBridgeHostname = "mqtt.googleapis.com";
            if (commandLine.hasOption(S4)) {
                mqttBridgeHostname = commandLine.getOptionValue(S4);
            }
            short mqttBridgePort = 8883;
            if (commandLine.hasOption(S2)) {
                mqttBridgePort =
                    ((Number) commandLine.getParsedOptionValue(S2)).shortValue();
            }
            String messageType = "event";
            if (commandLine.hasOption(S5)) {
                messageType = commandLine.getOptionValue(S5);
            }

            return new CliOptions(projectId, registryId, deviceId, privateKeyFile, algorithm, cloudRegion, tokenExpMins,
                mqttBridgeHostname, mqttBridgePort, messageType);
        } catch (ParseException e) {
            logger.debug(e.getMessage());
            return null;
        }
    }

    public static void printHelp() {
        HelpFormatter formatter = new HelpFormatter();
        formatter.printHelp("S7PlcToGoogleIoTCoreSample", options);
    }

    public CliOptions(String projectId, String registryId, String deviceId, String privateKeyFile, String algorithm,
                      String cloudRegion, int tokenExpMins, String mqttBridgeHostname, short mqttBridgePort,
                      String messageType) {
        this.projectId = projectId;
        this.registryId = registryId;
        this.deviceId = deviceId;
        this.privateKeyFile = privateKeyFile;
        this.algorithm = algorithm;
        this.cloudRegion = cloudRegion;
        this.tokenExpMins = tokenExpMins;
        this.mqttBridgeHostname = mqttBridgeHostname;
        this.mqttBridgePort = mqttBridgePort;
        this.messageType = messageType;
    }

    public String getProjectId() {
        return projectId;
    }

    public String getRegistryId() {
        return registryId;
    }

    public String getDeviceId() {
        return deviceId;
    }

    public String getPrivateKeyFile() {
        return privateKeyFile;
    }

    public String getAlgorithm() {
        return algorithm;
    }

    public String getCloudRegion() {
        return cloudRegion;
    }

    public int getTokenExpMins() {
        return tokenExpMins;
    }

    public String getMqttBridgeHostname() {
        return mqttBridgeHostname;
    }

    public short getMqttBridgePort() {
        return mqttBridgePort;
    }

    public String getMessageType() {
        return messageType;
    }

}