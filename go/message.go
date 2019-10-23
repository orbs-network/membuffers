// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

import (
	"math/big"
	"sync"
)

type InternalMessage struct {
	bytes  []byte
	size   Offset
	scheme []FieldType
	unions [][]FieldType

	// lazily generated
	offsets []Offset

	sync.Mutex // only used by to trick race detector (see calc_offsets_mutex.go)
}

func (m *InternalMessage) Init(buf []byte, size Offset, scheme []FieldType, unions [][]FieldType) {
	m.bytes = buf
	m.size = size
	m.scheme = scheme
	m.unions = unions
}

func alignOffsetToType(off Offset, fieldType FieldType) Offset {
	fieldSize := FieldAlignment[fieldType]
	return (off + fieldSize - 1) / fieldSize * fieldSize
}

func alignDynamicFieldContentOffset(off Offset, fieldType FieldType) Offset {
	contentAlignment := FieldDynamicContentAlignment[fieldType]
	return (off + contentAlignment - 1) / contentAlignment * contentAlignment
}

func (m *InternalMessage) _lazyCalcOffsets() bool {
	if m.offsets != nil {
		return true
	}
	res := make([]Offset, 0, len(m.scheme))
	var off Offset = 0
	var unionNum = 0
	for _, fieldType := range m.scheme {
		if off == m.size { // This means we are at end of field (but may be postfix newer fields we ignore) stop parsing
			break
		}
		// write the current offset
		off = alignOffsetToType(off, fieldType) // if last type needs padding
		if off > m.size {
			return false
		}
		res = append(res, off)

		// skip over the content to the next field
		if fieldType == TypeUnion {
			if off+FieldSizes[TypeUnion] > m.size {
				return false
			}
			unionType := GetUnionType(m.bytes[off:])
			off += FieldSizes[TypeUnion]
			if unionNum >= len(m.unions) || int(unionType) >= len(m.unions[unionNum]) {
				return false
			}
			fieldType = m.unions[unionNum][unionType]
			unionNum += 1
			off = alignOffsetToType(off, fieldType)
		}
		if FieldDynamic[fieldType] {
			if off+FieldSizes[fieldType] > m.size {
				return false
			}
			contentSize := GetOffset(m.bytes[off:])
			off += FieldSizes[fieldType]
			off = alignDynamicFieldContentOffset(off, fieldType)
			off += contentSize
		} else {
			off += FieldSizes[fieldType]
		}
	}
	if off > m.size || off == 0 { // past end of buffer or empty buffer fail
		return false
	}
	m.offsets = res
	return true
}

func (m *InternalMessage) IsValid() bool {
	if m.bytes == nil {
		panic("uninitialized membuffer, did you create it directly without a Builder or a Reader?")
	}
	return m.lazyCalcOffsets()
}

func (m *InternalMessage) RawBuffer() []byte {
	return m.bytes[:m.size]
}

func (m *InternalMessage) RawBufferForField(fieldNum int, unionNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) || fieldNum >= len(m.scheme) {
		return []byte{}
	}
	fieldType := m.scheme[fieldNum]
	off := m.offsets[fieldNum]
	if fieldType == TypeUnion {
		unionType := GetUnionType(m.bytes[off:])
		off += FieldSizes[TypeUnion]
		if unionNum >= len(m.unions) || int(unionType) >= len(m.unions[unionNum]) {
			return []byte{}
		}
		fieldType = m.unions[unionNum][unionType]
		off = alignOffsetToType(off, fieldType)
	}
	if FieldDynamic[fieldType] {
		contentSize := GetOffset(m.bytes[off:])
		off += FieldSizes[fieldType]
		off = alignDynamicFieldContentOffset(off, fieldType)
		return m.bytes[off : off+contentSize]
	} else {
		return m.bytes[off : off+FieldSizes[fieldType]]
	}
}

func (m *InternalMessage) RawBufferWithHeaderForField(fieldNum int, unionNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) || fieldNum >= len(m.scheme) {
		return []byte{}
	}
	fieldType := m.scheme[fieldNum]
	off := m.offsets[fieldNum]
	fieldHeaderOff := off
	if fieldType == TypeUnion {
		unionType := GetUnionType(m.bytes[off:])
		off += FieldSizes[TypeUnion]
		if unionNum >= len(m.unions) || int(unionType) >= len(m.unions[unionNum]) {
			return []byte{}
		}
		fieldType = m.unions[unionNum][unionType]
		off = alignOffsetToType(off, fieldType)
	}
	if FieldDynamic[fieldType] {
		contentSize := GetOffset(m.bytes[off:])
		off += FieldSizes[fieldType]
		off = alignDynamicFieldContentOffset(off, fieldType)
		return m.bytes[fieldHeaderOff : off+contentSize]
	} else {
		return m.bytes[fieldHeaderOff : off+FieldSizes[fieldType]]
	}
}

func (m *InternalMessage) GetOffsetInOffset(off Offset) Offset {
	return GetOffset(m.bytes[off:])
}

func (m *InternalMessage) GetBoolInOffset(off Offset) bool {
	return GetBool(m.bytes[off:])
}

func (m *InternalMessage) SetBoolInOffset(off Offset, v bool) {
	WriteBool(m.bytes[off:], v)
}

func (m *InternalMessage) GetBool(fieldNum int) bool {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return false
	}
	off := m.offsets[fieldNum]
	return m.GetBoolInOffset(off)
}

func (m *InternalMessage) SetBool(fieldNum int, v bool) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	m.SetBoolInOffset(off, v)
	return nil
}

func (m *InternalMessage) GetUint8InOffset(off Offset) uint8 {
	return GetUint8(m.bytes[off:])
}

func (m *InternalMessage) SetUint8InOffset(off Offset, v uint8) {
	WriteUint8(m.bytes[off:], v)
}

func (m *InternalMessage) GetUint8(fieldNum int) uint8 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return 0
	}
	off := m.offsets[fieldNum]
	return m.GetUint8InOffset(off)
}

func (m *InternalMessage) SetUint8(fieldNum int, v uint8) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	m.SetUint8InOffset(off, v)
	return nil
}

func (m *InternalMessage) GetUint16InOffset(off Offset) uint16 {
	return GetUint16(m.bytes[off:])
}

func (m *InternalMessage) SetUint16InOffset(off Offset, v uint16) {
	WriteUint16(m.bytes[off:], v)
}

func (m *InternalMessage) GetUint16(fieldNum int) uint16 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return 0
	}
	off := m.offsets[fieldNum]
	return m.GetUint16InOffset(off)
}

func (m *InternalMessage) SetUint16(fieldNum int, v uint16) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	m.SetUint16InOffset(off, v)
	return nil
}

func (m *InternalMessage) GetUint32InOffset(off Offset) uint32 {
	return GetUint32(m.bytes[off:])
}

func (m *InternalMessage) SetUint32InOffset(off Offset, v uint32) {
	WriteUint32(m.bytes[off:], v)
}

func (m *InternalMessage) GetUint32(fieldNum int) uint32 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return 0
	}
	off := m.offsets[fieldNum]
	return m.GetUint32InOffset(off)
}

func (m *InternalMessage) SetUint32(fieldNum int, v uint32) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	m.SetUint32InOffset(off, v)
	return nil
}

func (m *InternalMessage) GetUint64InOffset(off Offset) uint64 {
	return GetUint64(m.bytes[off:])
}

func (m *InternalMessage) SetUint64InOffset(off Offset, v uint64) {
	WriteUint64(m.bytes[off:], v)
}

func (m *InternalMessage) GetUint64(fieldNum int) uint64 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return 0
	}
	off := m.offsets[fieldNum]
	return m.GetUint64InOffset(off)
}

func (m *InternalMessage) SetUint64(fieldNum int, v uint64) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	m.SetUint64InOffset(off, v)
	return nil
}

func (m *InternalMessage) GetMessageInOffset(off Offset) (buf []byte, size Offset) {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeMessage]
	off = alignDynamicFieldContentOffset(off, TypeMessage)
	return m.bytes[off : off+contentSize], contentSize
}

func (m *InternalMessage) GetMessage(fieldNum int) (buf []byte, size Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return []byte{}, 0
	}
	off := m.offsets[fieldNum]
	return m.GetMessageInOffset(off)
}

func (m *InternalMessage) SetBytesInOffset(off Offset, v []byte) error {
	contentSize := GetOffset(m.bytes[off:])
	if contentSize != Offset(len(v)) {
		return &ErrSizeMismatch{}
	}
	off += FieldSizes[TypeBytes]
	off = alignDynamicFieldContentOffset(off, TypeBytes)
	copy(m.bytes[off:], v)
	return nil
}

func (m *InternalMessage) SetBytes(fieldNum int, v []byte) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	return m.SetBytesInOffset(off, v)
}

func (m *InternalMessage) SetBytes20InOffset(off Offset, v [20]byte) error {
	WriteBytes20(m.bytes[off:], v)
	return nil
}

func (m *InternalMessage) SetBytes20(fieldNum int, v [20]byte) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	return m.SetBytes20InOffset(off, v)
}

func (m *InternalMessage) SetBytes32InOffset(off Offset, v [32]byte) error {
	WriteBytes32(m.bytes[off:], v)
	return nil
}

func (m *InternalMessage) SetBytes32(fieldNum int, v [32]byte) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	return m.SetBytes32InOffset(off, v)
}

func (m *InternalMessage) SetUint256InOffset(off Offset, v *big.Int) error {
	WriteUint256(m.bytes[off:], v)
	return nil
}

func (m *InternalMessage) SetUint256(fieldNum int, v *big.Int) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	return m.SetUint256InOffset(off, v)
}

func (m *InternalMessage) GetBytesInOffset(off Offset) []byte {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeBytes]
	off = alignDynamicFieldContentOffset(off, TypeBytes)
	if off+contentSize > Offset(len(m.bytes)) {
		return []byte{}
	}
	return m.bytes[off : off+contentSize]
}

func (m *InternalMessage) GetBytes(fieldNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return []byte{}
	}
	off := m.offsets[fieldNum]
	return m.GetBytesInOffset(off)
}

func (m *InternalMessage) GetBytes20InOffset(off Offset) [20]byte {
	return GetBytes20(m.bytes[off:])
}

func (m *InternalMessage) GetBytes20(fieldNum int) [20]byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return [20]byte{}
	}
	off := m.offsets[fieldNum]
	return m.GetBytes20InOffset(off)
}

func (m *InternalMessage) GetBytes32InOffset(off Offset) [32]byte {
	return GetBytes32(m.bytes[off:])
}

func (m *InternalMessage) GetBytes32(fieldNum int) [32]byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return [32]byte{}
	}
	off := m.offsets[fieldNum]
	return m.GetBytes32InOffset(off)
}

func (m *InternalMessage) GetUint256InOffset(off Offset) *big.Int {
	return GetUint256(m.bytes[off:])
}

func (m *InternalMessage) GetUint256(fieldNum int) *big.Int {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return big.NewInt(0)
	}
	off := m.offsets[fieldNum]
	return m.GetUint256InOffset(off)
}

func (m *InternalMessage) GetStringInOffset(off Offset) string {
	b := m.GetBytesInOffset(off)
	return byteSliceToString(b)
}

func (m *InternalMessage) SetStringInOffset(off Offset, v string) error {
	return m.SetBytesInOffset(off, []byte(v))
}

func (m *InternalMessage) GetString(fieldNum int) string {
	b := m.GetBytes(fieldNum)
	return byteSliceToString(b)
}

func (m *InternalMessage) SetString(fieldNum int, v string) error {
	return m.SetBytes(fieldNum, []byte(v))
}

func (m *InternalMessage) GetUnionIndex(fieldNum int, unionNum int) uint16 {
	const invalidUnionIndex = 0xffff
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return invalidUnionIndex
	}
	off := m.offsets[fieldNum]
	unionType := GetUnionType(m.bytes[off:])
	off += FieldSizes[TypeUnion]
	if unionNum >= len(m.unions) || int(unionType) >= len(m.unions[unionNum]) {
		return invalidUnionIndex
	}
	fieldType := m.unions[unionNum][unionType]
	off = alignOffsetToType(off, fieldType)
	return unionType
}

func (m *InternalMessage) IsUnionIndex(fieldNum int, unionNum int, unionIndex uint16) (bool, Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return false, 0
	}
	off := m.offsets[fieldNum]
	unionType := GetUnionType(m.bytes[off:])
	off += FieldSizes[TypeUnion]
	if unionNum >= len(m.unions) || int(unionType) >= len(m.unions[unionNum]) {
		return false, 0
	}
	fieldType := m.unions[unionNum][unionType]
	off = alignOffsetToType(off, fieldType)
	return unionType == unionIndex, off
}

func (m *InternalMessage) GetBoolArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeBoolArray]
	off = alignDynamicFieldContentOffset(off, TypeBoolArray)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeBool,
		m:         m,
	}
}

func (m *InternalMessage) GetBoolArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeBool, m}
	}
	off := m.offsets[fieldNum]
	return m.GetBoolArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint8ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeUint8Array]
	off = alignDynamicFieldContentOffset(off, TypeUint8Array)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeUint8,
		m:         m,
	}
}

func (m *InternalMessage) GetUint8ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeUint8, m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint8ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint16ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeUint16Array]
	off = alignDynamicFieldContentOffset(off, TypeUint16Array)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeUint16,
		m:         m,
	}
}

func (m *InternalMessage) GetUint16ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeUint16, m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint16ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint32ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeUint32Array]
	off = alignDynamicFieldContentOffset(off, TypeUint32Array)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeUint32,
		m:         m,
	}
}

func (m *InternalMessage) GetUint32ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeUint32, m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint32ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint64ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeUint64Array]
	off = alignDynamicFieldContentOffset(off, TypeUint64Array)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeUint64,
		m:         m,
	}
}

func (m *InternalMessage) GetUint64ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeUint64, m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint64ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetMessageArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeMessageArray]
	off = alignDynamicFieldContentOffset(off, TypeMessageArray)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeMessage,
		m:         m,
	}
}

func (m *InternalMessage) GetMessageArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeMessage, m}
	}
	off := m.offsets[fieldNum]
	return m.GetMessageArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetBytesArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeBytesArray]
	off = alignDynamicFieldContentOffset(off, TypeBytesArray)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeBytes,
		m:         m,
	}
}

func (m *InternalMessage) GetBytesArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeBytes, m}
	}
	off := m.offsets[fieldNum]
	return m.GetBytesArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetBytes20ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeBytes20Array]
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeBytes20,
		m:         m,
	}
}

func (m *InternalMessage) GetBytes20ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeBytes20, m}
	}
	off := m.offsets[fieldNum]
	return m.GetBytes20ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetBytes32ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeBytes32Array]
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeBytes32,
		m:         m,
	}
}

func (m *InternalMessage) GetBytes32ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeBytes32, m}
	}
	off := m.offsets[fieldNum]
	return m.GetBytes32ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint256ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeUint256Array]
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeUint256,
		m:         m,
	}
}

func (m *InternalMessage) GetUint256ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeUint256, m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint256ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetStringArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.bytes[off:])
	off += FieldSizes[TypeStringArray]
	off = alignDynamicFieldContentOffset(off, TypeStringArray)
	return &Iterator{
		cursor:    off,
		endCursor: off + contentSize,
		fieldType: TypeString,
		m:         m,
	}
}

func (m *InternalMessage) GetStringArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0, 0, TypeString, m}
	}
	off := m.offsets[fieldNum]
	return m.GetStringArrayIteratorInOffset(off)
}
