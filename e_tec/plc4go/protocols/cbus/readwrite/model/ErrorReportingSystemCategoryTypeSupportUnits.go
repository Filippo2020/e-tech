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

// ErrorReportingSystemCategoryTypeSupportUnits is the corresponding interface of ErrorReportingSystemCategoryTypeSupportUnits
type ErrorReportingSystemCategoryTypeSupportUnits interface {
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForSupportUnits
}

// ErrorReportingSystemCategoryTypeSupportUnitsExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryTypeSupportUnits.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeSupportUnitsExactly interface {
	ErrorReportingSystemCategoryTypeSupportUnits
	isErrorReportingSystemCategoryTypeSupportUnits() bool
}

// _ErrorReportingSystemCategoryTypeSupportUnits is the data-structure of this message
type _ErrorReportingSystemCategoryTypeSupportUnits struct {
	*_ErrorReportingSystemCategoryType
	CategoryForType ErrorReportingSystemCategoryTypeForSupportUnits
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
	return ErrorReportingSystemCategoryClass_SUPPORT_UNITS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) InitializeParent(parent ErrorReportingSystemCategoryType) {
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetParent() ErrorReportingSystemCategoryType {
	return m._ErrorReportingSystemCategoryType
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetCategoryForType() ErrorReportingSystemCategoryTypeForSupportUnits {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewErrorReportingSystemCategoryTypeSupportUnits factory function for _ErrorReportingSystemCategoryTypeSupportUnits
func NewErrorReportingSystemCategoryTypeSupportUnits(categoryForType ErrorReportingSystemCategoryTypeForSupportUnits) *_ErrorReportingSystemCategoryTypeSupportUnits {
	_result := &_ErrorReportingSystemCategoryTypeSupportUnits{
		CategoryForType:                   categoryForType,
		_ErrorReportingSystemCategoryType: NewErrorReportingSystemCategoryType(),
	}
	_result._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeSupportUnits(structType interface{}) ErrorReportingSystemCategoryTypeSupportUnits {
	if casted, ok := structType.(ErrorReportingSystemCategoryTypeSupportUnits); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeSupportUnits); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeSupportUnits"
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ErrorReportingSystemCategoryTypeSupportUnitsParse(theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeSupportUnits, error) {
	return ErrorReportingSystemCategoryTypeSupportUnitsParseWithBuffer(utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeSupportUnitsParseWithBuffer(readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeSupportUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeSupportUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeSupportUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (categoryForType)
	if pullErr := readBuffer.PullContext("categoryForType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for categoryForType")
	}
	_categoryForType, _categoryForTypeErr := ErrorReportingSystemCategoryTypeForSupportUnitsParseWithBuffer(readBuffer)
	if _categoryForTypeErr != nil {
		return nil, errors.Wrap(_categoryForTypeErr, "Error parsing 'categoryForType' field of ErrorReportingSystemCategoryTypeSupportUnits")
	}
	categoryForType := _categoryForType
	if closeErr := readBuffer.CloseContext("categoryForType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for categoryForType")
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeSupportUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeSupportUnits")
	}

	// Create a partially initialized instance
	_child := &_ErrorReportingSystemCategoryTypeSupportUnits{
		_ErrorReportingSystemCategoryType: &_ErrorReportingSystemCategoryType{},
		CategoryForType:                   categoryForType,
	}
	_child._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _child
	return _child, nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeSupportUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeSupportUnits")
		}

		// Simple Field (categoryForType)
		if pushErr := writeBuffer.PushContext("categoryForType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for categoryForType")
		}
		_categoryForTypeErr := writeBuffer.WriteSerializable(m.GetCategoryForType())
		if popErr := writeBuffer.PopContext("categoryForType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for categoryForType")
		}
		if _categoryForTypeErr != nil {
			return errors.Wrap(_categoryForTypeErr, "Error serializing 'categoryForType' field")
		}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeSupportUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeSupportUnits")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) isErrorReportingSystemCategoryTypeSupportUnits() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryTypeSupportUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}