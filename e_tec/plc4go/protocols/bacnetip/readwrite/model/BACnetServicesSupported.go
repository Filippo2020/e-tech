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

// BACnetServicesSupported is an enum
type BACnetServicesSupported uint8

type IBACnetServicesSupported interface {
	utils.Serializable
}

const (
	BACnetServicesSupported_ACKNOWLEDGE_ALARM                     BACnetServicesSupported = 0
	BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION            BACnetServicesSupported = 1
	BACnetServicesSupported_CONFIRMED_EVENT_NOTIFICATION          BACnetServicesSupported = 2
	BACnetServicesSupported_GET_ALARM_SUMMARY                     BACnetServicesSupported = 3
	BACnetServicesSupported_GET_ENROLLMENT_SUMMARY                BACnetServicesSupported = 4
	BACnetServicesSupported_SUBSCRIBE_COV                         BACnetServicesSupported = 5
	BACnetServicesSupported_ATOMIC_READ_FILE                      BACnetServicesSupported = 6
	BACnetServicesSupported_ATOMIC_WRITE_FILE                     BACnetServicesSupported = 7
	BACnetServicesSupported_ADD_LIST_ELEMENT                      BACnetServicesSupported = 8
	BACnetServicesSupported_REMOVE_LIST_ELEMENT                   BACnetServicesSupported = 9
	BACnetServicesSupported_CREATE_OBJECT                         BACnetServicesSupported = 10
	BACnetServicesSupported_DELETE_OBJECT                         BACnetServicesSupported = 11
	BACnetServicesSupported_READ_PROPERTY                         BACnetServicesSupported = 12
	BACnetServicesSupported_READ_PROPERTY_MULTIPLE                BACnetServicesSupported = 14
	BACnetServicesSupported_WRITE_PROPERTY                        BACnetServicesSupported = 15
	BACnetServicesSupported_WRITE_PROPERTY_MULTIPLE               BACnetServicesSupported = 16
	BACnetServicesSupported_DEVICE_COMMUNICATION_CONTROL          BACnetServicesSupported = 17
	BACnetServicesSupported_CONFIRMED_PRIVATE_TRANSFER            BACnetServicesSupported = 18
	BACnetServicesSupported_CONFIRMED_TEXT_MESSAGE                BACnetServicesSupported = 19
	BACnetServicesSupported_REINITIALIZE_DEVICE                   BACnetServicesSupported = 20
	BACnetServicesSupported_VT_OPEN                               BACnetServicesSupported = 21
	BACnetServicesSupported_VT_CLOSE                              BACnetServicesSupported = 22
	BACnetServicesSupported_VT_DATA                               BACnetServicesSupported = 23
	BACnetServicesSupported_I_AM                                  BACnetServicesSupported = 26
	BACnetServicesSupported_I_HAVE                                BACnetServicesSupported = 27
	BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION          BACnetServicesSupported = 28
	BACnetServicesSupported_UNCONFIRMED_EVENT_NOTIFICATION        BACnetServicesSupported = 29
	BACnetServicesSupported_UNCONFIRMED_PRIVATE_TRANSFER          BACnetServicesSupported = 30
	BACnetServicesSupported_UNCONFIRMED_TEXT_MESSAGE              BACnetServicesSupported = 31
	BACnetServicesSupported_TIME_SYNCHRONIZATION                  BACnetServicesSupported = 32
	BACnetServicesSupported_WHO_HAS                               BACnetServicesSupported = 33
	BACnetServicesSupported_WHO_IS                                BACnetServicesSupported = 34
	BACnetServicesSupported_READ_RANGE                            BACnetServicesSupported = 35
	BACnetServicesSupported_UTC_TIME_SYNCHRONIZATION              BACnetServicesSupported = 36
	BACnetServicesSupported_LIFE_SAFETY_OPERATION                 BACnetServicesSupported = 37
	BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY                BACnetServicesSupported = 38
	BACnetServicesSupported_GET_EVENT_INFORMATION                 BACnetServicesSupported = 39
	BACnetServicesSupported_WRITE_GROUP                           BACnetServicesSupported = 40
	BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY_MULTIPLE       BACnetServicesSupported = 41
	BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION_MULTIPLE   BACnetServicesSupported = 42
	BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE BACnetServicesSupported = 43
)

var BACnetServicesSupportedValues []BACnetServicesSupported

func init() {
	_ = errors.New
	BACnetServicesSupportedValues = []BACnetServicesSupported{
		BACnetServicesSupported_ACKNOWLEDGE_ALARM,
		BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION,
		BACnetServicesSupported_CONFIRMED_EVENT_NOTIFICATION,
		BACnetServicesSupported_GET_ALARM_SUMMARY,
		BACnetServicesSupported_GET_ENROLLMENT_SUMMARY,
		BACnetServicesSupported_SUBSCRIBE_COV,
		BACnetServicesSupported_ATOMIC_READ_FILE,
		BACnetServicesSupported_ATOMIC_WRITE_FILE,
		BACnetServicesSupported_ADD_LIST_ELEMENT,
		BACnetServicesSupported_REMOVE_LIST_ELEMENT,
		BACnetServicesSupported_CREATE_OBJECT,
		BACnetServicesSupported_DELETE_OBJECT,
		BACnetServicesSupported_READ_PROPERTY,
		BACnetServicesSupported_READ_PROPERTY_MULTIPLE,
		BACnetServicesSupported_WRITE_PROPERTY,
		BACnetServicesSupported_WRITE_PROPERTY_MULTIPLE,
		BACnetServicesSupported_DEVICE_COMMUNICATION_CONTROL,
		BACnetServicesSupported_CONFIRMED_PRIVATE_TRANSFER,
		BACnetServicesSupported_CONFIRMED_TEXT_MESSAGE,
		BACnetServicesSupported_REINITIALIZE_DEVICE,
		BACnetServicesSupported_VT_OPEN,
		BACnetServicesSupported_VT_CLOSE,
		BACnetServicesSupported_VT_DATA,
		BACnetServicesSupported_I_AM,
		BACnetServicesSupported_I_HAVE,
		BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION,
		BACnetServicesSupported_UNCONFIRMED_EVENT_NOTIFICATION,
		BACnetServicesSupported_UNCONFIRMED_PRIVATE_TRANSFER,
		BACnetServicesSupported_UNCONFIRMED_TEXT_MESSAGE,
		BACnetServicesSupported_TIME_SYNCHRONIZATION,
		BACnetServicesSupported_WHO_HAS,
		BACnetServicesSupported_WHO_IS,
		BACnetServicesSupported_READ_RANGE,
		BACnetServicesSupported_UTC_TIME_SYNCHRONIZATION,
		BACnetServicesSupported_LIFE_SAFETY_OPERATION,
		BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY,
		BACnetServicesSupported_GET_EVENT_INFORMATION,
		BACnetServicesSupported_WRITE_GROUP,
		BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY_MULTIPLE,
		BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION_MULTIPLE,
		BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE,
	}
}

func BACnetServicesSupportedByValue(value uint8) (enum BACnetServicesSupported, ok bool) {
	switch value {
	case 0:
		return BACnetServicesSupported_ACKNOWLEDGE_ALARM, true
	case 1:
		return BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION, true
	case 10:
		return BACnetServicesSupported_CREATE_OBJECT, true
	case 11:
		return BACnetServicesSupported_DELETE_OBJECT, true
	case 12:
		return BACnetServicesSupported_READ_PROPERTY, true
	case 14:
		return BACnetServicesSupported_READ_PROPERTY_MULTIPLE, true
	case 15:
		return BACnetServicesSupported_WRITE_PROPERTY, true
	case 16:
		return BACnetServicesSupported_WRITE_PROPERTY_MULTIPLE, true
	case 17:
		return BACnetServicesSupported_DEVICE_COMMUNICATION_CONTROL, true
	case 18:
		return BACnetServicesSupported_CONFIRMED_PRIVATE_TRANSFER, true
	case 19:
		return BACnetServicesSupported_CONFIRMED_TEXT_MESSAGE, true
	case 2:
		return BACnetServicesSupported_CONFIRMED_EVENT_NOTIFICATION, true
	case 20:
		return BACnetServicesSupported_REINITIALIZE_DEVICE, true
	case 21:
		return BACnetServicesSupported_VT_OPEN, true
	case 22:
		return BACnetServicesSupported_VT_CLOSE, true
	case 23:
		return BACnetServicesSupported_VT_DATA, true
	case 26:
		return BACnetServicesSupported_I_AM, true
	case 27:
		return BACnetServicesSupported_I_HAVE, true
	case 28:
		return BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION, true
	case 29:
		return BACnetServicesSupported_UNCONFIRMED_EVENT_NOTIFICATION, true
	case 3:
		return BACnetServicesSupported_GET_ALARM_SUMMARY, true
	case 30:
		return BACnetServicesSupported_UNCONFIRMED_PRIVATE_TRANSFER, true
	case 31:
		return BACnetServicesSupported_UNCONFIRMED_TEXT_MESSAGE, true
	case 32:
		return BACnetServicesSupported_TIME_SYNCHRONIZATION, true
	case 33:
		return BACnetServicesSupported_WHO_HAS, true
	case 34:
		return BACnetServicesSupported_WHO_IS, true
	case 35:
		return BACnetServicesSupported_READ_RANGE, true
	case 36:
		return BACnetServicesSupported_UTC_TIME_SYNCHRONIZATION, true
	case 37:
		return BACnetServicesSupported_LIFE_SAFETY_OPERATION, true
	case 38:
		return BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY, true
	case 39:
		return BACnetServicesSupported_GET_EVENT_INFORMATION, true
	case 4:
		return BACnetServicesSupported_GET_ENROLLMENT_SUMMARY, true
	case 40:
		return BACnetServicesSupported_WRITE_GROUP, true
	case 41:
		return BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY_MULTIPLE, true
	case 42:
		return BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	case 43:
		return BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	case 5:
		return BACnetServicesSupported_SUBSCRIBE_COV, true
	case 6:
		return BACnetServicesSupported_ATOMIC_READ_FILE, true
	case 7:
		return BACnetServicesSupported_ATOMIC_WRITE_FILE, true
	case 8:
		return BACnetServicesSupported_ADD_LIST_ELEMENT, true
	case 9:
		return BACnetServicesSupported_REMOVE_LIST_ELEMENT, true
	}
	return 0, false
}

func BACnetServicesSupportedByName(value string) (enum BACnetServicesSupported, ok bool) {
	switch value {
	case "ACKNOWLEDGE_ALARM":
		return BACnetServicesSupported_ACKNOWLEDGE_ALARM, true
	case "CONFIRMED_COV_NOTIFICATION":
		return BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION, true
	case "CREATE_OBJECT":
		return BACnetServicesSupported_CREATE_OBJECT, true
	case "DELETE_OBJECT":
		return BACnetServicesSupported_DELETE_OBJECT, true
	case "READ_PROPERTY":
		return BACnetServicesSupported_READ_PROPERTY, true
	case "READ_PROPERTY_MULTIPLE":
		return BACnetServicesSupported_READ_PROPERTY_MULTIPLE, true
	case "WRITE_PROPERTY":
		return BACnetServicesSupported_WRITE_PROPERTY, true
	case "WRITE_PROPERTY_MULTIPLE":
		return BACnetServicesSupported_WRITE_PROPERTY_MULTIPLE, true
	case "DEVICE_COMMUNICATION_CONTROL":
		return BACnetServicesSupported_DEVICE_COMMUNICATION_CONTROL, true
	case "CONFIRMED_PRIVATE_TRANSFER":
		return BACnetServicesSupported_CONFIRMED_PRIVATE_TRANSFER, true
	case "CONFIRMED_TEXT_MESSAGE":
		return BACnetServicesSupported_CONFIRMED_TEXT_MESSAGE, true
	case "CONFIRMED_EVENT_NOTIFICATION":
		return BACnetServicesSupported_CONFIRMED_EVENT_NOTIFICATION, true
	case "REINITIALIZE_DEVICE":
		return BACnetServicesSupported_REINITIALIZE_DEVICE, true
	case "VT_OPEN":
		return BACnetServicesSupported_VT_OPEN, true
	case "VT_CLOSE":
		return BACnetServicesSupported_VT_CLOSE, true
	case "VT_DATA":
		return BACnetServicesSupported_VT_DATA, true
	case "I_AM":
		return BACnetServicesSupported_I_AM, true
	case "I_HAVE":
		return BACnetServicesSupported_I_HAVE, true
	case "UNCONFIRMED_COV_NOTIFICATION":
		return BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION, true
	case "UNCONFIRMED_EVENT_NOTIFICATION":
		return BACnetServicesSupported_UNCONFIRMED_EVENT_NOTIFICATION, true
	case "GET_ALARM_SUMMARY":
		return BACnetServicesSupported_GET_ALARM_SUMMARY, true
	case "UNCONFIRMED_PRIVATE_TRANSFER":
		return BACnetServicesSupported_UNCONFIRMED_PRIVATE_TRANSFER, true
	case "UNCONFIRMED_TEXT_MESSAGE":
		return BACnetServicesSupported_UNCONFIRMED_TEXT_MESSAGE, true
	case "TIME_SYNCHRONIZATION":
		return BACnetServicesSupported_TIME_SYNCHRONIZATION, true
	case "WHO_HAS":
		return BACnetServicesSupported_WHO_HAS, true
	case "WHO_IS":
		return BACnetServicesSupported_WHO_IS, true
	case "READ_RANGE":
		return BACnetServicesSupported_READ_RANGE, true
	case "UTC_TIME_SYNCHRONIZATION":
		return BACnetServicesSupported_UTC_TIME_SYNCHRONIZATION, true
	case "LIFE_SAFETY_OPERATION":
		return BACnetServicesSupported_LIFE_SAFETY_OPERATION, true
	case "SUBSCRIBE_COV_PROPERTY":
		return BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY, true
	case "GET_EVENT_INFORMATION":
		return BACnetServicesSupported_GET_EVENT_INFORMATION, true
	case "GET_ENROLLMENT_SUMMARY":
		return BACnetServicesSupported_GET_ENROLLMENT_SUMMARY, true
	case "WRITE_GROUP":
		return BACnetServicesSupported_WRITE_GROUP, true
	case "SUBSCRIBE_COV_PROPERTY_MULTIPLE":
		return BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY_MULTIPLE, true
	case "CONFIRMED_COV_NOTIFICATION_MULTIPLE":
		return BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	case "UNCONFIRMED_COV_NOTIFICATION_MULTIPLE":
		return BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE, true
	case "SUBSCRIBE_COV":
		return BACnetServicesSupported_SUBSCRIBE_COV, true
	case "ATOMIC_READ_FILE":
		return BACnetServicesSupported_ATOMIC_READ_FILE, true
	case "ATOMIC_WRITE_FILE":
		return BACnetServicesSupported_ATOMIC_WRITE_FILE, true
	case "ADD_LIST_ELEMENT":
		return BACnetServicesSupported_ADD_LIST_ELEMENT, true
	case "REMOVE_LIST_ELEMENT":
		return BACnetServicesSupported_REMOVE_LIST_ELEMENT, true
	}
	return 0, false
}

func BACnetServicesSupportedKnows(value uint8) bool {
	for _, typeValue := range BACnetServicesSupportedValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetServicesSupported(structType interface{}) BACnetServicesSupported {
	castFunc := func(typ interface{}) BACnetServicesSupported {
		if sBACnetServicesSupported, ok := typ.(BACnetServicesSupported); ok {
			return sBACnetServicesSupported
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetServicesSupported) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetServicesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServicesSupportedParse(theBytes []byte) (BACnetServicesSupported, error) {
	return BACnetServicesSupportedParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetServicesSupportedParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetServicesSupported, error) {
	val, err := readBuffer.ReadUint8("BACnetServicesSupported", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetServicesSupported")
	}
	if enum, ok := BACnetServicesSupportedByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetServicesSupported(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetServicesSupported) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetServicesSupported) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetServicesSupported", 8, uint8(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetServicesSupported) PLC4XEnumName() string {
	switch e {
	case BACnetServicesSupported_ACKNOWLEDGE_ALARM:
		return "ACKNOWLEDGE_ALARM"
	case BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION:
		return "CONFIRMED_COV_NOTIFICATION"
	case BACnetServicesSupported_CREATE_OBJECT:
		return "CREATE_OBJECT"
	case BACnetServicesSupported_DELETE_OBJECT:
		return "DELETE_OBJECT"
	case BACnetServicesSupported_READ_PROPERTY:
		return "READ_PROPERTY"
	case BACnetServicesSupported_READ_PROPERTY_MULTIPLE:
		return "READ_PROPERTY_MULTIPLE"
	case BACnetServicesSupported_WRITE_PROPERTY:
		return "WRITE_PROPERTY"
	case BACnetServicesSupported_WRITE_PROPERTY_MULTIPLE:
		return "WRITE_PROPERTY_MULTIPLE"
	case BACnetServicesSupported_DEVICE_COMMUNICATION_CONTROL:
		return "DEVICE_COMMUNICATION_CONTROL"
	case BACnetServicesSupported_CONFIRMED_PRIVATE_TRANSFER:
		return "CONFIRMED_PRIVATE_TRANSFER"
	case BACnetServicesSupported_CONFIRMED_TEXT_MESSAGE:
		return "CONFIRMED_TEXT_MESSAGE"
	case BACnetServicesSupported_CONFIRMED_EVENT_NOTIFICATION:
		return "CONFIRMED_EVENT_NOTIFICATION"
	case BACnetServicesSupported_REINITIALIZE_DEVICE:
		return "REINITIALIZE_DEVICE"
	case BACnetServicesSupported_VT_OPEN:
		return "VT_OPEN"
	case BACnetServicesSupported_VT_CLOSE:
		return "VT_CLOSE"
	case BACnetServicesSupported_VT_DATA:
		return "VT_DATA"
	case BACnetServicesSupported_I_AM:
		return "I_AM"
	case BACnetServicesSupported_I_HAVE:
		return "I_HAVE"
	case BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION:
		return "UNCONFIRMED_COV_NOTIFICATION"
	case BACnetServicesSupported_UNCONFIRMED_EVENT_NOTIFICATION:
		return "UNCONFIRMED_EVENT_NOTIFICATION"
	case BACnetServicesSupported_GET_ALARM_SUMMARY:
		return "GET_ALARM_SUMMARY"
	case BACnetServicesSupported_UNCONFIRMED_PRIVATE_TRANSFER:
		return "UNCONFIRMED_PRIVATE_TRANSFER"
	case BACnetServicesSupported_UNCONFIRMED_TEXT_MESSAGE:
		return "UNCONFIRMED_TEXT_MESSAGE"
	case BACnetServicesSupported_TIME_SYNCHRONIZATION:
		return "TIME_SYNCHRONIZATION"
	case BACnetServicesSupported_WHO_HAS:
		return "WHO_HAS"
	case BACnetServicesSupported_WHO_IS:
		return "WHO_IS"
	case BACnetServicesSupported_READ_RANGE:
		return "READ_RANGE"
	case BACnetServicesSupported_UTC_TIME_SYNCHRONIZATION:
		return "UTC_TIME_SYNCHRONIZATION"
	case BACnetServicesSupported_LIFE_SAFETY_OPERATION:
		return "LIFE_SAFETY_OPERATION"
	case BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY:
		return "SUBSCRIBE_COV_PROPERTY"
	case BACnetServicesSupported_GET_EVENT_INFORMATION:
		return "GET_EVENT_INFORMATION"
	case BACnetServicesSupported_GET_ENROLLMENT_SUMMARY:
		return "GET_ENROLLMENT_SUMMARY"
	case BACnetServicesSupported_WRITE_GROUP:
		return "WRITE_GROUP"
	case BACnetServicesSupported_SUBSCRIBE_COV_PROPERTY_MULTIPLE:
		return "SUBSCRIBE_COV_PROPERTY_MULTIPLE"
	case BACnetServicesSupported_CONFIRMED_COV_NOTIFICATION_MULTIPLE:
		return "CONFIRMED_COV_NOTIFICATION_MULTIPLE"
	case BACnetServicesSupported_UNCONFIRMED_COV_NOTIFICATION_MULTIPLE:
		return "UNCONFIRMED_COV_NOTIFICATION_MULTIPLE"
	case BACnetServicesSupported_SUBSCRIBE_COV:
		return "SUBSCRIBE_COV"
	case BACnetServicesSupported_ATOMIC_READ_FILE:
		return "ATOMIC_READ_FILE"
	case BACnetServicesSupported_ATOMIC_WRITE_FILE:
		return "ATOMIC_WRITE_FILE"
	case BACnetServicesSupported_ADD_LIST_ELEMENT:
		return "ADD_LIST_ELEMENT"
	case BACnetServicesSupported_REMOVE_LIST_ELEMENT:
		return "REMOVE_LIST_ELEMENT"
	}
	return ""
}

func (e BACnetServicesSupported) String() string {
	return e.PLC4XEnumName()
}
