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

// BACnetConfirmedServiceRequestVTClose is the corresponding interface of BACnetConfirmedServiceRequestVTClose
type BACnetConfirmedServiceRequestVTClose interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetListOfRemoteVtSessionIdentifiers returns ListOfRemoteVtSessionIdentifiers (property field)
	GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger
}

// BACnetConfirmedServiceRequestVTCloseExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestVTClose.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestVTCloseExactly interface {
	BACnetConfirmedServiceRequestVTClose
	isBACnetConfirmedServiceRequestVTClose() bool
}

// _BACnetConfirmedServiceRequestVTClose is the data-structure of this message
type _BACnetConfirmedServiceRequestVTClose struct {
	*_BACnetConfirmedServiceRequest
	ListOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger

	// Arguments.
	ServiceRequestPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_VT_CLOSE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestVTClose) GetListOfRemoteVtSessionIdentifiers() []BACnetApplicationTagUnsignedInteger {
	return m.ListOfRemoteVtSessionIdentifiers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestVTClose factory function for _BACnetConfirmedServiceRequestVTClose
func NewBACnetConfirmedServiceRequestVTClose(listOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger, serviceRequestLength uint32, serviceRequestPayloadLength uint32) *_BACnetConfirmedServiceRequestVTClose {
	_result := &_BACnetConfirmedServiceRequestVTClose{
		ListOfRemoteVtSessionIdentifiers: listOfRemoteVtSessionIdentifiers,
		_BACnetConfirmedServiceRequest:   NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestVTClose(structType interface{}) BACnetConfirmedServiceRequestVTClose {
	if casted, ok := structType.(BACnetConfirmedServiceRequestVTClose); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestVTClose); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetTypeName() string {
	return "BACnetConfirmedServiceRequestVTClose"
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ListOfRemoteVtSessionIdentifiers) > 0 {
		for _, element := range m.ListOfRemoteVtSessionIdentifiers {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestVTClose) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestVTCloseParse(theBytes []byte, serviceRequestLength uint32, serviceRequestPayloadLength uint32) (BACnetConfirmedServiceRequestVTClose, error) {
	return BACnetConfirmedServiceRequestVTCloseParseWithBuffer(utils.NewReadBufferByteBased(theBytes), serviceRequestLength, serviceRequestPayloadLength)
}

func BACnetConfirmedServiceRequestVTCloseParseWithBuffer(readBuffer utils.ReadBuffer, serviceRequestLength uint32, serviceRequestPayloadLength uint32) (BACnetConfirmedServiceRequestVTClose, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestVTClose"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestVTClose")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (listOfRemoteVtSessionIdentifiers)
	if pullErr := readBuffer.PullContext("listOfRemoteVtSessionIdentifiers", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfRemoteVtSessionIdentifiers")
	}
	// Length array
	var listOfRemoteVtSessionIdentifiers []BACnetApplicationTagUnsignedInteger
	{
		_listOfRemoteVtSessionIdentifiersLength := serviceRequestPayloadLength
		_listOfRemoteVtSessionIdentifiersEndPos := positionAware.GetPos() + uint16(_listOfRemoteVtSessionIdentifiersLength)
		for positionAware.GetPos() < _listOfRemoteVtSessionIdentifiersEndPos {
			_item, _err := BACnetApplicationTagParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfRemoteVtSessionIdentifiers' field of BACnetConfirmedServiceRequestVTClose")
			}
			listOfRemoteVtSessionIdentifiers = append(listOfRemoteVtSessionIdentifiers, _item.(BACnetApplicationTagUnsignedInteger))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfRemoteVtSessionIdentifiers", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfRemoteVtSessionIdentifiers")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestVTClose"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestVTClose")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestVTClose{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ListOfRemoteVtSessionIdentifiers: listOfRemoteVtSessionIdentifiers,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestVTClose) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestVTClose"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestVTClose")
		}

		// Array Field (listOfRemoteVtSessionIdentifiers)
		if pushErr := writeBuffer.PushContext("listOfRemoteVtSessionIdentifiers", utils.withRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfRemoteVtSessionIdentifiers")
		}
		for _, _element := range m.GetListOfRemoteVtSessionIdentifiers() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfRemoteVtSessionIdentifiers' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfRemoteVtSessionIdentifiers", utils.withRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfRemoteVtSessionIdentifiers")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestVTClose"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestVTClose")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestVTClose) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}

//
////

func (m *_BACnetConfirmedServiceRequestVTClose) isBACnetConfirmedServiceRequestVTClose() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestVTClose) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
