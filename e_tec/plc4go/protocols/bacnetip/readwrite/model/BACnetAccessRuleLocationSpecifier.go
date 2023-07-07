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

// BACnetAccessRuleLocationSpecifier is an enum
type BACnetAccessRuleLocationSpecifier uint8

type IBACnetAccessRuleLocationSpecifier interface {
	utils.Serializable
}

const (
	BACnetAccessRuleLocationSpecifier_SPECIFIED BACnetAccessRuleLocationSpecifier = 0
	BACnetAccessRuleLocationSpecifier_ALL       BACnetAccessRuleLocationSpecifier = 1
)

var BACnetAccessRuleLocationSpecifierValues []BACnetAccessRuleLocationSpecifier

func init() {
	_ = errors.New
	BACnetAccessRuleLocationSpecifierValues = []BACnetAccessRuleLocationSpecifier{
		BACnetAccessRuleLocationSpecifier_SPECIFIED,
		BACnetAccessRuleLocationSpecifier_ALL,
	}
}

func BACnetAccessRuleLocationSpecifierByValue(value uint8) (enum BACnetAccessRuleLocationSpecifier, ok bool) {
	switch value {
	case 0:
		return BACnetAccessRuleLocationSpecifier_SPECIFIED, true
	case 1:
		return BACnetAccessRuleLocationSpecifier_ALL, true
	}
	return 0, false
}

func BACnetAccessRuleLocationSpecifierByName(value string) (enum BACnetAccessRuleLocationSpecifier, ok bool) {
	switch value {
	case "SPECIFIED":
		return BACnetAccessRuleLocationSpecifier_SPECIFIED, true
	case "ALL":
		return BACnetAccessRuleLocationSpecifier_ALL, true
	}
	return 0, false
}

func BACnetAccessRuleLocationSpecifierKnows(value uint8) bool {
	for _, typeValue := range BACnetAccessRuleLocationSpecifierValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessRuleLocationSpecifier(structType interface{}) BACnetAccessRuleLocationSpecifier {
	castFunc := func(typ interface{}) BACnetAccessRuleLocationSpecifier {
		if sBACnetAccessRuleLocationSpecifier, ok := typ.(BACnetAccessRuleLocationSpecifier); ok {
			return sBACnetAccessRuleLocationSpecifier
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessRuleLocationSpecifier) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetAccessRuleLocationSpecifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAccessRuleLocationSpecifierParse(theBytes []byte) (BACnetAccessRuleLocationSpecifier, error) {
	return BACnetAccessRuleLocationSpecifierParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessRuleLocationSpecifierParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetAccessRuleLocationSpecifier, error) {
	val, err := readBuffer.ReadUint8("BACnetAccessRuleLocationSpecifier", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessRuleLocationSpecifier")
	}
	if enum, ok := BACnetAccessRuleLocationSpecifierByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAccessRuleLocationSpecifier(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessRuleLocationSpecifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessRuleLocationSpecifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetAccessRuleLocationSpecifier", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessRuleLocationSpecifier) PLC4XEnumName() string {
	switch e {
	case BACnetAccessRuleLocationSpecifier_SPECIFIED:
		return "SPECIFIED"
	case BACnetAccessRuleLocationSpecifier_ALL:
		return "ALL"
	}
	return ""
}

func (e BACnetAccessRuleLocationSpecifier) String() string {
	return e.PLC4XEnumName()
}
