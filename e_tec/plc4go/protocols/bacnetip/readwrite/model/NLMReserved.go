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

// NLMReserved is the corresponding interface of NLMReserved
type NLMReserved interface {
	utils.LengthAware
	utils.Serializable
	NLM
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
}

// NLMReservedExactly can be used when we want exactly this type and not a type which fulfills NLMReserved.
// This is useful for switch cases.
type NLMReservedExactly interface {
	NLMReserved
	isNLMReserved() bool
}

// _NLMReserved is the data-structure of this message
type _NLMReserved struct {
	*_NLM
	UnknownBytes []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NLMReserved) GetMessageType() uint8 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NLMReserved) InitializeParent(parent NLM) {}

func (m *_NLMReserved) GetParent() NLM {
	return m._NLM
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NLMReserved) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNLMReserved factory function for _NLMReserved
func NewNLMReserved(unknownBytes []byte, apduLength uint16) *_NLMReserved {
	_result := &_NLMReserved{
		UnknownBytes: unknownBytes,
		_NLM:         NewNLM(apduLength),
	}
	_result._NLM._NLMChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNLMReserved(structType interface{}) NLMReserved {
	if casted, ok := structType.(NLMReserved); ok {
		return casted
	}
	if casted, ok := structType.(*NLMReserved); ok {
		return *casted
	}
	return nil
}

func (m *_NLMReserved) GetTypeName() string {
	return "NLMReserved"
}

func (m *_NLMReserved) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_NLMReserved) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_NLMReserved) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NLMReservedParse(theBytes []byte, apduLength uint16) (NLMReserved, error) {
	return NLMReservedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), apduLength)
}

func NLMReservedParseWithBuffer(readBuffer utils.ReadBuffer, apduLength uint16) (NLMReserved, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NLMReserved"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NLMReserved")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(utils.InlineIf((bool((apduLength) > (0))), func() interface{} { return uint16((uint16(apduLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field of NLMReserved")
	}

	if closeErr := readBuffer.CloseContext("NLMReserved"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NLMReserved")
	}

	// Create a partially initialized instance
	_child := &_NLMReserved{
		_NLM: &_NLM{
			ApduLength: apduLength,
		},
		UnknownBytes: unknownBytes,
	}
	_child._NLM._NLMChildRequirements = _child
	return _child, nil
}

func (m *_NLMReserved) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NLMReserved) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NLMReserved"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NLMReserved")
		}

		// Array Field (unknownBytes)
		// Byte Array field (unknownBytes)
		if err := writeBuffer.WriteByteArray("unknownBytes", m.GetUnknownBytes()); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("NLMReserved"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NLMReserved")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_NLMReserved) isNLMReserved() bool {
	return true
}

func (m *_NLMReserved) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
