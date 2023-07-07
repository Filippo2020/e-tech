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

// BVLCBroadcastDistributionTableEntry is the corresponding interface of BVLCBroadcastDistributionTableEntry
type BVLCBroadcastDistributionTableEntry interface {
	utils.LengthAware
	utils.Serializable
	// GetIp returns Ip (property field)
	GetIp() []uint8
	// GetPort returns Port (property field)
	GetPort() uint16
	// GetBroadcastDistributionMap returns BroadcastDistributionMap (property field)
	GetBroadcastDistributionMap() []uint8
}

// BVLCBroadcastDistributionTableEntryExactly can be used when we want exactly this type and not a type which fulfills BVLCBroadcastDistributionTableEntry.
// This is useful for switch cases.
type BVLCBroadcastDistributionTableEntryExactly interface {
	BVLCBroadcastDistributionTableEntry
	isBVLCBroadcastDistributionTableEntry() bool
}

// _BVLCBroadcastDistributionTableEntry is the data-structure of this message
type _BVLCBroadcastDistributionTableEntry struct {
	Ip                       []uint8
	Port                     uint16
	BroadcastDistributionMap []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BVLCBroadcastDistributionTableEntry) GetIp() []uint8 {
	return m.Ip
}

func (m *_BVLCBroadcastDistributionTableEntry) GetPort() uint16 {
	return m.Port
}

func (m *_BVLCBroadcastDistributionTableEntry) GetBroadcastDistributionMap() []uint8 {
	return m.BroadcastDistributionMap
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCBroadcastDistributionTableEntry factory function for _BVLCBroadcastDistributionTableEntry
func NewBVLCBroadcastDistributionTableEntry(ip []uint8, port uint16, broadcastDistributionMap []uint8) *_BVLCBroadcastDistributionTableEntry {
	return &_BVLCBroadcastDistributionTableEntry{Ip: ip, Port: port, BroadcastDistributionMap: broadcastDistributionMap}
}

// Deprecated: use the interface for direct cast
func CastBVLCBroadcastDistributionTableEntry(structType interface{}) BVLCBroadcastDistributionTableEntry {
	if casted, ok := structType.(BVLCBroadcastDistributionTableEntry); ok {
		return casted
	}
	if casted, ok := structType.(*BVLCBroadcastDistributionTableEntry); ok {
		return *casted
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) GetTypeName() string {
	return "BVLCBroadcastDistributionTableEntry"
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Array field
	if len(m.Ip) > 0 {
		lengthInBits += 8 * uint16(len(m.Ip))
	}

	// Simple field (port)
	lengthInBits += 16

	// Array field
	if len(m.BroadcastDistributionMap) > 0 {
		lengthInBits += 8 * uint16(len(m.BroadcastDistributionMap))
	}

	return lengthInBits
}

func (m *_BVLCBroadcastDistributionTableEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCBroadcastDistributionTableEntryParse(theBytes []byte) (BVLCBroadcastDistributionTableEntry, error) {
	return BVLCBroadcastDistributionTableEntryParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BVLCBroadcastDistributionTableEntryParseWithBuffer(readBuffer utils.ReadBuffer) (BVLCBroadcastDistributionTableEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BVLCBroadcastDistributionTableEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BVLCBroadcastDistributionTableEntry")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (ip)
	if pullErr := readBuffer.PullContext("ip", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ip")
	}
	// Count array
	ip := make([]uint8, uint16(4))
	// This happens when the size is set conditional to 0
	if len(ip) == 0 {
		ip = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'ip' field of BVLCBroadcastDistributionTableEntry")
			}
			ip[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("ip", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ip")
	}

	// Simple Field (port)
	_port, _portErr := readBuffer.ReadUint16("port", 16)
	if _portErr != nil {
		return nil, errors.Wrap(_portErr, "Error parsing 'port' field of BVLCBroadcastDistributionTableEntry")
	}
	port := _port

	// Array field (broadcastDistributionMap)
	if pullErr := readBuffer.PullContext("broadcastDistributionMap", utils.withRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for broadcastDistributionMap")
	}
	// Count array
	broadcastDistributionMap := make([]uint8, uint16(4))
	// This happens when the size is set conditional to 0
	if len(broadcastDistributionMap) == 0 {
		broadcastDistributionMap = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'broadcastDistributionMap' field of BVLCBroadcastDistributionTableEntry")
			}
			broadcastDistributionMap[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("broadcastDistributionMap", utils.withRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for broadcastDistributionMap")
	}

	if closeErr := readBuffer.CloseContext("BVLCBroadcastDistributionTableEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BVLCBroadcastDistributionTableEntry")
	}

	// Create the instance
	return &_BVLCBroadcastDistributionTableEntry{
		Ip:                       ip,
		Port:                     port,
		BroadcastDistributionMap: broadcastDistributionMap,
	}, nil
}

func (m *_BVLCBroadcastDistributionTableEntry) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BVLCBroadcastDistributionTableEntry) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BVLCBroadcastDistributionTableEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BVLCBroadcastDistributionTableEntry")
	}

	// Array Field (ip)
	if pushErr := writeBuffer.PushContext("ip", utils.withRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ip")
	}
	for _, _element := range m.GetIp() {
		_elementErr := writeBuffer.WriteUint8("", 8, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'ip' field")
		}
	}
	if popErr := writeBuffer.PopContext("ip", utils.withRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ip")
	}

	// Simple Field (port)
	port := uint16(m.GetPort())
	_portErr := writeBuffer.WriteUint16("port", 16, (port))
	if _portErr != nil {
		return errors.Wrap(_portErr, "Error serializing 'port' field")
	}

	// Array Field (broadcastDistributionMap)
	if pushErr := writeBuffer.PushContext("broadcastDistributionMap", utils.withRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for broadcastDistributionMap")
	}
	for _, _element := range m.GetBroadcastDistributionMap() {
		_elementErr := writeBuffer.WriteUint8("", 8, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'broadcastDistributionMap' field")
		}
	}
	if popErr := writeBuffer.PopContext("broadcastDistributionMap", utils.withRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for broadcastDistributionMap")
	}

	if popErr := writeBuffer.PopContext("BVLCBroadcastDistributionTableEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BVLCBroadcastDistributionTableEntry")
	}
	return nil
}

func (m *_BVLCBroadcastDistributionTableEntry) isBVLCBroadcastDistributionTableEntry() bool {
	return true
}

func (m *_BVLCBroadcastDistributionTableEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}