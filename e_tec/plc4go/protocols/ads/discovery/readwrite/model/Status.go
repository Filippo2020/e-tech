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

// Status is an enum
type Status uint32

type IStatus interface {
	utils.Serializable
}

const (
	Status_SUCCESS              Status = 0x00000000
	Status_FAILURE_INVALID_DATA Status = 0x00000704
	Status_FAILURE_MISSING_DATA Status = 0x00000706
)

var StatusValues []Status

func init() {
	_ = errors.New
	StatusValues = []Status{
		Status_SUCCESS,
		Status_FAILURE_INVALID_DATA,
		Status_FAILURE_MISSING_DATA,
	}
}

func StatusByValue(value uint32) (enum Status, ok bool) {
	switch value {
	case 0x00000000:
		return Status_SUCCESS, true
	case 0x00000704:
		return Status_FAILURE_INVALID_DATA, true
	case 0x00000706:
		return Status_FAILURE_MISSING_DATA, true
	}
	return 0, false
}

func StatusByName(value string) (enum Status, ok bool) {
	switch value {
	case "SUCCESS":
		return Status_SUCCESS, true
	case "FAILURE_INVALID_DATA":
		return Status_FAILURE_INVALID_DATA, true
	case "FAILURE_MISSING_DATA":
		return Status_FAILURE_MISSING_DATA, true
	}
	return 0, false
}

func StatusKnows(value uint32) bool {
	for _, typeValue := range StatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastStatus(structType interface{}) Status {
	castFunc := func(typ interface{}) Status {
		if sStatus, ok := typ.(Status); ok {
			return sStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m Status) GetLengthInBits() uint16 {
	return 32
}

func (m Status) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusParse(theBytes []byte) (Status, error) {
	return StatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func StatusParseWithBuffer(readBuffer utils.ReadBuffer) (Status, error) {
	val, err := readBuffer.ReadUint32("Status", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading Status")
	}
	if enum, ok := StatusByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return Status(val), nil
	} else {
		return enum, nil
	}
}

func (e Status) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e Status) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint32("Status", 32, uint32(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e Status) PLC4XEnumName() string {
	switch e {
	case Status_SUCCESS:
		return "SUCCESS"
	case Status_FAILURE_INVALID_DATA:
		return "FAILURE_INVALID_DATA"
	case Status_FAILURE_MISSING_DATA:
		return "FAILURE_MISSING_DATA"
	}
	return ""
}

func (e Status) String() string {
	return e.PLC4XEnumName()
}
