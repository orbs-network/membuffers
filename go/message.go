package membuffers

type Message struct {
	Bytes []byte
	Size Offset
	Scheme []FieldType
	Unions [][]FieldType

	// lazily generated
	Offsets map[int]Offset
}

func (m *Message) Init(buf []byte, size Offset, scheme []FieldType, unions [][]FieldType) {
	m.Bytes = buf
	m.Size = size
	m.Scheme = scheme
	m.Unions = unions
}

func alignOffsetToType(off Offset, fieldType FieldType) Offset {
	fieldSize := FieldAlignment[fieldType]
	return (off + fieldSize - 1) / fieldSize * fieldSize
}

func alignDynamicFieldContentOffset(off Offset, fieldType FieldType) Offset {
	contentAlignment := FieldDynamicContentAlignment[fieldType]
	return (off + contentAlignment - 1) / contentAlignment * contentAlignment
}

func (m *Message) lazyCalcOffsets() bool {
	if m.Offsets != nil {
		return true
	}
	res := make(map[int]Offset)
	var off Offset = 0
	var unionNum = 0
	for fieldNum, fieldType := range m.Scheme {
		// write the current offset
		off = alignOffsetToType(off, fieldType)
		if off >= m.Size {
			return false
		}
		res[fieldNum] = off

		// skip over the content to the next field
		if fieldType == TypeUnion {
			if off + FieldSizes[TypeUnion] > m.Size {
				return false
			}
			unionType := GetUnionType(m.Bytes[off:])
			off += FieldSizes[TypeUnion]
			if unionNum >= len(m.Unions) || unionType >= len(m.Unions[unionNum]) {
				return false
			}
			fieldType = m.Unions[unionNum][unionType]
			unionNum += 1
			off = alignOffsetToType(off, fieldType)
		}
		if FieldDynamic[fieldType] {
			if off + FieldSizes[fieldType] > m.Size {
				return false
			}
			contentSize := GetOffset(m.Bytes[off:])
			off += FieldSizes[fieldType]
			off = alignDynamicFieldContentOffset(off, fieldType)
			off += contentSize
		} else {
			off += FieldSizes[fieldType]
		}
	}
	if off > m.Size {
		return false
	}
	m.Offsets = res
	return true
}

func (m *Message) IsValid() bool {
	return m.lazyCalcOffsets()
}

func (m *Message) RawBuffer() []byte {
	return m.Bytes[:m.Size]
}

func (m *Message) RawBufferForField(fieldNum int, unionNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) || fieldNum >= len(m.Scheme) {
		return []byte{}
	}
	fieldType := m.Scheme[fieldNum]
	off := m.Offsets[fieldNum]
	if fieldType == TypeUnion {
		unionType := GetUnionType(m.Bytes[off:])
		off += FieldSizes[TypeUnion]
		if unionNum >= len(m.Unions) || unionType >= len(m.Unions[unionNum]) {
			return []byte{}
		}
		fieldType = m.Unions[unionNum][unionType]
		off = alignOffsetToType(off, fieldType)
	}
	if FieldDynamic[fieldType] {
		contentSize := GetOffset(m.Bytes[off:])
		off += FieldSizes[fieldType]
		off = alignDynamicFieldContentOffset(off, fieldType)
		return m.Bytes[off:off+contentSize]
	} else {
		return m.Bytes[off:off+FieldSizes[fieldType]]
	}
}

func (m *Message) GetUint8InOffset(off Offset) uint8 {
	return GetUint8(m.Bytes[off:])
}

func (m *Message) GetUint8(fieldNum int) uint8 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return m.GetUint8InOffset(off)
}

func (m *Message) GetUint16InOffset(off Offset) uint16 {
	return GetUint16(m.Bytes[off:])
}

func (m *Message) GetUint16(fieldNum int) uint16 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return m.GetUint16InOffset(off)
}

func (m *Message) GetUint32InOffset(off Offset) uint32 {
	return GetUint32(m.Bytes[off:])
}

func (m *Message) GetUint32(fieldNum int) uint32 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return m.GetUint32InOffset(off)
}

func (m *Message) GetUint64InOffset(off Offset) uint64 {
	return GetUint64(m.Bytes[off:])
}

func (m *Message) GetUint64(fieldNum int) uint64 {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return m.GetUint64InOffset(off)
}

func (m *Message) GetMessageInOffset(off Offset) (buf []byte, size Offset) {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeMessage]
	off = alignDynamicFieldContentOffset(off, TypeMessage)
	return m.Bytes[off:off+contentSize], contentSize
}

func (m *Message) GetMessage(fieldNum int) (buf []byte, size Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return []byte{}, 0
	}
	off := m.Offsets[fieldNum]
	return m.GetMessageInOffset(off)
}

func (m *Message) GetBytesInOffset(off Offset) []byte {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeBytes]
	off = alignDynamicFieldContentOffset(off, TypeBytes)
	return m.Bytes[off:off+contentSize]
}

func (m *Message) GetBytes(fieldNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return []byte{}
	}
	off := m.Offsets[fieldNum]
	return m.GetBytesInOffset(off)
}

func (m *Message) GetStringInOffset(off Offset) string {
	b := m.GetBytesInOffset(off)
	if len(b) == 0 {
		return ""
	}
	return byteSliceToString(b)
}

func (m *Message) GetString(fieldNum int) string {
	b := m.GetBytes(fieldNum)
	if len(b) == 0 {
		return ""
	}
	return byteSliceToString(b)
}

func (m *Message) IsUnionIndex(fieldNum int, unionNum int, unionIndex int) (bool, Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return false, 0
	}
	off := m.Offsets[fieldNum]
	unionType := GetUnionType(m.Bytes[off:])
	off += FieldSizes[TypeUnion]
	if unionNum >= len(m.Unions) || unionType >= len(m.Unions[unionNum]) {
		return false, 0
	}
	fieldType := m.Unions[unionNum][unionType]
	off = alignOffsetToType(off, fieldType)
	return unionType == unionIndex, off
}

func (m *Message) GetUint8ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeUint8Array]
	off = alignDynamicFieldContentOffset(off, TypeUint8Array)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeUint8,
		m: m,
	}
}

func (m *Message) GetUint8ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return &Iterator{0,0,TypeUint8,m}
	}
	off := m.Offsets[fieldNum]
	return m.GetUint8ArrayIteratorInOffset(off)
}

func (m *Message) GetUint16ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeUint16Array]
	off = alignDynamicFieldContentOffset(off, TypeUint16Array)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeUint16,
		m: m,
	}
}

func (m *Message) GetUint16ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return &Iterator{0,0,TypeUint16,m}
	}
	off := m.Offsets[fieldNum]
	return m.GetUint16ArrayIteratorInOffset(off)
}

func (m *Message) GetUint32ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeUint32Array]
	off = alignDynamicFieldContentOffset(off, TypeUint32Array)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeUint32,
		m: m,
	}
}

func (m *Message) GetUint32ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return &Iterator{0,0,TypeUint32,m}
	}
	off := m.Offsets[fieldNum]
	return m.GetUint32ArrayIteratorInOffset(off)
}

func (m *Message) GetUint64ArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeUint64Array]
	off = alignDynamicFieldContentOffset(off, TypeUint64Array)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeUint64,
		m: m,
	}
}

func (m *Message) GetUint64ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return &Iterator{0,0,TypeUint64,m}
	}
	off := m.Offsets[fieldNum]
	return m.GetUint64ArrayIteratorInOffset(off)
}

func (m *Message) GetMessageArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeMessageArray]
	off = alignDynamicFieldContentOffset(off, TypeMessageArray)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeMessage,
		m: m,
	}
}

func (m *Message) GetMessageArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return &Iterator{0,0,TypeMessage,m}
	}
	off := m.Offsets[fieldNum]
	return m.GetMessageArrayIteratorInOffset(off)
}