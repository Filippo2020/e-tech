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

// BACnetConstructedDataTimeSynchronizationRecipients is the corresponding interface of BACnetConstructedDataTimeSynchronizationRecipients
type BACnetConstructedDataTimeSynchronizationRecipients interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeSynchronizationRecipients returns TimeSynchronizationRecipients (property field)
	GetTimeSynchronizationRecipients() []BACnetRecipient
}

// BACnetConstructedDataTimeSynchronizationRecipientsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeSynchronizationRecipients.
// This is useful for switch cases.
type BACnetConstructedDataTimeSynchronizationRecipientsExactly interface {
	BACnetConstructedDataTimeSynchronizationRecipients
	isBACnetConstructedDataTimeSynchronizationRecipients() bool
}

// _BACnetConstructedDataTimeSynchronizationRecipients is the data-structure of this message
type _BACnetConstructedDataTimeSynchronizationRecipients struct {
	*_BACnetConstructedData
	TimeSynchronizationRecipients []BACnetRecipient
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_RECIPIENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetTimeSynchronizationRecipients() []BACnetRecipient {
	return m.TimeSynchronizationRecipients
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeSynchronizationRecipients factory function for _BACnetConstructedDataTimeSynchronizationRecipients
func NewBACnetConstructedDataTimeSynchronizationRecipients(timeSynchronizationRecipients []BACnetRecipient, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeSynchronizationRecipients {
	_result := &_BACnetConstructedDataTimeSynchronizationRecipients{
		TimeSynchronizationRecipients: timeSynchronizationRecipients,
		_BACnetConstructedData:        NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeSynchronizationRecipients(structType interface{}) BACnetConstructedDataTimeSynchronizationRecipients {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationRecipients); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationRecipients"
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.TimeSynchronizationRecipients) > 0 {
		for _, element := range m.TimeSynchronizationRecipients {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeSynchronizationRecipientsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeSynchronizationRecipients, error) {
	return BACnetConstructedDataTimeSynchronizationRecipientsParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimeSynchronizationRecipientsParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeSynchronizationRecipients, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationRecipients"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeSynchronizationRecipients")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (timeSynchronizationRecipients)
	if pullErr := readBuffer.PullContext("timeSynchronizationRecipients", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeSynchronizationRecipients")
	}
	// Terminated array
	var timeSynchronizationRecipients []BACnetRecipient
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRecipientParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'timeSynchronizationRecipients' field of BACnetConstructedDataTimeSynchronizationRecipients")
			}
			timeSynchronizationRecipients = append(timeSynchronizationRecipients, _item.(BACnetRecipient))
		}
	}
	if closeErr := readBuffer.CloseContext("timeSynchronizationRecipients", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeSynchronizationRecipients")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationRecipients"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeSynchronizationRecipients")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeSynchronizationRecipients{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimeSynchronizationRecipients: timeSynchronizationRecipients,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationRecipients"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeSynchronizationRecipients")
		}

		// Array Field (timeSynchronizationRecipients)
		if pushErr := writeBuffer.PushContext("timeSynchronizationRecipients", utils.withRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeSynchronizationRecipients")
		}
		for _, _element := range m.GetTimeSynchronizationRecipients() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'timeSynchronizationRecipients' field")
			}
		}
		if popErr := writeBuffer.PopContext("timeSynchronizationRecipients", utils.withRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeSynchronizationRecipients")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationRecipients"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeSynchronizationRecipients")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) isBACnetConstructedDataTimeSynchronizationRecipients() bool {
	return true
}

func (m *_BACnetConstructedDataTimeSynchronizationRecipients) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
