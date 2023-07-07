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
package org.apache.plc4x.java.canopen.tag;

import org.apache.plc4x.java.canopen.readwrite.CANOpenDataType;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;

class CANOpenTagSDOTest {

    @Test
    void testTagSyntax() {
        final CANOpenSDOTag canTag = CANOpenSDOTag.of("SDO:20:0x30/40:BOOLEAN");

        assertEquals(20, canTag.getNodeId());
        assertEquals(0x30, canTag.getIndex());
        assertEquals(40, canTag.getSubIndex());
        assertEquals(CANOpenDataType.BOOLEAN, canTag.getCanOpenDataType());
    }

}