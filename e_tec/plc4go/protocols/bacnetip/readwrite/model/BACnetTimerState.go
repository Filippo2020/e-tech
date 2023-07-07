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

// BACnetTimerState is an enum
type BACnetTimerState uint8

type IBACnetTimerState interface {
	utils.Serializable
}

const (
	BACnetTimerState_IDLE    BACnetTimerState = 0
	BACnetTimerState_RUNNING BACnetTimerState = 1
	BACnetTimerState_EXPIRED BACnetTimerState = 2
)

var BACnetTimerStateValues []BACnetTimerState

func init() {
	_ = errors.New
	BACnetTimerStateValues = []BACnetTimerState{
		BACnetTimerState_IDLE,
		BACnetTimerState_RUNNING,
		BACnetTimerState_EXPIRED,
	}
}

func BACnetTimerStateByValue(value uint8) (enum BACnetTimerState, ok bool) {
	switch value {
	case 0:
		return BACnetTimerState_IDLE, true
	case 1:
		return BACnetTimerState_RUNNING, true
	case 2:
		return BACnetTimerState_EXPIRED, true
	}
	return 0, false
}

func BACnetTimerStateByName(value string) (enum BACnetTimerState, ok bool) {
	switch value {
	case "IDLE":
		return BACnetTimerState_IDLE, true
	case "RUNNING":
		return BACnetTimerState_RUNNING, true
	case "EXPIRED":
		return BACnetTimerState_EXPIRED, true
	}
	return 0, false
}

func BACnetTimerStateKnows(value uint8) bool {
	for _, typeValue := range BACnetTimerStateValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetTimerState(structType interface{}) BACnetTimerState {
	castFunc := func(typ interface{}) BACnetTimerState {
		if sBACnetTimerState, ok := typ.(BACnetTimerState); ok {
			return sBACnetTimerState
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetTimerState) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetTimerState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateParse(theBytes []byte) (BACnetTimerState, error) {
	return BACnetTimerStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetTimerStateParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetTimerState, error) {
	val, err := readBuffer.ReadUint8("BACnetTimerState", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetTimerState")
	}
	if enum, ok := BACnetTimerStateByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetTimerState(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetTimerState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetTimerState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetTimerState", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetTimerState) PLC4XEnumName() string {
	switch e {
	case BACnetTimerState_IDLE:
		return "IDLE"
	case BACnetTimerState_RUNNING:
		return "RUNNING"
	case BACnetTimerState_EXPIRED:
		return "EXPIRED"
	}
	return ""
}

func (e BACnetTimerState) String() string {
	return e.PLC4XEnumName()
}