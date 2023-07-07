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

// LightingLabelFlavour is an enum
type LightingLabelFlavour uint8

type ILightingLabelFlavour interface {
	utils.Serializable
}

const (
	LightingLabelFlavour_FLAVOUR_0 LightingLabelFlavour = 0
	LightingLabelFlavour_FLAVOUR_1 LightingLabelFlavour = 1
	LightingLabelFlavour_FLAVOUR_2 LightingLabelFlavour = 2
	LightingLabelFlavour_FLAVOUR_3 LightingLabelFlavour = 3
)

var LightingLabelFlavourValues []LightingLabelFlavour

func init() {
	_ = errors.New
	LightingLabelFlavourValues = []LightingLabelFlavour{
		LightingLabelFlavour_FLAVOUR_0,
		LightingLabelFlavour_FLAVOUR_1,
		LightingLabelFlavour_FLAVOUR_2,
		LightingLabelFlavour_FLAVOUR_3,
	}
}

func LightingLabelFlavourByValue(value uint8) (enum LightingLabelFlavour, ok bool) {
	switch value {
	case 0:
		return LightingLabelFlavour_FLAVOUR_0, true
	case 1:
		return LightingLabelFlavour_FLAVOUR_1, true
	case 2:
		return LightingLabelFlavour_FLAVOUR_2, true
	case 3:
		return LightingLabelFlavour_FLAVOUR_3, true
	}
	return 0, false
}

func LightingLabelFlavourByName(value string) (enum LightingLabelFlavour, ok bool) {
	switch value {
	case "FLAVOUR_0":
		return LightingLabelFlavour_FLAVOUR_0, true
	case "FLAVOUR_1":
		return LightingLabelFlavour_FLAVOUR_1, true
	case "FLAVOUR_2":
		return LightingLabelFlavour_FLAVOUR_2, true
	case "FLAVOUR_3":
		return LightingLabelFlavour_FLAVOUR_3, true
	}
	return 0, false
}

func LightingLabelFlavourKnows(value uint8) bool {
	for _, typeValue := range LightingLabelFlavourValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastLightingLabelFlavour(structType interface{}) LightingLabelFlavour {
	castFunc := func(typ interface{}) LightingLabelFlavour {
		if sLightingLabelFlavour, ok := typ.(LightingLabelFlavour); ok {
			return sLightingLabelFlavour
		}
		return 0
	}
	return castFunc(structType)
}

func (m LightingLabelFlavour) GetLengthInBits() uint16 {
	return 2
}

func (m LightingLabelFlavour) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LightingLabelFlavourParse(theBytes []byte) (LightingLabelFlavour, error) {
	return LightingLabelFlavourParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func LightingLabelFlavourParseWithBuffer(readBuffer utils.ReadBuffer) (LightingLabelFlavour, error) {
	val, err := readBuffer.ReadUint8("LightingLabelFlavour", 2)
	if err != nil {
		return 0, errors.Wrap(err, "error reading LightingLabelFlavour")
	}
	if enum, ok := LightingLabelFlavourByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return LightingLabelFlavour(val), nil
	} else {
		return enum, nil
	}
}

func (e LightingLabelFlavour) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e LightingLabelFlavour) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("LightingLabelFlavour", 2, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e LightingLabelFlavour) PLC4XEnumName() string {
	switch e {
	case LightingLabelFlavour_FLAVOUR_0:
		return "FLAVOUR_0"
	case LightingLabelFlavour_FLAVOUR_1:
		return "FLAVOUR_1"
	case LightingLabelFlavour_FLAVOUR_2:
		return "FLAVOUR_2"
	case LightingLabelFlavour_FLAVOUR_3:
		return "FLAVOUR_3"
	}
	return ""
}

func (e LightingLabelFlavour) String() string {
	return e.PLC4XEnumName()
}
