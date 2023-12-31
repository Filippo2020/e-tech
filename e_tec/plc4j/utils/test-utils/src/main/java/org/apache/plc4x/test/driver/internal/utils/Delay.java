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
package org.apache.plc4x.test.driver.internal.utils;

import org.apache.plc4x.test.driver.exceptions.DriverTestsuiteException;

import java.util.concurrent.TimeUnit;

public class Delay {
    private Delay() {
        throw new IllegalStateException("Utility class");
      }
    

    // replace with signals
    public static void shortDelay() throws DriverTestsuiteException {
        delay(23);
    }

    // replace with signals
    public static void delay(int milliseconds) throws DriverTestsuiteException {
        try {
            TimeUnit.MILLISECONDS.sleep(milliseconds);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new DriverTestsuiteException("Interrupted during delay.");
        }
    }
}
