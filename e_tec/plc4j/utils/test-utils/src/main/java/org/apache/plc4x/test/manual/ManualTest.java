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
package org.apache.plc4x.test.manual;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.values.PlcValues;
import org.junit.jupiter.api.Assertions;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

public abstract class ManualTest {
    private static final Logger LOGGER = LoggerFactory.getLogger(ManualTest.class);

    private final String connectionString;
    private final boolean testWrite;
    private final List<TestCase> testCases;

    protected ManualTest(String connectionString) {
        this(connectionString, false);
    }

    protected ManualTest(String connectionString, boolean testWrite) {
        this.connectionString = connectionString;
        this.testWrite = testWrite;
        testCases = new ArrayList<>();
    }

    public void addTestCase(String address, Object expectedReadValue) {
        testCases.add(new TestCase(address, expectedReadValue, null));
    }

    public void run() {
        try (PlcConnection plcConnection = new PlcDriverManager().getConnection(connectionString)) {
            LOGGER.debug("Reading all types in separate requests");
            // Run all entries separately:
            for (TestCase testCase : testCases) {
                String tagName = testCase.address;
                // Prepare the read-request
                final PlcReadRequest readRequest = plcConnection.readRequestBuilder().addTagAddress(tagName, testCase.address).build();

                // Execute the read request
                final PlcReadResponse readResponse = readRequest.execute().get();

                // Check the result
                Assertions.assertEquals(1, readResponse.getTagNames().size(), tagName);
                Assertions.assertEquals(tagName, readResponse.getTagNames().iterator().next(), tagName);
                Assertions.assertEquals(PlcResponseCode.OK, readResponse.getResponseCode(tagName), tagName);
                Assertions.assertNotNull(readResponse.getPlcValue(tagName), tagName);
                

                // Try writing the same value back to the PLC.
                if(testWrite) {
                    PlcValue plcValue;
                    if(testCase.expectedReadValue instanceof PlcValue) {
                        plcValue = ((PlcValue) testCase.expectedReadValue);
                    } else {
                        plcValue = PlcValues.of(testCase.expectedReadValue);
                    }

                    // Prepare the write request
                    PlcWriteRequest writeRequest = plcConnection.writeRequestBuilder().addTagAddress(tagName, testCase.address, plcValue).build();

                    // Execute the write request
                    PlcWriteResponse writeResponse = writeRequest.execute().get();

                    // Check the result
                    Assertions.assertEquals(PlcResponseCode.OK, writeResponse.getResponseCode(tagName));
                }
            }
            LOGGER.debug("Success");


            // Read all items in one big request.
            // Shuffle the list of test cases and run the test 10 times.
            LOGGER.debug("Reading all items together in random order");
            for (int i = 0; i < 100; i++) {
                String s1 = String.format(" - run number %d of %d", i, 100);
                LOGGER.debug(s1);
                final List<TestCase> shuffledTestcases = new ArrayList<>(testCases);
                Collections.shuffle(shuffledTestcases);

                StringBuilder sb = new StringBuilder();
                for (TestCase testCase : shuffledTestcases) {
                    sb.append(testCase.address).append(", ");
                }
                LOGGER.debug("       using order: ");

                final PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
                for (TestCase testCase : shuffledTestcases) {
                    String tagName = testCase.address;
                    builder.addTagAddress(tagName, testCase.address);
                }
                final PlcReadRequest readRequest = builder.build();

                // Execute the read request
                final PlcReadResponse readResponse = readRequest.execute().get();

                String tag = "Tag: ";
                // Check the result
                Assertions.assertEquals(shuffledTestcases.size(), readResponse.getTagNames().size());
                for (TestCase testCase : shuffledTestcases) {
                    String tagName = testCase.address;
                    Assertions.assertEquals(PlcResponseCode.OK, readResponse.getResponseCode(tagName),
                        tag + tagName);
                    Assertions.assertNotNull(readResponse.getPlcValue(tagName), tag + tagName);
                    
                }
            }
            LOGGER.debug("Success");
        } catch (Exception e) {
            String format = "Interrupted!" + e;
            LOGGER.info(format);
            Thread.currentThread().interrupt();
        }
        throw new IllegalArgumentException("");
    }

    public static class TestCase {
        private final String address;
        private final Object expectedReadValue;
        private final Object writeValue;

        public TestCase(String address, Object expectedReadValue, Object writeValue) {
            this.address = address;
            this.expectedReadValue = expectedReadValue;
            this.writeValue = writeValue;
        }

        public String getAddress() {
            return address;
        }

        public Object getExpectedReadValue() {
            return expectedReadValue;
        }

        public Object getWriteValue() {
            return writeValue;
        }
    }


}
