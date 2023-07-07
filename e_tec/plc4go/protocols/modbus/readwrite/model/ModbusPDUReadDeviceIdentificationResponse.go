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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusPDUReadDeviceIdentificationResponse_MEITYPE uint8 = 0x0E

// ModbusPDUReadDeviceIdentificationResponse is the corresponding interface of ModbusPDUReadDeviceIdentificationResponse
type ModbusPDUReadDeviceIdentificationResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetLevel returns Level (property field)
	GetLevel() ModbusDeviceInformationLevel
	// GetIndividualAccess returns IndividualAccess (property field)
	GetIndividualAccess() bool
	// GetConformityLevel returns ConformityLevel (property field)
	GetConformityLevel() ModbusDeviceInformationConformityLevel
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() ModbusDeviceInformationMoreFollows
	// GetNextObjectId returns NextObjectId (property field)
	GetNextObjectId() uint8
	// GetObjects returns Objects (property field)
	GetObjects() []ModbusDeviceInformationObject
}

// ModbusPDUReadDeviceIdentificationResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadDeviceIdentificationResponse.
// This is useful for switch cases.
type ModbusPDUReadDeviceIdentificationResponseExactly interface {
	ModbusPDUReadDeviceIdentificationResponse
	isModbusPDUReadDeviceIdentificationResponse() bool
}

// _ModbusPDUReadDeviceIdentificationResponse is the data-structure of this message
type _ModbusPDUReadDeviceIdentificationResponse struct {
	*_ModbusPDU
	Level            ModbusDeviceInformationLevel
	IndividualAccess bool
	ConformityLevel  ModbusDeviceInformationConformityLevel
	MoreFollows      ModbusDeviceInformationMoreFollows
	NextObjectId     uint8
	Objects          []ModbusDeviceInformationObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetFunctionFlag() uint8 {
	return 0x2B
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadDeviceIdentificationResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetLevel() ModbusDeviceInformationLevel {
	return m.Level
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetIndividualAccess() bool {
	return m.IndividualAccess
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetConformityLevel() ModbusDeviceInformationConformityLevel {
	return m.ConformityLevel
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetMoreFollows() ModbusDeviceInformationMoreFollows {
	return m.MoreFollows
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetNextObjectId() uint8 {
	return m.NextObjectId
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetObjects() []ModbusDeviceInformationObject {
	return m.Objects
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetMeiType() uint8 {
	return ModbusPDUReadDeviceIdentificationResponse_MEITYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDeviceIdentificationResponse factory function for _ModbusPDUReadDeviceIdentificationResponse
func NewModbusPDUReadDeviceIdentificationResponse(level ModbusDeviceInformationLevel, individualAccess bool, conformityLevel ModbusDeviceInformationConformityLevel, moreFollows ModbusDeviceInformationMoreFollows, nextObjectId uint8, objects []ModbusDeviceInformationObject) *_ModbusPDUReadDeviceIdentificationResponse {
	_result := &_ModbusPDUReadDeviceIdentificationResponse{
		Level:            level,
		IndividualAccess: individualAccess,
		ConformityLevel:  conformityLevel,
		MoreFollows:      moreFollows,
		NextObjectId:     nextObjectId,
		Objects:          objects,
		_ModbusPDU:       NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadDeviceIdentificationResponse(structType interface{}) ModbusPDUReadDeviceIdentificationResponse {
	if casted, ok := structType.(ModbusPDUReadDeviceIdentificationResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadDeviceIdentificationResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationResponse"
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (meiType)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	// Simple field (individualAccess)
	lengthInBits += 1

	// Simple field (conformityLevel)
	lengthInBits += 7

	// Simple field (moreFollows)
	lengthInBits += 8

	// Simple field (nextObjectId)
	lengthInBits += 8

	// Implicit Field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.Objects) > 0 {
		for i, element := range m.Objects {
			last := i == len(m.Objects)-1
			lengthInBits += element.(interface{ GetLengthInBitsConditional(bool) uint16 }).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadDeviceIdentificationResponseParse(theBytes []byte, response bool) (ModbusPDUReadDeviceIdentificationResponse, error) {
	return ModbusPDUReadDeviceIdentificationResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReadDeviceIdentificationResponseParseWithBuffer(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadDeviceIdentificationResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadDeviceIdentificationResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadDeviceIdentificationResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (meiType)
	meiType, _meiTypeErr := readBuffer.ReadUint8("meiType", 8)
	if _meiTypeErr != nil {
		return nil, errors.Wrap(_meiTypeErr, "Error parsing 'meiType' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	if meiType != ModbusPDUReadDeviceIdentificationResponse_MEITYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ModbusPDUReadDeviceIdentificationResponse_MEITYPE) + " but got " + fmt.Sprintf("%d", meiType))
	}

	// Simple Field (level)
	if pullErr := readBuffer.PullContext("level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for level")
	}
	_level, _levelErr := ModbusDeviceInformationLevelParseWithBuffer(readBuffer)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	level := _level
	if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for level")
	}

	// Simple Field (individualAccess)
	_individualAccess, _individualAccessErr := readBuffer.ReadBit("individualAccess")
	if _individualAccessErr != nil {
		return nil, errors.Wrap(_individualAccessErr, "Error parsing 'individualAccess' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	individualAccess := _individualAccess

	// Simple Field (conformityLevel)
	if pullErr := readBuffer.PullContext("conformityLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for conformityLevel")
	}
	_conformityLevel, _conformityLevelErr := ModbusDeviceInformationConformityLevelParseWithBuffer(readBuffer)
	if _conformityLevelErr != nil {
		return nil, errors.Wrap(_conformityLevelErr, "Error parsing 'conformityLevel' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	conformityLevel := _conformityLevel
	if closeErr := readBuffer.CloseContext("conformityLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for conformityLevel")
	}

	// Simple Field (moreFollows)
	if pullErr := readBuffer.PullContext("moreFollows"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for moreFollows")
	}
	_moreFollows, _moreFollowsErr := ModbusDeviceInformationMoreFollowsParseWithBuffer(readBuffer)
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	moreFollows := _moreFollows
	if closeErr := readBuffer.CloseContext("moreFollows"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for moreFollows")
	}

	// Simple Field (nextObjectId)
	_nextObjectId, _nextObjectIdErr := readBuffer.ReadUint8("nextObjectId", 8)
	if _nextObjectIdErr != nil {
		return nil, errors.Wrap(_nextObjectIdErr, "Error parsing 'nextObjectId' field of ModbusPDUReadDeviceIdentificationResponse")
	}
	nextObjectId := _nextObjectId

	// Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	_ = numberOfObjects
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of ModbusPDUReadDeviceIdentificationResponse")
	}

	// Array field (objects)
	if pullErr := readBuffer.PullContext("objects", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objects")
	}
	// Count array
	objects := make([]ModbusDeviceInformationObject, numberOfObjects)
	// This happens when the size is set conditional to 0
	if len(objects) == 0 {
		objects = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(numberOfObjects); curItem++ {
			_item, _err := ModbusDeviceInformationObjectParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'objects' field of ModbusPDUReadDeviceIdentificationResponse")
			}
			objects[curItem] = _item.(ModbusDeviceInformationObject)
		}
	}
	if closeErr := readBuffer.CloseContext("objects", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objects")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDeviceIdentificationResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadDeviceIdentificationResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadDeviceIdentificationResponse{
		_ModbusPDU:       &_ModbusPDU{},
		Level:            level,
		IndividualAccess: individualAccess,
		ConformityLevel:  conformityLevel,
		MoreFollows:      moreFollows,
		NextObjectId:     nextObjectId,
		Objects:          objects,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDeviceIdentificationResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadDeviceIdentificationResponse")
		}

		// Const Field (meiType)
		_meiTypeErr := writeBuffer.WriteUint8("meiType", 8, 0x0E)
		if _meiTypeErr != nil {
			return errors.Wrap(_meiTypeErr, "Error serializing 'meiType' field")
		}

		// Simple Field (level)
		if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for level")
		}
		_levelErr := writeBuffer.WriteSerializable(m.GetLevel())
		if popErr := writeBuffer.PopContext("level"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for level")
		}
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		// Simple Field (individualAccess)
		individualAccess := bool(m.GetIndividualAccess())
		_individualAccessErr := writeBuffer.WriteBit("individualAccess", (individualAccess))
		if _individualAccessErr != nil {
			return errors.Wrap(_individualAccessErr, "Error serializing 'individualAccess' field")
		}

		// Simple Field (conformityLevel)
		if pushErr := writeBuffer.PushContext("conformityLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for conformityLevel")
		}
		_conformityLevelErr := writeBuffer.WriteSerializable(m.GetConformityLevel())
		if popErr := writeBuffer.PopContext("conformityLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for conformityLevel")
		}
		if _conformityLevelErr != nil {
			return errors.Wrap(_conformityLevelErr, "Error serializing 'conformityLevel' field")
		}

		// Simple Field (moreFollows)
		if pushErr := writeBuffer.PushContext("moreFollows"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for moreFollows")
		}
		_moreFollowsErr := writeBuffer.WriteSerializable(m.GetMoreFollows())
		if popErr := writeBuffer.PopContext("moreFollows"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for moreFollows")
		}
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Simple Field (nextObjectId)
		nextObjectId := uint8(m.GetNextObjectId())
		_nextObjectIdErr := writeBuffer.WriteUint8("nextObjectId", 8, (nextObjectId))
		if _nextObjectIdErr != nil {
			return errors.Wrap(_nextObjectIdErr, "Error serializing 'nextObjectId' field")
		}

		// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		numberOfObjects := uint8(uint8(len(m.GetObjects())))
		_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
		if _numberOfObjectsErr != nil {
			return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
		}

		// Array Field (objects)
		if pushErr := writeBuffer.PushContext("objects", utils.withRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objects")
		}
		for _, _element := range m.GetObjects() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'objects' field")
			}
		}
		if popErr := writeBuffer.PopContext("objects", utils.withRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objects")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDeviceIdentificationResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadDeviceIdentificationResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) isModbusPDUReadDeviceIdentificationResponse() bool {
	return true
}

func (m *_ModbusPDUReadDeviceIdentificationResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
