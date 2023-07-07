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
package org.apache.plc4x.kafka;

import org.apache.kafka.common.config.AbstractConfig;
import org.apache.kafka.common.config.ConfigDef;
import org.apache.kafka.connect.data.*;

import org.apache.kafka.connect.sink.SinkRecord;
import org.apache.kafka.connect.sink.SinkTask;
import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.utils.connectionpool2.PooledDriverManager;
import org.apache.plc4x.kafka.config.Constants;
import org.apache.plc4x.kafka.util.VersionUtil;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.*;

/**
 * Source Connector Task polling the data source at a given rate.
 * A timer thread is scheduled which sets the fetch flag to true every rate milliseconds.
 * When poll() is invoked, the calling thread waits until the fetch flag is set for WAIT_LIMIT_MILLIS.
 * If the flag does not become true, the method returns null, otherwise a fetch is performed.
 */
public class Plc4xSinkTask extends SinkTask {

    private static final Logger log = LoggerFactory.getLogger(Plc4xSinkTask.class);

    /*
     * Config of the task.
     */
    static final String CONNECTION_NAME_CONFIG = "connection-name";
    private static final String CONNECTION_NAME_STRING_DOC = "Connection Name";

    static final String PLC4X_CONNECTION_STRING_CONFIG = "connectionString";
    private static final String PLC4X_CONNECTION_STRING_DOC = "PLC4X Connection String";

    static final String PLC4X_TOPIC_CONFIG = "topic";
    private static final String PLC4X_TOPIC_DOC = "Task Topic";

    private static final String PLC4X_RETRIES_CONFIG = "retries";
    private static final String PLC4X_RETRIES_DOC = "Number of times to retry after failed write";

    private static final String PLC4X_TIMEOUT_CONFIG = "timeout";
    private static final String PLC4X_TIMEOUT_DOC = "Time between retries";

    // Syntax for the queries: {tag-alias}#{tag-address}:{tag-alias}#{tag-address}...,{topic}:{rate}:....
    static final String QUERIES_CONFIG = "queries";
    private static final String QUERIES_DOC = "Tags to be sent to the PLC";

    private static final ConfigDef CONFIG_DEF = new ConfigDef()
        .define(CONNECTION_NAME_CONFIG,
                ConfigDef.Type.STRING,
                ConfigDef.Importance.HIGH,
                CONNECTION_NAME_STRING_DOC)
        .define(PLC4X_CONNECTION_STRING_CONFIG,
                ConfigDef.Type.STRING,
                ConfigDef.Importance.HIGH,
                PLC4X_CONNECTION_STRING_DOC)
        .define(PLC4X_TOPIC_CONFIG,
                ConfigDef.Type.STRING,
                ConfigDef.Importance.HIGH,
                PLC4X_TOPIC_DOC)
        .define(PLC4X_RETRIES_CONFIG,
                ConfigDef.Type.INT,
                ConfigDef.Importance.HIGH,
                PLC4X_RETRIES_DOC)
        .define(PLC4X_TIMEOUT_CONFIG,
                ConfigDef.Type.INT,
                ConfigDef.Importance.HIGH,
                PLC4X_TIMEOUT_DOC)
        .define(QUERIES_CONFIG,
                ConfigDef.Type.STRING,
                ConfigDef.Importance.HIGH,
                QUERIES_DOC);

    /*
     * Configuration of the output.
     */

    @Override
    public String version() {
        return VersionUtil.getVersion();
    }

    private PlcDriverManager driverManager;
    private String plc4xConnectionString;
    private String plc4xTopic;
    private Integer plc4xRetries;

    private Integer remainingRetries;

    @Override
    public void start(Map<String, String> props) {
        Map<String, String> tags;
        AbstractConfig config;
        config = new AbstractConfig(CONFIG_DEF, props);
        plc4xConnectionString = config.getString(PLC4X_CONNECTION_STRING_CONFIG);
        plc4xTopic = config.getString(PLC4X_TOPIC_CONFIG);
        plc4xRetries = config.getInt(PLC4X_RETRIES_CONFIG);
        remainingRetries = plc4xRetries;

        String queries = config.getString(QUERIES_CONFIG);
        tags = new HashMap<>();

        String[] tagsConfigSegments = queries.split("\\|");
        for (String tagsConfigSegment : tagsConfigSegments) {
            String[] tagSegments = tagsConfigSegment.split("#");
            if (tagSegments.length != 2) {
                log.info("Error in tag configuration. " +
                        "The tag segment expects a format {tag-alias}#{tag-address}, but got '%s'",
                    tagsConfigSegment);
                continue;
            }
            String tagAlias = tagSegments[0];
            String tagAddress = tagSegments[1];

            tags.put(tagAlias, tagAddress);
        }

        log.info("Creating Pooled PLC4x driver manager");
        driverManager = new PooledDriverManager();
    }

    @Override
    public void stop() {
        synchronized (this) {
            notifyAll(); // wake up thread waiting in awaitFetch
        }
    }

    @Override
    public void put(Collection<SinkRecord> records) {
        if (records.isEmpty()) {
            return;
        }

        PlcConnection connection = null;
        try {
            connection = driverManager.getConnection(plc4xConnectionString);
        } catch (PlcConnectionException e) {
            log.warn("Failed to Open Connection {}", plc4xConnectionString);
            remainingRetries--;
            
            log.warn("Failed to write after {} retries", plc4xRetries);
            return;
        }

        for (SinkRecord r: records) {
            Struct recordz = (Struct) r.value();
            String topic = r.topic();

            Struct plcTags = recordz.getStruct(Constants.TAGS_CONFIG);
            Schema plcTagsSchema = plcTags.schema();

            for (Field plcTag : plcTagsSchema.fields()) {
                String tagName = plcTag.name();
                Object value = plcTags.get(tagName);
                if (value != null &&  (!topic.equals(plc4xTopic) || plc4xTopic.equals(""))) {
                        log.debug("Ignoring write request received on wrong topic");
                     
                }

            }
        }

        

        try {
            connection.close();
        } catch (Exception e) {
            log.warn("Failed to Close {}", plc4xConnectionString);
        }

        remainingRetries = plc4xRetries;
      
    }
}
