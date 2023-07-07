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

// CALCommandType is an enum
type CALCommandType uint8

type ICALCommandType interface {
	utils.Serializable
}

const (
	CALCommandType_RESET           CALCommandType = 0x00
	CALCommandType_RECALL          CALCommandType = 0x01
	CALCommandType_IDENTIFY        CALCommandType = 0x02
	CALCommandType_GET_STATUS      CALCommandType = 0x03
	CALCommandType_WRITE           CALCommandType = 0x04
	CALCommandType_REPLY           CALCommandType = 0x0F
	CALCommandType_ACKNOWLEDGE     CALCommandType = 0x10
	CALCommandType_STATUS          CALCommandType = 0x11
	CALCommandType_STATUS_EXTENDED CALCommandType = 0x12
)

var CALCommandTypeValues []CALCommandType

func init() {
	_ = errors.New
	CALCommandTypeValues = []CALCommandType{
		CALCommandType_RESET,
		CALCommandType_RECALL,
		CALCommandType_IDENTIFY,
		CALCommandType_GET_STATUS,
		CALCommandType_WRITE,
		CALCommandType_REPLY,
		CALCommandType_ACKNOWLEDGE,
		CALCommandType_STATUS,
		CALCommandType_STATUS_EXTENDED,
	}
}

func CALCommandTypeByValue(value uint8) (enum CALCommandType, ok bool) {
	switch value {
	case 0x00:
		return CALCommandType_RESET, true
	case 0x01:
		return CALCommandType_RECALL, true
	case 0x02:
		return CALCommandType_IDENTIFY, true
	case 0x03:
		return CALCommandType_GET_STATUS, true
	case 0x04:
		return CALCommandType_WRITE, true
	case 0x0F:
		return CALCommandType_REPLY, true
	case 0x10:
		return CALCommandType_ACKNOWLEDGE, true
	case 0x11:
		return CALCommandType_STATUS, true
	case 0x12:
		return CALCommandType_STATUS_EXTENDED, true
	}
	return 0, false
}

func CALCommandTypeByName(value string) (enum CALCommandType, ok bool) {
	switch value {
	case "RESET":
		return CALCommandType_RESET, true
	case "RECALL":
		return CALCommandType_RECALL, true
	case "IDENTIFY":
		return CALCommandType_IDENTIFY, true
	case "GET_STATUS":
		return CALCommandType_GET_STATUS, true
	case "WRITE":
		return CALCommandType_WRITE, true
	case "REPLY":
		return CALCommandType_REPLY, true
	case "ACKNOWLEDGE":
		return CALCommandType_ACKNOWLEDGE, true
	case "STATUS":
		return CALCommandType_STATUS, true
	case "STATUS_EXTENDED":
		return CALCommandType_STATUS_EXTENDED, true
	}
	return 0, false
}

func CALCommandTypeKnows(value uint8) bool {
	for _, typeValue := range CALCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCALCommandType(structType interface{}) CALCommandType {
	castFunc := func(typ interface{}) CALCommandType {
		if sCALCommandType, ok := typ.(CALCommandType); ok {
			return sCALCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m CALCommandType) GetLengthInBits() uint16 {
	return 8
}

func (m CALCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALCommandTypeParse(theBytes []byte) (CALCommandType, error) {
	return CALCommandTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func CALCommandTypeParseWithBuffer(readBuffer utils.ReadBuffer) (CALCommandType, error) {
	val, err := readBuffer.ReadUint8("CALCommandType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading CALCommandType")
	}
	if enum, ok := CALCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return CALCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e CALCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e CALCommandType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("CALCommandType", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e CALCommandType) PLC4XEnumName() string {
	switch e {
	case CALCommandType_RESET:
		return "RESET"
	case CALCommandType_RECALL:
		return "RECALL"
	case CALCommandType_IDENTIFY:
		return "IDENTIFY"
	case CALCommandType_GET_STATUS:
		return "GET_STATUS"
	case CALCommandType_WRITE:
		return "WRITE"
	case CALCommandType_REPLY:
		return "REPLY"
	case CALCommandType_ACKNOWLEDGE:
		return "ACKNOWLEDGE"
	case CALCommandType_STATUS:
		return "STATUS"
	case CALCommandType_STATUS_EXTENDED:
		return "STATUS_EXTENDED"
	}
	return ""
}

func (e CALCommandType) String() string {
	return e.PLC4XEnumName()
}