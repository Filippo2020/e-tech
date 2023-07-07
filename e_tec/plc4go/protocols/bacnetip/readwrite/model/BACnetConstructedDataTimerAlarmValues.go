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

// BACnetConstructedDataTimerAlarmValues is the corresponding interface of BACnetConstructedDataTimerAlarmValues
type BACnetConstructedDataTimerAlarmValues interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetTimerStateTagged
}

// BACnetConstructedDataTimerAlarmValuesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimerAlarmValues.
// This is useful for switch cases.
type BACnetConstructedDataTimerAlarmValuesExactly interface {
	BACnetConstructedDataTimerAlarmValues
	isBACnetConstructedDataTimerAlarmValues() bool
}

// _BACnetConstructedDataTimerAlarmValues is the data-structure of this message
type _BACnetConstructedDataTimerAlarmValues struct {
	*_BACnetConstructedData
	AlarmValues []BACnetTimerStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerAlarmValues) GetAlarmValues() []BACnetTimerStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerAlarmValues factory function for _BACnetConstructedDataTimerAlarmValues
func NewBACnetConstructedDataTimerAlarmValues(alarmValues []BACnetTimerStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerAlarmValues {
	_result := &_BACnetConstructedDataTimerAlarmValues{
		AlarmValues:            alarmValues,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerAlarmValues(structType interface{}) BACnetConstructedDataTimerAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataTimerAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataTimerAlarmValues"
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerAlarmValues) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerAlarmValuesParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerAlarmValues, error) {
	return BACnetConstructedDataTimerAlarmValuesParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimerAlarmValuesParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerAlarmValues, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (alarmValues)
	if pullErr := readBuffer.PullContext("alarmValues", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alarmValues")
	}
	// Terminated array
	var alarmValues []BACnetTimerStateTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetTimerStateTaggedParseWithBuffer(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'alarmValues' field of BACnetConstructedDataTimerAlarmValues")
			}
			alarmValues = append(alarmValues, _item.(BACnetTimerStateTagged))
		}
	}
	if closeErr := readBuffer.CloseContext("alarmValues", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alarmValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerAlarmValues")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerAlarmValues{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AlarmValues: alarmValues,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerAlarmValues) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerAlarmValues")
		}

		// Array Field (alarmValues)
		if pushErr := writeBuffer.PushContext("alarmValues", utils.withRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alarmValues")
		}
		for _, _element := range m.GetAlarmValues() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'alarmValues' field")
			}
		}
		if popErr := writeBuffer.PopContext("alarmValues", utils.withRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alarmValues")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerAlarmValues")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerAlarmValues) isBACnetConstructedDataTimerAlarmValues() bool {
	return true
}

func (m *_BACnetConstructedDataTimerAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
