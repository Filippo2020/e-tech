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
package org.apache.plc4x.java.modbus;

import org.apache.plc4x.java.modbus.base.tag.*;
import org.junit.jupiter.api.Assertions;
import org.junit.jupiter.api.Test;

class ModbusTagDataTypeTest {

    @Test
    void testHolding_DataType() {
        //Datatype, Length in Words
        String[][] datatypes = {{"INT","1"},
                                {"UINT","1"},
                                {"REAL","2"}};
        for(int i = 0; i< datatypes.length; i++){
            ModbusTagHoldingRegister holdingregister = ModbusTagHoldingRegister.of("holding-register:1:" + datatypes[i][0]);
            Assertions.assertEquals(datatypes[i][0], holdingregister.getDataType().name());
            Assertions.assertEquals(1, holdingregister.getNumberOfElements());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]) * 2, holdingregister.getLengthBytes());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]), holdingregister.getLengthWords());
        }
    }

    @Test
    void testInput_DataType() {
        //Datatype, Length in Words
        String[][] datatypes = {{"INT","1"},
                                {"UINT","1"},
                                {"REAL","2"}};
        for(int i = 0; i< datatypes.length; i++){
            ModbusTagInputRegister inputregister = ModbusTagInputRegister.of("input-register:1:" + datatypes[i][0]);
            Assertions.assertEquals(datatypes[i][0], inputregister.getDataType().name());
            Assertions.assertEquals(1, inputregister.getNumberOfElements());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]) * 2, inputregister.getLengthBytes());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]), inputregister.getLengthWords());
        }
    }

    @Test
    void testExtended_DataType() {
        //Datatype, Length in Words
        String[][] datatypes = {{"INT","1"},
                                {"UINT","1"},
                                {"DINT","2"},
                                {"REAL","2"}};
        for(int i = 0; i< datatypes.length; i++){
            ModbusTagExtendedRegister extendedregister = ModbusTagExtendedRegister.of("extended-register:1:" + datatypes[i][0]);
            Assertions.assertEquals(datatypes[i][0], extendedregister.getDataType().name());
            Assertions.assertEquals(1, extendedregister.getNumberOfElements());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]) * 2, extendedregister.getLengthBytes());
            Assertions.assertEquals(Integer.parseInt(datatypes[i][1]), extendedregister.getLengthWords());
        }
    }

    @Test
    void testCoil_DataType() {
        //Datatype, Length in Bytes
        String[][] datatypes = {{"BOOL","1"}};
        for(int i = 0; i< datatypes.length; i++){
            ModbusTagCoil coil = ModbusTagCoil.of("coil:1:" + datatypes[i][0]);
            Assertions.assertEquals(datatypes[i][0], coil.getDataType().name());
            Assertions.assertEquals(1, coil.getNumberOfElements());
        }
    }

    @Test
    void testDiscreteInput_DataType() {
        //Datatype, Length in Bytes
        String[][] datatypes = {{"BOOL","1"}};
        for(int i = 0; i< datatypes.length; i++){
            ModbusTagDiscreteInput discrete = ModbusTagDiscreteInput.of("discrete-input:1:" + datatypes[i][0]);
            Assertions.assertEquals(datatypes[i][0], discrete.getDataType().name());
            Assertions.assertEquals(1, discrete.getNumberOfElements());
        }
    }
}
