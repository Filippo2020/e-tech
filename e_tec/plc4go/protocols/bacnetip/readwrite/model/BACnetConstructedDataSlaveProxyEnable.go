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

// BACnetConstructedDataSlaveProxyEnable is the corresponding interface of BACnetConstructedDataSlaveProxyEnable
type BACnetConstructedDataSlaveProxyEnable interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSlaveProxyEnable returns SlaveProxyEnable (property field)
	GetSlaveProxyEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataSlaveProxyEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSlaveProxyEnable.
// This is useful for switch cases.
type BACnetConstructedDataSlaveProxyEnableExactly interface {
	BACnetConstructedDataSlaveProxyEnable
	isBACnetConstructedDataSlaveProxyEnable() bool
}

// _BACnetConstructedDataSlaveProxyEnable is the data-structure of this message
type _BACnetConstructedDataSlaveProxyEnable struct {
	*_BACnetConstructedData
	SlaveProxyEnable BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSlaveProxyEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SLAVE_PROXY_ENABLE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSlaveProxyEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSlaveProxyEnable) GetSlaveProxyEnable() BACnetApplicationTagBoolean {
	return m.SlaveProxyEnable
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSlaveProxyEnable) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetSlaveProxyEnable())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSlaveProxyEnable factory function for _BACnetConstructedDataSlaveProxyEnable
func NewBACnetConstructedDataSlaveProxyEnable(slaveProxyEnable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSlaveProxyEnable {
	_result := &_BACnetConstructedDataSlaveProxyEnable{
		SlaveProxyEnable:       slaveProxyEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSlaveProxyEnable(structType interface{}) BACnetConstructedDataSlaveProxyEnable {
	if casted, ok := structType.(BACnetConstructedDataSlaveProxyEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSlaveProxyEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetTypeName() string {
	return "BACnetConstructedDataSlaveProxyEnable"
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (slaveProxyEnable)
	lengthInBits += m.SlaveProxyEnable.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSlaveProxyEnable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSlaveProxyEnableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSlaveProxyEnable, error) {
	return BACnetConstructedDataSlaveProxyEnableParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSlaveProxyEnableParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSlaveProxyEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSlaveProxyEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSlaveProxyEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (slaveProxyEnable)
	if pullErr := readBuffer.PullContext("slaveProxyEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for slaveProxyEnable")
	}
	_slaveProxyEnable, _slaveProxyEnableErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _slaveProxyEnableErr != nil {
		return nil, errors.Wrap(_slaveProxyEnableErr, "Error parsing 'slaveProxyEnable' field of BACnetConstructedDataSlaveProxyEnable")
	}
	slaveProxyEnable := _slaveProxyEnable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("slaveProxyEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for slaveProxyEnable")
	}

	// Virtual field
	_actualValue := slaveProxyEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSlaveProxyEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSlaveProxyEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSlaveProxyEnable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SlaveProxyEnable: slaveProxyEnable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSlaveProxyEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSlaveProxyEnable) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSlaveProxyEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSlaveProxyEnable")
		}

		// Simple Field (slaveProxyEnable)
		if pushErr := writeBuffer.PushContext("slaveProxyEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for slaveProxyEnable")
		}
		_slaveProxyEnableErr := writeBuffer.WriteSerializable(m.GetSlaveProxyEnable())
		if popErr := writeBuffer.PopContext("slaveProxyEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for slaveProxyEnable")
		}
		if _slaveProxyEnableErr != nil {
			return errors.Wrap(_slaveProxyEnableErr, "Error serializing 'slaveProxyEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSlaveProxyEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSlaveProxyEnable")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSlaveProxyEnable) isBACnetConstructedDataSlaveProxyEnable() bool {
	return true
}

func (m *_BACnetConstructedDataSlaveProxyEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}