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

// ApduType is an enum
type ApduType uint8

type IApduType interface {
	utils.Serializable
}

const (
	ApduType_CONFIRMED_REQUEST_PDU   ApduType = 0x0
	ApduType_UNCONFIRMED_REQUEST_PDU ApduType = 0x1
	ApduType_SIMPLE_ACK_PDU          ApduType = 0x2
	ApduType_COMPLEX_ACK_PDU         ApduType = 0x3
	ApduType_SEGMENT_ACK_PDU         ApduType = 0x4
	ApduType_ERROR_PDU               ApduType = 0x5
	ApduType_REJECT_PDU              ApduType = 0x6
	ApduType_ABORT_PDU               ApduType = 0x7
	ApduType_APDU_UNKNOWN_8          ApduType = 0x8
	ApduType_APDU_UNKNOWN_9          ApduType = 0x9
	ApduType_APDU_UNKNOWN_A          ApduType = 0xA
	ApduType_APDU_UNKNOWN_B          ApduType = 0xB
	ApduType_APDU_UNKNOWN_C          ApduType = 0xC
	ApduType_APDU_UNKNOWN_D          ApduType = 0xD
	ApduType_APDU_UNKNOWN_E          ApduType = 0xE
	ApduType_APDU_UNKNOWN_F          ApduType = 0xF
)

var ApduTypeValues []ApduType

func init() {
	_ = errors.New
	ApduTypeValues = []ApduType{
		ApduType_CONFIRMED_REQUEST_PDU,
		ApduType_UNCONFIRMED_REQUEST_PDU,
		ApduType_SIMPLE_ACK_PDU,
		ApduType_COMPLEX_ACK_PDU,
		ApduType_SEGMENT_ACK_PDU,
		ApduType_ERROR_PDU,
		ApduType_REJECT_PDU,
		ApduType_ABORT_PDU,
		ApduType_APDU_UNKNOWN_8,
		ApduType_APDU_UNKNOWN_9,
		ApduType_APDU_UNKNOWN_A,
		ApduType_APDU_UNKNOWN_B,
		ApduType_APDU_UNKNOWN_C,
		ApduType_APDU_UNKNOWN_D,
		ApduType_APDU_UNKNOWN_E,
		ApduType_APDU_UNKNOWN_F,
	}
}

func ApduTypeByValue(value uint8) (enum ApduType, ok bool) {
	switch value {
	case 0x0:
		return ApduType_CONFIRMED_REQUEST_PDU, true
	case 0x1:
		return ApduType_UNCONFIRMED_REQUEST_PDU, true
	case 0x2:
		return ApduType_SIMPLE_ACK_PDU, true
	case 0x3:
		return ApduType_COMPLEX_ACK_PDU, true
	case 0x4:
		return ApduType_SEGMENT_ACK_PDU, true
	case 0x5:
		return ApduType_ERROR_PDU, true
	case 0x6:
		return ApduType_REJECT_PDU, true
	case 0x7:
		return ApduType_ABORT_PDU, true
	case 0x8:
		return ApduType_APDU_UNKNOWN_8, true
	case 0x9:
		return ApduType_APDU_UNKNOWN_9, true
	case 0xA:
		return ApduType_APDU_UNKNOWN_A, true
	case 0xB:
		return ApduType_APDU_UNKNOWN_B, true
	case 0xC:
		return ApduType_APDU_UNKNOWN_C, true
	case 0xD:
		return ApduType_APDU_UNKNOWN_D, true
	case 0xE:
		return ApduType_APDU_UNKNOWN_E, true
	case 0xF:
		return ApduType_APDU_UNKNOWN_F, true
	}
	return 0, false
}

func ApduTypeByName(value string) (enum ApduType, ok bool) {
	switch value {
	case "CONFIRMED_REQUEST_PDU":
		return ApduType_CONFIRMED_REQUEST_PDU, true
	case "UNCONFIRMED_REQUEST_PDU":
		return ApduType_UNCONFIRMED_REQUEST_PDU, true
	case "SIMPLE_ACK_PDU":
		return ApduType_SIMPLE_ACK_PDU, true
	case "COMPLEX_ACK_PDU":
		return ApduType_COMPLEX_ACK_PDU, true
	case "SEGMENT_ACK_PDU":
		return ApduType_SEGMENT_ACK_PDU, true
	case "ERROR_PDU":
		return ApduType_ERROR_PDU, true
	case "REJECT_PDU":
		return ApduType_REJECT_PDU, true
	case "ABORT_PDU":
		return ApduType_ABORT_PDU, true
	case "APDU_UNKNOWN_8":
		return ApduType_APDU_UNKNOWN_8, true
	case "APDU_UNKNOWN_9":
		return ApduType_APDU_UNKNOWN_9, true
	case "APDU_UNKNOWN_A":
		return ApduType_APDU_UNKNOWN_A, true
	case "APDU_UNKNOWN_B":
		return ApduType_APDU_UNKNOWN_B, true
	case "APDU_UNKNOWN_C":
		return ApduType_APDU_UNKNOWN_C, true
	case "APDU_UNKNOWN_D":
		return ApduType_APDU_UNKNOWN_D, true
	case "APDU_UNKNOWN_E":
		return ApduType_APDU_UNKNOWN_E, true
	case "APDU_UNKNOWN_F":
		return ApduType_APDU_UNKNOWN_F, true
	}
	return 0, false
}

func ApduTypeKnows(value uint8) bool {
	for _, typeValue := range ApduTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastApduType(structType interface{}) ApduType {
	castFunc := func(typ interface{}) ApduType {
		if sApduType, ok := typ.(ApduType); ok {
			return sApduType
		}
		return 0
	}
	return castFunc(structType)
}

func (m ApduType) GetLengthInBits() uint16 {
	return 4
}

func (m ApduType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduTypeParse(theBytes []byte) (ApduType, error) {
	return ApduTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func ApduTypeParseWithBuffer(readBuffer utils.ReadBuffer) (ApduType, error) {
	val, err := readBuffer.ReadUint8("ApduType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ApduType")
	}
	if enum, ok := ApduTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ApduType(val), nil
	} else {
		return enum, nil
	}
}

func (e ApduType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ApduType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ApduType", 4, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ApduType) PLC4XEnumName() string {
	switch e {
	case ApduType_CONFIRMED_REQUEST_PDU:
		return "CONFIRMED_REQUEST_PDU"
	case ApduType_UNCONFIRMED_REQUEST_PDU:
		return "UNCONFIRMED_REQUEST_PDU"
	case ApduType_SIMPLE_ACK_PDU:
		return "SIMPLE_ACK_PDU"
	case ApduType_COMPLEX_ACK_PDU:
		return "COMPLEX_ACK_PDU"
	case ApduType_SEGMENT_ACK_PDU:
		return "SEGMENT_ACK_PDU"
	case ApduType_ERROR_PDU:
		return "ERROR_PDU"
	case ApduType_REJECT_PDU:
		return "REJECT_PDU"
	case ApduType_ABORT_PDU:
		return "ABORT_PDU"
	case ApduType_APDU_UNKNOWN_8:
		return "APDU_UNKNOWN_8"
	case ApduType_APDU_UNKNOWN_9:
		return "APDU_UNKNOWN_9"
	case ApduType_APDU_UNKNOWN_A:
		return "APDU_UNKNOWN_A"
	case ApduType_APDU_UNKNOWN_B:
		return "APDU_UNKNOWN_B"
	case ApduType_APDU_UNKNOWN_C:
		return "APDU_UNKNOWN_C"
	case ApduType_APDU_UNKNOWN_D:
		return "APDU_UNKNOWN_D"
	case ApduType_APDU_UNKNOWN_E:
		return "APDU_UNKNOWN_E"
	case ApduType_APDU_UNKNOWN_F:
		return "APDU_UNKNOWN_F"
	}
	return ""
}

func (e ApduType) String() string {
	return e.PLC4XEnumName()
}
