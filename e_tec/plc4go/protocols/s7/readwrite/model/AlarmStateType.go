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

package model

import (
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// AlarmStateType is an enum
type AlarmStateType uint8

type IAlarmStateType interface {
	utils.Serializable
}

const (
	AlarmStateType_SCAN_ABORT       AlarmStateType = 0x00
	AlarmStateType_SCAN_INITIATE    AlarmStateType = 0x01
	AlarmStateType_ALARM_ABORT      AlarmStateType = 0x04
	AlarmStateType_ALARM_INITIATE   AlarmStateType = 0x05
	AlarmStateType_ALARM_S_ABORT    AlarmStateType = 0x08
	AlarmStateType_ALARM_S_INITIATE AlarmStateType = 0x09
)

var AlarmStateTypeValues []AlarmStateType

func init() {
	_ = errors.New
	AlarmStateTypeValues = []AlarmStateType{
		AlarmStateType_SCAN_ABORT,
		AlarmStateType_SCAN_INITIATE,
		AlarmStateType_ALARM_ABORT,
		AlarmStateType_ALARM_INITIATE,
		AlarmStateType_ALARM_S_ABORT,
		AlarmStateType_ALARM_S_INITIATE,
	}
}

func AlarmStateTypeByValue(value uint8) (enum AlarmStateType, ok bool) {
	switch value {
	case 0x00:
		return AlarmStateType_SCAN_ABORT, true
	case 0x01:
		return AlarmStateType_SCAN_INITIATE, true
	case 0x04:
		return AlarmStateType_ALARM_ABORT, true
	case 0x05:
		return AlarmStateType_ALARM_INITIATE, true
	case 0x08:
		return AlarmStateType_ALARM_S_ABORT, true
	case 0x09:
		return AlarmStateType_ALARM_S_INITIATE, true
	}
	return 0, false
}

func AlarmStateTypeByName(value string) (enum AlarmStateType, ok bool) {
	switch value {
	case "SCAN_ABORT":
		return AlarmStateType_SCAN_ABORT, true
	case "SCAN_INITIATE":
		return AlarmStateType_SCAN_INITIATE, true
	case "ALARM_ABORT":
		return AlarmStateType_ALARM_ABORT, true
	case "ALARM_INITIATE":
		return AlarmStateType_ALARM_INITIATE, true
	case "ALARM_S_ABORT":
		return AlarmStateType_ALARM_S_ABORT, true
	case "ALARM_S_INITIATE":
		return AlarmStateType_ALARM_S_INITIATE, true
	}
	return 0, false
}

func AlarmStateTypeKnows(value uint8) bool {
	for _, typeValue := range AlarmStateTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAlarmStateType(structType interface{}) AlarmStateType {
	castFunc := func(typ interface{}) AlarmStateType {
		if sAlarmStateType, ok := typ.(AlarmStateType); ok {
			return sAlarmStateType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AlarmStateType) GetLengthInBits() uint16 {
	return 8
}

func (m AlarmStateType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AlarmStateTypeParse(theBytes []byte) (AlarmStateType, error) {
	return AlarmStateTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AlarmStateTypeParseWithBuffer(readBuffer utils.ReadBuffer) (AlarmStateType, error) {
	val, err := readBuffer.ReadUint8("AlarmStateType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AlarmStateType")
	}
	if enum, ok := AlarmStateTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AlarmStateType(val), nil
	} else {
		return enum, nil
	}
}

func (e AlarmStateType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AlarmStateType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("AlarmStateType", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AlarmStateType) PLC4XEnumName() string {
	switch e {
	case AlarmStateType_SCAN_ABORT:
		return "SCAN_ABORT"
	case AlarmStateType_SCAN_INITIATE:
		return "SCAN_INITIATE"
	case AlarmStateType_ALARM_ABORT:
		return "ALARM_ABORT"
	case AlarmStateType_ALARM_INITIATE:
		return "ALARM_INITIATE"
	case AlarmStateType_ALARM_S_ABORT:
		return "ALARM_S_ABORT"
	case AlarmStateType_ALARM_S_INITIATE:
		return "ALARM_S_INITIATE"
	}
	return ""
}

func (e AlarmStateType) String() string {
	return e.PLC4XEnumName()
}
