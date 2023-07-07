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

// MeasurementCommandType is an enum
type MeasurementCommandType uint8

type IMeasurementCommandType interface {
	utils.Serializable
	NumberOfArguments() uint8
}

const (
	MeasurementCommandType_MEASUREMENT_EVENT MeasurementCommandType = 0x00
)

var MeasurementCommandTypeValues []MeasurementCommandType

func init() {
	_ = errors.New
	MeasurementCommandTypeValues = []MeasurementCommandType{
		MeasurementCommandType_MEASUREMENT_EVENT,
	}
}

func (e MeasurementCommandType) NumberOfArguments() uint8 {
	switch e {
	case 0x00:
		{ /* '0x00' */
			return 6
		}
	default:
		{
			return 0
		}
	}
}

func MeasurementCommandTypeFirstEnumForFieldNumberOfArguments(value uint8) (MeasurementCommandType, error) {
	for _, sizeValue := range MeasurementCommandTypeValues {
		if sizeValue.NumberOfArguments() == value {
			return sizeValue, nil
		}
	}
	return 0, errors.Errorf("enum for %v describing NumberOfArguments not found", value)
}
func MeasurementCommandTypeByValue(value uint8) (enum MeasurementCommandType, ok bool) {
	switch value {
	case 0x00:
		return MeasurementCommandType_MEASUREMENT_EVENT, true
	}
	return 0, false
}

func MeasurementCommandTypeByName(value string) (enum MeasurementCommandType, ok bool) {
	switch value {
	case "MEASUREMENT_EVENT":
		return MeasurementCommandType_MEASUREMENT_EVENT, true
	}
	return 0, false
}

func MeasurementCommandTypeKnows(value uint8) bool {
	for _, typeValue := range MeasurementCommandTypeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastMeasurementCommandType(structType interface{}) MeasurementCommandType {
	castFunc := func(typ interface{}) MeasurementCommandType {
		if sMeasurementCommandType, ok := typ.(MeasurementCommandType); ok {
			return sMeasurementCommandType
		}
		return 0
	}
	return castFunc(structType)
}

func (m MeasurementCommandType) GetLengthInBits() uint16 {
	return 4
}

func (m MeasurementCommandType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeasurementCommandTypeParse(theBytes []byte) (MeasurementCommandType, error) {
	return MeasurementCommandTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func MeasurementCommandTypeParseWithBuffer(readBuffer utils.ReadBuffer) (MeasurementCommandType, error) {
	val, err := readBuffer.ReadUint8("MeasurementCommandType", 4)
	if err != nil {
		return 0, errors.Wrap(err, "error reading MeasurementCommandType")
	}
	if enum, ok := MeasurementCommandTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return MeasurementCommandType(val), nil
	} else {
		return enum, nil
	}
}

func (e MeasurementCommandType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e MeasurementCommandType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("MeasurementCommandType", 4, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e MeasurementCommandType) PLC4XEnumName() string {
	switch e {
	case MeasurementCommandType_MEASUREMENT_EVENT:
		return "MEASUREMENT_EVENT"
	}
	return ""
}

func (e MeasurementCommandType) String() string {
	return e.PLC4XEnumName()
}
