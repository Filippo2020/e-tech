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

// AdsDiscoveryBlockType is an enum
type AdsDiscoveryBlockType uint16

type IAdsDiscoveryBlockType interface {
	utils.Serializable
}

const (
	AdsDiscoveryBlockType_STATUS      AdsDiscoveryBlockType = 0x0001
	AdsDiscoveryBlockType_PASSWORD    AdsDiscoveryBlockType = 0x0002
	AdsDiscoveryBlockType_VERSION     AdsDiscoveryBlockType = 0x0003
	AdsDiscoveryBlockType_OS_DATA     AdsDiscoveryBlockType = 0x0004
	AdsDiscoveryBlockType_HOST_NAME   AdsDiscoveryBlockType = 0x0005
	AdsDiscoveryBlockType_AMS_NET_ID  AdsDiscoveryBlockType = 0x0007
	AdsDiscoveryBlockType_ROUTE_NAME  AdsDiscoveryBlockType = 0x000C
	AdsDiscoveryBlockType_USER_NAME   AdsDiscoveryBlockType = 0x000D
	AdsDiscoveryBlockType_FINGERPRINT AdsDiscoveryBlockType = 0x0012
)

var AdsDiscoveryBlockTypeValues []AdsDiscoveryBlockType

func init() {
	_ = errors.New
	AdsDiscoveryBlockTypeValues = []AdsDiscoveryBlockType{
		AdsDiscoveryBlockType_STATUS,
		AdsDiscoveryBlockType_PASSWORD,
		AdsDiscoveryBlockType_VERSION,
		AdsDiscoveryBlockType_OS_DATA,
		AdsDiscoveryBlockType_HOST_NAME,
		AdsDiscoveryBlockType_AMS_NET_ID,
		AdsDiscoveryBlockType_ROUTE_NAME,
		AdsDiscoveryBlockType_USER_NAME,
		AdsDiscoveryBlockType_FINGERPRINT,
	}
}

func AdsDiscoveryBlockTypeByValue(value uint16) (enum AdsDiscoveryBlockType, ok bool) {
	switch value {
	case 0x0001:
		return AdsDiscoveryBlockType_STATUS, true
	case 0x0002:
		return AdsDiscoveryBlockType_PASSWORD, true
	case 0x0003:
		return AdsDiscoveryBlockType_VERSION, true
	case 0x0004:
		return AdsDiscoveryBlockType_OS_DATA, true
	case 0x0005:
		return AdsDiscoveryBlockType_HOST_NAME, true
	case 0x0007:
		return AdsDiscoveryBlockType_AMS_NET_ID, true
	case 0x000C:
		return AdsDiscoveryBlockType_ROUTE_NAME, true
	case 0x000D:
		return AdsDiscoveryBlockType_USER_NAME, true
	case 0x0012:
		return AdsDiscoveryBlockType_FINGERPRINT, true
	}
	return 0, false
}

func AdsDiscoveryBlockTypeByName(value string) (enum AdsDiscoveryBlockType, ok bool) {
	switch value {
	case "STATUS":
		return AdsDiscoveryBlockType_STATUS, true
	case "PASSWORD":
		return AdsDiscoveryBlockType_PASSWORD, true
	case "VERSION":
		return AdsDiscoveryBlockType_VERSION, true
	case "OS_DATA":
		return AdsDiscoveryBlockType_OS_DATA, true
	case "HOST_NAME":
		return AdsDiscoveryBlockType_HOST_NAME, true
	case "AMS_NET_ID":
		return AdsDiscoveryBlockType_AMS_NET_ID, true
	case "ROUTE_NAME":
		return AdsDiscoveryBlockType_ROUTE_NAME, true
	case "USER_NAME":
		return AdsDiscoveryBlockType_USER_NAME, true
	case "FINGERPRINT":
		return AdsDiscoveryBlockType_FINGERPRINT, true
	}
	return 0, false
}

func AdsDiscoveryBlockTypeKnows(value uint16) bool {
	for _, typeValue := range AdsDiscoveryBlockTypeValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAdsDiscoveryBlockType(structType interface{}) AdsDiscoveryBlockType {
	castFunc := func(typ interface{}) AdsDiscoveryBlockType {
		if sAdsDiscoveryBlockType, ok := typ.(AdsDiscoveryBlockType); ok {
			return sAdsDiscoveryBlockType
		}
		return 0
	}
	return castFunc(structType)
}

func (m AdsDiscoveryBlockType) GetLengthInBits() uint16 {
	return 16
}

func (m AdsDiscoveryBlockType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsDiscoveryBlockTypeParse(theBytes []byte) (AdsDiscoveryBlockType, error) {
	return AdsDiscoveryBlockTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AdsDiscoveryBlockTypeParseWithBuffer(readBuffer utils.ReadBuffer) (AdsDiscoveryBlockType, error) {
	val, err := readBuffer.ReadUint16("AdsDiscoveryBlockType", 16)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AdsDiscoveryBlockType")
	}
	if enum, ok := AdsDiscoveryBlockTypeByValue(val); !ok {
		Plc4xModelLog.Debug().Msgf("no value %x found for RequestType", val)
		return AdsDiscoveryBlockType(val), nil
	} else {
		return enum, nil
	}
}

func (e AdsDiscoveryBlockType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AdsDiscoveryBlockType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("AdsDiscoveryBlockType", 16, uint16(e), utils.withAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AdsDiscoveryBlockType) PLC4XEnumName() string {
	switch e {
	case AdsDiscoveryBlockType_STATUS:
		return "STATUS"
	case AdsDiscoveryBlockType_PASSWORD:
		return "PASSWORD"
	case AdsDiscoveryBlockType_VERSION:
		return "VERSION"
	case AdsDiscoveryBlockType_OS_DATA:
		return "OS_DATA"
	case AdsDiscoveryBlockType_HOST_NAME:
		return "HOST_NAME"
	case AdsDiscoveryBlockType_AMS_NET_ID:
		return "AMS_NET_ID"
	case AdsDiscoveryBlockType_ROUTE_NAME:
		return "ROUTE_NAME"
	case AdsDiscoveryBlockType_USER_NAME:
		return "USER_NAME"
	case AdsDiscoveryBlockType_FINGERPRINT:
		return "FINGERPRINT"
	}
	return ""
}

func (e AdsDiscoveryBlockType) String() string {
	return e.PLC4XEnumName()
}
