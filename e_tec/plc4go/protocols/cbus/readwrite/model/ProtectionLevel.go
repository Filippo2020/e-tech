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

// ProtectionLevel is an enum
type ProtectionLevel uint8

type IProtectionLevel interface {
	utils.Serializable
	Description() string
}

const (
	ProtectionLevel_UNLOCK_REQUIRED ProtectionLevel = 0
	ProtectionLevel_NO_WRITE_ACCESS ProtectionLevel = 1
	ProtectionLevel_NONE            ProtectionLevel = 2
	ProtectionLevel_READ_ONLY       ProtectionLevel = 3
)

var ProtectionLevelValues []ProtectionLevel

func init() {
	_ = errors.New
	ProtectionLevelValues = []ProtectionLevel{
		ProtectionLevel_UNLOCK_REQUIRED,
		ProtectionLevel_NO_WRITE_ACCESS,
		ProtectionLevel_NONE,
		ProtectionLevel_READ_ONLY,
	}
}

func (e ProtectionLevel) Description() string {
	switch e {
	case 0:
		{ /* '0' */
			return "Unlock required from C-BUS port"
		}
	case 1:
		{ /* '1' */
			return "No write access via C-BUS port"
		}
	case 2:
		{ /* '2' */
			return "None"
		}
	case 3:
		{ /* '3' */
			return "Read only"
		}
	default:
		{
			return ""
		}
	}
}

func ProtectionLevelFirstEnumForFieldDescription(value string) (ProtectionLevel, error) {
	for _, sizeValue := range ProtectionLevelValues {
		if sizeValue.Description() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing Description not found", value)
}
func ProtectionLevelByValue(value uint8) (enum ProtectionLevel, ok bool) {
	switch value {
	case 0:
		return ProtectionLevel_UNLOCK_REQUIRED, true
	case 1:
		return ProtectionLevel_NO_WRITE_ACCESS, true
	case 2:
		return ProtectionLevel_NONE, true
	case 3:
		return ProtectionLevel_READ_ONLY, true
	}
	return 0, false
}

func ProtectionLevelByName(value string) (enum ProtectionLevel, ok bool) {
	switch value {
	case "UNLOCK_REQUIRED":
		return ProtectionLevel_UNLOCK_REQUIRED, true
	case "NO_WRITE_ACCESS":
		return ProtectionLevel_NO_WRITE_ACCESS, true
	case "NONE":
		return ProtectionLevel_NONE, true
	case "READ_ONLY":
		return ProtectionLevel_READ_ONLY, true
	}
	return 0, false
}

func ProtectionLevelKnows(value uint8) bool {
	for _, typeValue := range ProtectionLevelValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastProtectionLevel(structType interface{}) ProtectionLevel {
	castFunc := func(typ interface{}) ProtectionLevel {
		if sProtectionLevel, ok := typ.(ProtectionLevel); ok {
			return sProtectionLevel
		}
		return 0
	}
	return castFunc(structType)
}

func (m ProtectionLevel) GetLengthInBits() uint16 {
	return 4
}

func (m ProtectionLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ProtectionLevelParse(theBytes []byte) (ProtectionLevel, error) {
	return ProtectionLevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func ProtectionLevelParseWithBuffer(readBuffer utils.ReadBuffer) (ProtectionLevel, error) {
	val, err := readBuffer.ReadUint8("ProtectionLevel", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading ProtectionLevel")
	}
	if enum, ok := ProtectionLevelByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return ProtectionLevel(val), nil
	} else {
		return enum, nil
	}
}

func (e ProtectionLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e ProtectionLevel) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("ProtectionLevel", 4, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e ProtectionLevel) PLC4XEnumName() string {
	switch e {
	case ProtectionLevel_UNLOCK_REQUIRED:
		return "UNLOCK_REQUIRED"
	case ProtectionLevel_NO_WRITE_ACCESS:
		return "NO_WRITE_ACCESS"
	case ProtectionLevel_NONE:
		return "NONE"
	case ProtectionLevel_READ_ONLY:
		return "READ_ONLY"
	}
	return ""
}

func (e ProtectionLevel) String() string {
	return e.PLC4XEnumName()
}
