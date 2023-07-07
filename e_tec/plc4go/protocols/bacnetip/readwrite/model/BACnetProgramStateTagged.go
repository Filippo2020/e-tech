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

// BACnetProgramStateTagged is the corresponding interface of BACnetProgramStateTagged
type BACnetProgramStateTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetProgramState
}

// BACnetProgramStateTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetProgramStateTagged.
// This is useful for switch cases.
type BACnetProgramStateTaggedExactly interface {
	BACnetProgramStateTagged
	isBACnetProgramStateTagged() bool
}

// _BACnetProgramStateTagged is the data-structure of this message
type _BACnetProgramStateTagged struct {
	Header BACnetTagHeader
	Value  BACnetProgramState

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetProgramStateTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetProgramStateTagged) GetValue() BACnetProgramState {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetProgramStateTagged factory function for _BACnetProgramStateTagged
func NewBACnetProgramStateTagged(header BACnetTagHeader, value BACnetProgramState, tagNumber uint8, tagClass TagClass) *_BACnetProgramStateTagged {
	return &_BACnetProgramStateTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetProgramStateTagged(structType interface{}) BACnetProgramStateTagged {
	if casted, ok := structType.(BACnetProgramStateTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetProgramStateTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetProgramStateTagged) GetTypeName() string {
	return "BACnetProgramStateTagged"
}

func (m *_BACnetProgramStateTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetProgramStateTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetProgramStateTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetProgramStateTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetProgramStateTagged, error) {
	return BACnetProgramStateTaggedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetProgramStateTaggedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetProgramStateTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetProgramStateTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetProgramStateTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetProgramStateTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetProgramState_IDLE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetProgramStateTagged")
	}
	var value BACnetProgramState
	if _value != nil {
		value = _value.(BACnetProgramState)
	}

	if closeErr := readBuffer.CloseContext("BACnetProgramStateTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetProgramStateTagged")
	}

	// Create the instance
	return &_BACnetProgramStateTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetProgramStateTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetProgramStateTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetProgramStateTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetProgramStateTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetProgramStateTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetProgramStateTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetProgramStateTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetProgramStateTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetProgramStateTagged) isBACnetProgramStateTagged() bool {
	return true
}

func (m *_BACnetProgramStateTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
