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

// KnxNetObjectServer is the corresponding interface of KnxNetObjectServer
type KnxNetObjectServer interface {
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetObjectServerExactly can be used when we want exactly this type and not a type which fulfills KnxNetObjectServer.
// This is useful for switch cases.
type KnxNetObjectServerExactly interface {
	KnxNetObjectServer
	isKnxNetObjectServer() bool
}

// _KnxNetObjectServer is the data-structure of this message
type _KnxNetObjectServer struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetObjectServer) GetServiceType() uint8 {
	return 0x08
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetObjectServer) InitializeParent(parent ServiceId) {}

func (m *_KnxNetObjectServer) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetObjectServer) GetVersion() uint8 {
	return m.Version
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetObjectServer factory function for _KnxNetObjectServer
func NewKnxNetObjectServer(version uint8) *_KnxNetObjectServer {
	_result := &_KnxNetObjectServer{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetObjectServer(structType interface{}) KnxNetObjectServer {
	if casted, ok := structType.(KnxNetObjectServer); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetObjectServer); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetObjectServer) GetTypeName() string {
	return "KnxNetObjectServer"
}

func (m *_KnxNetObjectServer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_KnxNetObjectServer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetObjectServer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxNetObjectServerParse(theBytes []byte) (KnxNetObjectServer, error) {
	return KnxNetObjectServerParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func KnxNetObjectServerParseWithBuffer(readBuffer utils.ReadBuffer) (KnxNetObjectServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("KnxNetObjectServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetObjectServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of KnxNetObjectServer")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetObjectServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetObjectServer")
	}

	// Create a partially initialized instance
	_child := &_KnxNetObjectServer{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetObjectServer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetObjectServer) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetObjectServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetObjectServer")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetObjectServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetObjectServer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_KnxNetObjectServer) isKnxNetObjectServer() bool {
	return true
}

func (m *_KnxNetObjectServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}