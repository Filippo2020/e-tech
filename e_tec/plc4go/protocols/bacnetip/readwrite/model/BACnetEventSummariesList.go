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

// BACnetEventSummariesList is the corresponding interface of BACnetEventSummariesList
type BACnetEventSummariesList interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfEventSummaries returns ListOfEventSummaries (property field)
	GetListOfEventSummaries() []BACnetEventSummary
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventSummariesListExactly can be used when we want exactly this type and not a type which fulfills BACnetEventSummariesList.
// This is useful for switch cases.
type BACnetEventSummariesListExactly interface {
	BACnetEventSummariesList
	isBACnetEventSummariesList() bool
}

// _BACnetEventSummariesList is the data-structure of this message
type _BACnetEventSummariesList struct {
	OpeningTag           BACnetOpeningTag
	ListOfEventSummaries []BACnetEventSummary
	ClosingTag           BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventSummariesList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventSummariesList) GetListOfEventSummaries() []BACnetEventSummary {
	return m.ListOfEventSummaries
}

func (m *_BACnetEventSummariesList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventSummariesList factory function for _BACnetEventSummariesList
func NewBACnetEventSummariesList(openingTag BACnetOpeningTag, listOfEventSummaries []BACnetEventSummary, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventSummariesList {
	return &_BACnetEventSummariesList{OpeningTag: openingTag, ListOfEventSummaries: listOfEventSummaries, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventSummariesList(structType interface{}) BACnetEventSummariesList {
	if casted, ok := structType.(BACnetEventSummariesList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventSummariesList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventSummariesList) GetTypeName() string {
	return "BACnetEventSummariesList"
}

func (m *_BACnetEventSummariesList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventSummariesList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfEventSummaries) > 0 {
		for _, element := range m.ListOfEventSummaries {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventSummariesList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventSummariesListParse(theBytes []byte, tagNumber uint8) (BACnetEventSummariesList, error) {
	return BACnetEventSummariesListParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventSummariesListParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventSummariesList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventSummariesList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventSummariesList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventSummariesList")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfEventSummaries)
	if pullErr := readBuffer.PullContext("listOfEventSummaries", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfEventSummaries")
	}
	// Terminated array
	var listOfEventSummaries []BACnetEventSummary
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetEventSummaryParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfEventSummaries' field of BACnetEventSummariesList")
			}
			listOfEventSummaries = append(listOfEventSummaries, _item.(BACnetEventSummary))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfEventSummaries", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfEventSummaries")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventSummariesList")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventSummariesList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventSummariesList")
	}

	// Create the instance
	return &_BACnetEventSummariesList{
		TagNumber:            tagNumber,
		OpeningTag:           openingTag,
		ListOfEventSummaries: listOfEventSummaries,
		ClosingTag:           closingTag,
	}, nil
}

func (m *_BACnetEventSummariesList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventSummariesList) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventSummariesList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventSummariesList")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfEventSummaries)
	if pushErr := writeBuffer.PushContext("listOfEventSummaries", utils.withRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfEventSummaries")
	}
	for _, _element := range m.GetListOfEventSummaries() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfEventSummaries' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfEventSummaries", utils.withRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfEventSummaries")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventSummariesList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventSummariesList")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventSummariesList) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventSummariesList) isBACnetEventSummariesList() bool {
	return true
}

func (m *_BACnetEventSummariesList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
