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

// BACnetPropertyStatesNetworkNumberQuality is the corresponding interface of BACnetPropertyStatesNetworkNumberQuality
type BACnetPropertyStatesNetworkNumberQuality interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetNetworkNumberQuality returns NetworkNumberQuality (property field)
	GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged
}

// BACnetPropertyStatesNetworkNumberQualityExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesNetworkNumberQuality.
// This is useful for switch cases.
type BACnetPropertyStatesNetworkNumberQualityExactly interface {
	BACnetPropertyStatesNetworkNumberQuality
	isBACnetPropertyStatesNetworkNumberQuality() bool
}

// _BACnetPropertyStatesNetworkNumberQuality is the data-structure of this message
type _BACnetPropertyStatesNetworkNumberQuality struct {
	*_BACnetPropertyStates
	NetworkNumberQuality BACnetNetworkNumberQualityTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNetworkNumberQuality) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetNetworkNumberQuality() BACnetNetworkNumberQualityTagged {
	return m.NetworkNumberQuality
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNetworkNumberQuality factory function for _BACnetPropertyStatesNetworkNumberQuality
func NewBACnetPropertyStatesNetworkNumberQuality(networkNumberQuality BACnetNetworkNumberQualityTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesNetworkNumberQuality {
	_result := &_BACnetPropertyStatesNetworkNumberQuality{
		NetworkNumberQuality:  networkNumberQuality,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNetworkNumberQuality(structType interface{}) BACnetPropertyStatesNetworkNumberQuality {
	if casted, ok := structType.(BACnetPropertyStatesNetworkNumberQuality); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNetworkNumberQuality); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetTypeName() string {
	return "BACnetPropertyStatesNetworkNumberQuality"
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkNumberQuality)
	lengthInBits += m.NetworkNumberQuality.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNetworkNumberQualityParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesNetworkNumberQuality, error) {
	return BACnetPropertyStatesNetworkNumberQualityParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesNetworkNumberQualityParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesNetworkNumberQuality, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNetworkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNetworkNumberQuality")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkNumberQuality)
	if pullErr := readBuffer.PullContext("networkNumberQuality"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkNumberQuality")
	}
	_networkNumberQuality, _networkNumberQualityErr := BACnetNetworkNumberQualityTaggedParseWithBuffer(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _networkNumberQualityErr != nil {
		return nil, errors.Wrap(_networkNumberQualityErr, "Error parsing 'networkNumberQuality' field of BACnetPropertyStatesNetworkNumberQuality")
	}
	networkNumberQuality := _networkNumberQuality.(BACnetNetworkNumberQualityTagged)
	if closeErr := readBuffer.CloseContext("networkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkNumberQuality")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNetworkNumberQuality"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNetworkNumberQuality")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesNetworkNumberQuality{
		_BACnetPropertyStates: &_BACnetPropertyStates{},
		NetworkNumberQuality:  networkNumberQuality,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNetworkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNetworkNumberQuality")
		}

		// Simple Field (networkNumberQuality)
		if pushErr := writeBuffer.PushContext("networkNumberQuality"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkNumberQuality")
		}
		_networkNumberQualityErr := writeBuffer.WriteSerializable(m.GetNetworkNumberQuality())
		if popErr := writeBuffer.PopContext("networkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkNumberQuality")
		}
		if _networkNumberQualityErr != nil {
			return errors.Wrap(_networkNumberQualityErr, "Error serializing 'networkNumberQuality' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNetworkNumberQuality"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNetworkNumberQuality")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) isBACnetPropertyStatesNetworkNumberQuality() bool {
	return true
}

func (m *_BACnetPropertyStatesNetworkNumberQuality) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
