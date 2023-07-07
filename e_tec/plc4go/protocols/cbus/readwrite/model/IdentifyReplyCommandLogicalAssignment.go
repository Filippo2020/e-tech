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

// IdentifyReplyCommandLogicalAssignment is the corresponding interface of IdentifyReplyCommandLogicalAssignment
type IdentifyReplyCommandLogicalAssignment interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetLogicAssigment returns LogicAssigment (property field)
	GetLogicAssigment() []LogicAssignment
}

// IdentifyReplyCommandLogicalAssignmentExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandLogicalAssignment.
// This is useful for switch cases.
type IdentifyReplyCommandLogicalAssignmentExactly interface {
	IdentifyReplyCommandLogicalAssignment
	isIdentifyReplyCommandLogicalAssignment() bool
}

// _IdentifyReplyCommandLogicalAssignment is the data-structure of this message
type _IdentifyReplyCommandLogicalAssignment struct {
	*_IdentifyReplyCommand
	LogicAssigment []LogicAssignment
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetAttribute() Attribute {
	return Attribute_LogicalAssignment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) InitializeParent(parent IdentifyReplyCommand) {}

func (m *_IdentifyReplyCommandLogicalAssignment) GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandLogicalAssignment) GetLogicAssigment() []LogicAssignment {
	return m.LogicAssigment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandLogicalAssignment factory function for _IdentifyReplyCommandLogicalAssignment
func NewIdentifyReplyCommandLogicalAssignment(logicAssigment []LogicAssignment, numBytes uint8) *_IdentifyReplyCommandLogicalAssignment {
	_result := &_IdentifyReplyCommandLogicalAssignment{
		LogicAssigment:        logicAssigment,
		_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandLogicalAssignment(structType interface{}) IdentifyReplyCommandLogicalAssignment {
	if casted, ok := structType.(IdentifyReplyCommandLogicalAssignment); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandLogicalAssignment); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetTypeName() string {
	return "IdentifyReplyCommandLogicalAssignment"
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.LogicAssigment) > 0 {
		for i, element := range m.LogicAssigment {
			last := i == len(m.LogicAssigment)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_IdentifyReplyCommandLogicalAssignment) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandLogicalAssignmentParse(theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandLogicalAssignment, error) {
	return IdentifyReplyCommandLogicalAssignmentParseWithBuffer(utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandLogicalAssignmentParseWithBuffer(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandLogicalAssignment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandLogicalAssignment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandLogicalAssignment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (logicAssigment)
	if pullErr := readBuffer.PullContext("logicAssigment", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logicAssigment")
	}
	// Count array
	logicAssigment := make([]LogicAssignment, numBytes)
	// This happens when the size is set conditional to 0
	if len(logicAssigment) == 0 {
		logicAssigment = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(numBytes); curItem++ {
			_item, _err := LogicAssignmentParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'logicAssigment' field of IdentifyReplyCommandLogicalAssignment")
			}
			logicAssigment[curItem] = _item.(LogicAssignment)
		}
	}
	if closeErr := readBuffer.CloseContext("logicAssigment", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logicAssigment")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandLogicalAssignment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandLogicalAssignment")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandLogicalAssignment{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		LogicAssigment: logicAssigment,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandLogicalAssignment) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandLogicalAssignment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandLogicalAssignment")
		}

		// Array Field (logicAssigment)
		if pushErr := writeBuffer.PushContext("logicAssigment", utils.withRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logicAssigment")
		}
		for _, _element := range m.GetLogicAssigment() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'logicAssigment' field")
			}
		}
		if popErr := writeBuffer.PopContext("logicAssigment", utils.withRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logicAssigment")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandLogicalAssignment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandLogicalAssignment")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_IdentifyReplyCommandLogicalAssignment) isIdentifyReplyCommandLogicalAssignment() bool {
	return true
}

func (m *_IdentifyReplyCommandLogicalAssignment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
