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

// BACnetUnconfirmedServiceChoice is an enum
type BACnetUnconfirmedServiceChoice uint8

type IBACnetUnconfirmedServiceChoice interface {
	utils.Serializable
}

const (
	BACnetUnconfirmedServiceChoice_I_AM                                  BACnetUnconfirmedServiceChoice = 0x00
	BACnetUnconfirmedServiceChoice_I_HAVE                                BACnetUnconfirmedServiceChoice = 0x01
	BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION          BACnetUnconfirmedServiceChoice = 0x02
	BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION        BACnetUnconfirmedServiceChoice = 0x03
	BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER          BACnetUnconfirmedServiceChoice = 0x04
	BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE              BACnetUnconfirmedServiceChoice = 0x05
	BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION                  BACnetUnconfirmedServiceChoice = 0x06
	BACnetUnconfirmedServiceChoice_WHO_HAS                               BACnetUnconfirmedServiceChoice = 0x07
	BACnetUnconfirmedServiceChoice_WHO_IS                                BACnetUnconfirmedServiceChoice = 0x08
	BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION              BACnetUnconfirmedServiceChoice = 0x09
	BACnetUnconfirmedServiceChoice_WRITE_GROUP                           BACnetUnconfirmedServiceChoice = 0x0A
	BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE BACnetUnconfirmedServiceChoice = 0x0B
)

var BACnetUnconfirmedServiceChoiceValues []BACnetUnconfirmedServiceChoice

func init() {
	_ = errors.New
	BACnetUnconfirmedServiceChoiceValues = []BACnetUnconfirmedServiceChoice{
		BACnetUnconfirmedServiceChoice_I_AM,
		BACnetUnconfirmedServiceChoice_I_HAVE,
		BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION,
		BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION,
		BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER,
		BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE,
		BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION,
		BACnetUnconfirmedServiceChoice_WHO_HAS,
		BACnetUnconfirmedServiceChoice_WHO_IS,
		BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION,
		BACnetUnconfirmedServiceChoice_WRITE_GROUP,
		BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE,
	}
}

func BACnetUnconfirmedServiceChoiceByValue(value uint8) (enum BACnetUnconfirmedServiceChoice, ok bool) {
	switch value {
	case 0x00:
		return BACnetUnconfirmedServiceChoice_I_AM, true
	case 0x01:
		return BACnetUnconfirmedServiceChoice_I_HAVE, true
	case 0x02:
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION, true
	case 0x03:
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION, true
	case 0x04:
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER, true
	case 0x05:
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE, true
	case 0x06:
		return BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION, true
	case 0x07:
		return BACnetUnconfirmedServiceChoice_WHO_HAS, true
	case 0x08:
		return BACnetUnconfirmedServiceChoice_WHO_IS, true
	case 0x09:
		return BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION, true
	case 0x0A:
		return BACnetUnconfirmedServiceChoice_WRITE_GROUP, true
	case 0x0B:
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	}
	return 0, false
}

func BACnetUnconfirmedServiceChoiceByName(value string) (enum BACnetUnconfirmedServiceChoice, ok bool) {
	switch value {
	case "I_AM":
		return BACnetUnconfirmedServiceChoice_I_AM, true
	case "I_HAVE":
		return BACnetUnconfirmedServiceChoice_I_HAVE, true
	case "UNCONFIRMED_COV_NOTIFICATION":
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION, true
	case "UNCONFIRMED_EVENT_NOTIFICATION":
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION, true
	case "UNCONFIRMED_PRIVATE_TRANSFER":
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER, true
	case "UNCONFIRMED_TEXT_MESSAGE":
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE, true
	case "TIME_SYNCHRONIZATION":
		return BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION, true
	case "WHO_HAS":
		return BACnetUnconfirmedServiceChoice_WHO_HAS, true
	case "WHO_IS":
		return BACnetUnconfirmedServiceChoice_WHO_IS, true
	case "UTC_TIME_SYNCHRONIZATION":
		return BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION, true
	case "WRITE_GROUP":
		return BACnetUnconfirmedServiceChoice_WRITE_GROUP, true
	case "UNCONFIRMED_COV_NOTIFICATION_MULTIPLE":
		return BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	}
	return 0, false
}

func BACnetUnconfirmedServiceChoiceKnows(value uint8) bool {
	for _, typeValue := range BACnetUnconfirmedServiceChoiceValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetUnconfirmedServiceChoice(structType interface{}) BACnetUnconfirmedServiceChoice {
	castFunc := func(typ interface{}) BACnetUnconfirmedServiceChoice {
		if sBACnetUnconfirmedServiceChoice, ok := typ.(BACnetUnconfirmedServiceChoice); ok {
			return sBACnetUnconfirmedServiceChoice
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetUnconfirmedServiceChoice) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetUnconfirmedServiceChoice) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceChoiceParse(theBytes []byte) (BACnetUnconfirmedServiceChoice, error) {
	return BACnetUnconfirmedServiceChoiceParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetUnconfirmedServiceChoiceParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetUnconfirmedServiceChoice, error) {
	val, err := readBuffer.ReadUint8("BACnetUnconfirmedServiceChoice", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetUnconfirmedServiceChoice")
	}
	if enum, ok := BACnetUnconfirmedServiceChoiceByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetUnconfirmedServiceChoice(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetUnconfirmedServiceChoice) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetUnconfirmedServiceChoice) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetUnconfirmedServiceChoice", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetUnconfirmedServiceChoice) PLC4XEnumName() string {
	switch e {
	case BACnetUnconfirmedServiceChoice_I_AM:
		return "I_AM"
	case BACnetUnconfirmedServiceChoice_I_HAVE:
		return "I_HAVE"
	case BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION:
		return "UNCONFIRMED_COV_NOTIFICATION"
	case BACnetUnconfirmedServiceChoice_UNCONFIRMED_EVENT_NOTIFICATION:
		return "UNCONFIRMED_EVENT_NOTIFICATION"
	case BACnetUnconfirmedServiceChoice_UNCONFIRMED_PRIVATE_TRANSFER:
		return "UNCONFIRMED_PRIVATE_TRANSFER"
	case BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE:
		return "UNCONFIRMED_TEXT_MESSAGE"
	case BACnetUnconfirmedServiceChoice_TIME_SYNCHRONIZATION:
		return "TIME_SYNCHRONIZATION"
	case BACnetUnconfirmedServiceChoice_WHO_HAS:
		return "WHO_HAS"
	case BACnetUnconfirmedServiceChoice_WHO_IS:
		return "WHO_IS"
	case BACnetUnconfirmedServiceChoice_UTC_TIME_SYNCHRONIZATION:
		return "UTC_TIME_SYNCHRONIZATION"
	case BACnetUnconfirmedServiceChoice_WRITE_GROUP:
		return "WRITE_GROUP"
	case BACnetUnconfirmedServiceChoice_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE:
		return "UNCONFIRMED_COV_NOTIFICATION_MULTIPLE"
	}
	return ""
}

func (e BACnetUnconfirmedServiceChoice) String() string {
	return e.PLC4XEnumName()
}
