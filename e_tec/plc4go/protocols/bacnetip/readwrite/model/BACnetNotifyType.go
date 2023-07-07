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

// BACnetNotifyType is an enum
type BACnetNotifyType uint8

type IBACnetNotifyType interface {
	utils.Serializable
}

const (
	BACnetNotifyType_ALARM            BACnetNotifyType = 0x0
	BACnetNotifyType_EVENT            BACnetNotifyType = 0x1
	BACnetNotifyType_ACK_NOTIFICATION BACnetNotifyType = 0x2
)

var BACnetNotifyTypeValues []BACnetNotifyType

func init() {
	_ = errors.New
	BACnetNotifyTypeValues = []BACnetNotifyType{
		BACnetNotifyType_ALARM,
		BACnetNotifyType_EVENT,
		BACnetNotifyType_ACK_NOTIFICATION,
	}
}

func BACnetNotifyTypeByValue(value uint8) (enum BACnetNotifyType, ok bool) {
	switch value {
	case 0x0:
		return BACnetNotifyType_ALARM, true
	case 0x1:
		return BACnetNotifyType_EVENT, true
	case 0x2:
		return BACnetNotifyType_ACK_NOTIFICATION, true
	}
	return 0, false
}

func BACnetNotifyTypeByName(value string) (enum BACnetNotifyType, ok bool) {
	switch value {
	case "ALARM":
		return BACnetNotifyType_ALARM, true
	case "EVENT":
		return BACnetNotifyType_EVENT, true
	case "ACK_NOTIFICATION":
		return BACnetNotifyType_ACK_NOTIFICATION, true
	}
	return 0, false
}

func BACnetNotifyTypeKnows(value uint8) bool {
	for _, typeValue := range BACnetNotifyTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetNotifyType(structType interface{}) BACnetNotifyType {
	castFunc := func(typ interface{}) BACnetNotifyType {
		if sBACnetNotifyType, ok := typ.(BACnetNotifyType); ok {
			return sBACnetNotifyType
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetNotifyType) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetNotifyType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotifyTypeParse(theBytes []byte) (BACnetNotifyType, error) {
	return BACnetNotifyTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetNotifyTypeParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetNotifyType, error) {
	val, err := readBuffer.ReadUint8("BACnetNotifyType", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetNotifyType")
	}
	if enum, ok := BACnetNotifyTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetNotifyType(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetNotifyType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetNotifyType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetNotifyType", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetNotifyType) PLC4XEnumName() string {
	switch e {
	case BACnetNotifyType_ALARM:
		return "ALARM"
	case BACnetNotifyType_EVENT:
		return "EVENT"
	case BACnetNotifyType_ACK_NOTIFICATION:
		return "ACK_NOTIFICATION"
	}
	return ""
}

func (e BACnetNotifyType) String() string {
	return e.PLC4XEnumName()
}
