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
const S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_FUNCTIONID uint8 = 0x00
const S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_NUMBERMESSAGEOBJ uint8 = 0x01

// S7PayloadUserDataItemCpuFunctionAlarmQueryResponse is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
type S7PayloadUserDataItemCpuFunctionAlarmQueryResponse interface {
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetPudicfReturnCode returns PudicfReturnCode (property field)
	GetPudicfReturnCode() DataTransportErrorCode
	// GetPudicftransportSize returns PudicftransportSize (property field)
	GetPudicftransportSize() DataTransportSize
}

// S7PayloadUserDataItemCpuFunctionAlarmQueryResponseExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionAlarmQueryResponse.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionAlarmQueryResponseExactly interface {
	S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
	isS7PayloadUserDataItemCpuFunctionAlarmQueryResponse() bool
}

// _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse struct {
	*_S7PayloadUserDataItem
	PudicfReturnCode    DataTransportErrorCode
	PudicftransportSize DataTransportSize
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetCpuFunctionType() uint8 {
	return 0x08
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetCpuSubfunction() uint8 {
	return 0x13
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetDataLength() uint16 {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) InitializeParent(parent S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.ReturnCode = returnCode
	m.TransportSize = transportSize
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetPudicfReturnCode() DataTransportErrorCode {
	return m.PudicfReturnCode
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetPudicftransportSize() DataTransportSize {
	return m.PudicftransportSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetFunctionId() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_FUNCTIONID
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetNumberMessageObj() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_NUMBERMESSAGEOBJ
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserDataItemCpuFunctionAlarmQueryResponse factory function for _S7PayloadUserDataItemCpuFunctionAlarmQueryResponse
func NewS7PayloadUserDataItemCpuFunctionAlarmQueryResponse(pudicfReturnCode DataTransportErrorCode, pudicftransportSize DataTransportSize, returnCode DataTransportErrorCode, transportSize DataTransportSize) *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse{
		PudicfReturnCode:       pudicfReturnCode,
		PudicftransportSize:    pudicftransportSize,
		_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmQueryResponse(structType interface{}) S7PayloadUserDataItemCpuFunctionAlarmQueryResponse {
	if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmQueryResponse); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmQueryResponse); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (functionId)
	lengthInBits += 8

	// Const Field (numberMessageObj)
	lengthInBits += 8

	// Simple field (pudicfReturnCode)
	lengthInBits += 8

	// Simple field (pudicftransportSize)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataItemCpuFunctionAlarmQueryResponseParse(theBytes []byte, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionAlarmQueryResponse, error) {
	return S7PayloadUserDataItemCpuFunctionAlarmQueryResponseParseWithBuffer(utils.NewReadBufferByteBased(theBytes), cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionAlarmQueryResponseParseWithBuffer(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionAlarmQueryResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	if functionId != S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_FUNCTIONID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_FUNCTIONID) + " but got " + fmt.Sprintf("%d", functionId))
	}

	// Const Field (numberMessageObj)
	numberMessageObj, _numberMessageObjErr := readBuffer.ReadUint8("numberMessageObj", 8)
	if _numberMessageObjErr != nil {
		return nil, errors.Wrap(_numberMessageObjErr, "Error parsing 'numberMessageObj' field of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	if numberMessageObj != S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_NUMBERMESSAGEOBJ {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmQueryResponse_NUMBERMESSAGEOBJ) + " but got " + fmt.Sprintf("%d", numberMessageObj))
	}

	// Simple Field (pudicfReturnCode)
	if pullErr := readBuffer.PullContext("pudicfReturnCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pudicfReturnCode")
	}
	_pudicfReturnCode, _pudicfReturnCodeErr := DataTransportErrorCodeParseWithBuffer(readBuffer)
	if _pudicfReturnCodeErr != nil {
		return nil, errors.Wrap(_pudicfReturnCodeErr, "Error parsing 'pudicfReturnCode' field of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	pudicfReturnCode := _pudicfReturnCode
	if closeErr := readBuffer.CloseContext("pudicfReturnCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pudicfReturnCode")
	}

	// Simple Field (pudicftransportSize)
	if pullErr := readBuffer.PullContext("pudicftransportSize"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pudicftransportSize")
	}
	_pudicftransportSize, _pudicftransportSizeErr := DataTransportSizeParseWithBuffer(readBuffer)
	if _pudicftransportSizeErr != nil {
		return nil, errors.Wrap(_pudicftransportSizeErr, "Error parsing 'pudicftransportSize' field of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}
	pudicftransportSize := _pudicftransportSize
	if closeErr := readBuffer.CloseContext("pudicftransportSize"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pudicftransportSize")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{},
		PudicfReturnCode:       pudicfReturnCode,
		PudicftransportSize:    pudicftransportSize,
		reservedField0:         reservedField0,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
		}

		// Const Field (functionId)
		_functionIdErr := writeBuffer.WriteUint8("functionId", 8, 0x00)
		if _functionIdErr != nil {
			return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
		}

		// Const Field (numberMessageObj)
		_numberMessageObjErr := writeBuffer.WriteUint8("numberMessageObj", 8, 0x01)
		if _numberMessageObjErr != nil {
			return errors.Wrap(_numberMessageObjErr, "Error serializing 'numberMessageObj' field")
		}

		// Simple Field (pudicfReturnCode)
		if pushErr := writeBuffer.PushContext("pudicfReturnCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pudicfReturnCode")
		}
		_pudicfReturnCodeErr := writeBuffer.WriteSerializable(m.GetPudicfReturnCode())
		if popErr := writeBuffer.PopContext("pudicfReturnCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pudicfReturnCode")
		}
		if _pudicfReturnCodeErr != nil {
			return errors.Wrap(_pudicfReturnCodeErr, "Error serializing 'pudicfReturnCode' field")
		}

		// Simple Field (pudicftransportSize)
		if pushErr := writeBuffer.PushContext("pudicftransportSize"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pudicftransportSize")
		}
		_pudicftransportSizeErr := writeBuffer.WriteSerializable(m.GetPudicftransportSize())
		if popErr := writeBuffer.PopContext("pudicftransportSize"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pudicftransportSize")
		}
		if _pudicftransportSizeErr != nil {
			return errors.Wrap(_pudicftransportSizeErr, "Error serializing 'pudicftransportSize' field")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]interface{}{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmQueryResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmQueryResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) isS7PayloadUserDataItemCpuFunctionAlarmQueryResponse() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmQueryResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}