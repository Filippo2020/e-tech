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
package org.apache.plc4x.java.api.exceptions;

import org.junit.jupiter.api.Test;

import static org.hamcrest.core.IsEqual.equalTo;
import static org.hamcrest.MatcherAssert.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;

class PlcInvalidTagExceptionTest {

    @Test
    void simpleTagStringConstructor() {
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("Pattern");
        });
    }

    @Test
    void simpleTagStringAndPatternConstructor() {
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("Pattern");
        });
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("foo");
        });
    }

    @Test
    void simpleTagStringPatternAndReadableStringConstructor() {
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("Pattern");
        });
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("foo");
        });
        assertThrows(PlcInvalidTagException.class, () -> {
            throw new PlcInvalidTagException("readable");
        });
    }

    @Test
    void getTagToBeParsed() {
        PlcInvalidTagException ex = new PlcInvalidTagException("Pattern");
        assertThat(ex.getTagToBeParsed(), equalTo("Pattern"));
    }

}