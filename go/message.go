package membuffers

type InternalMessage struct {
	Bytes []byte
	Size Offset
	Scheme []FieldType
	Unions [][]FieldType

	// lazily generated
	offsets map[int]Offset
}

func (m *InternalMessage) Init(buf []byte, size Offset, scheme []FieldType, unions [][]FieldType) {
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

func (m *InternalMessage) lazyCalcOffsets() bool {
	if m.offsets != nil {
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
			if unionNum >= len(m.Unions) || int(unionType) >= len(m.Unions[unionNum]) {
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
	m.offsets = res
	return true
}

func (m *InternalMessage) IsValid() bool {
	return m.lazyCalcOffsets()
}

func (m *InternalMessage) RawBuffer() []byte {
	return m.Bytes[:m.Size]
}

func (m *InternalMessage) RawBufferForField(fieldNum int, unionNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) || fieldNum >= len(m.Scheme) {
		return []byte{}
	}
	fieldType := m.Scheme[fieldNum]
	off := m.offsets[fieldNum]
	if fieldType == TypeUnion {
		unionType := GetUnionType(m.Bytes[off:])
		off += FieldSizes[TypeUnion]
		if unionNum >= len(m.Unions) || int(unionType) >= len(m.Unions[unionNum]) {
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

func (m *InternalMessage) GetOffsetInOffset(off Offset) Offset {
	return GetOffset(m.Bytes[off:])
}

func (m *InternalMessage) GetUint8InOffset(off Offset) uint8 {
	return GetUint8(m.Bytes[off:])
}

func (m *InternalMessage) SetUint8InOffset(off Offset, v uint8) {
	WriteUint8(m.Bytes[off:], v)
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
	return GetUint16(m.Bytes[off:])
}

func (m *InternalMessage) SetUint16InOffset(off Offset, v uint16) {
	WriteUint16(m.Bytes[off:], v)
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
	return GetUint32(m.Bytes[off:])
}

func (m *InternalMessage) SetUint32InOffset(off Offset, v uint32) {
	WriteUint32(m.Bytes[off:], v)
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
	return GetUint64(m.Bytes[off:])
}

func (m *InternalMessage) SetUint64InOffset(off Offset, v uint64) {
	WriteUint64(m.Bytes[off:], v)
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
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeMessage]
	off = alignDynamicFieldContentOffset(off, TypeMessage)
	return m.Bytes[off:off+contentSize], contentSize
}

func (m *InternalMessage) GetMessage(fieldNum int) (buf []byte, size Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return []byte{}, 0
	}
	off := m.offsets[fieldNum]
	return m.GetMessageInOffset(off)
}

func (m *InternalMessage) GetBytesInOffset(off Offset) []byte {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeBytes]
	off = alignDynamicFieldContentOffset(off, TypeBytes)
	if off+contentSize > Offset(len(m.Bytes)) {
		return []byte{}
	}
	return m.Bytes[off:off+contentSize]
}

func (m *InternalMessage) SetBytesInOffset(off Offset, v []byte) error {
	contentSize := GetOffset(m.Bytes[off:])
	if contentSize != Offset(len(v)) {
		return &ErrSizeMismatch{}
	}
	off += FieldSizes[TypeBytes]
	off = alignDynamicFieldContentOffset(off, TypeBytes)
	copy(m.Bytes[off:], v)
	return nil
}

func (m *InternalMessage) GetBytes(fieldNum int) []byte {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return []byte{}
	}
	off := m.offsets[fieldNum]
	return m.GetBytesInOffset(off)
}

func (m *InternalMessage) SetBytes(fieldNum int, v []byte) error {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &ErrInvalidField{}
	}
	off := m.offsets[fieldNum]
	return m.SetBytesInOffset(off, v)
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

func (m *InternalMessage) IsUnionIndex(fieldNum int, unionNum int, unionIndex uint16) (bool, Offset) {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return false, 0
	}
	off := m.offsets[fieldNum]
	unionType := GetUnionType(m.Bytes[off:])
	off += FieldSizes[TypeUnion]
	if unionNum >= len(m.Unions) || int(unionType) >= len(m.Unions[unionNum]) {
		return false, 0
	}
	fieldType := m.Unions[unionNum][unionType]
	off = alignOffsetToType(off, fieldType)
	return unionType == unionIndex, off
}

func (m *InternalMessage) GetUint8ArrayIteratorInOffset(off Offset) *Iterator {
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

func (m *InternalMessage) GetUint8ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeUint8,m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint8ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint16ArrayIteratorInOffset(off Offset) *Iterator {
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

func (m *InternalMessage) GetUint16ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeUint16,m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint16ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint32ArrayIteratorInOffset(off Offset) *Iterator {
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

func (m *InternalMessage) GetUint32ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeUint32,m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint32ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetUint64ArrayIteratorInOffset(off Offset) *Iterator {
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

func (m *InternalMessage) GetUint64ArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeUint64,m}
	}
	off := m.offsets[fieldNum]
	return m.GetUint64ArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetMessageArrayIteratorInOffset(off Offset) *Iterator {
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

func (m *InternalMessage) GetMessageArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeMessage,m}
	}
	off := m.offsets[fieldNum]
	return m.GetMessageArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetBytesArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeBytesArray]
	off = alignDynamicFieldContentOffset(off, TypeBytesArray)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeBytes,
		m: m,
	}
}

func (m *InternalMessage) GetBytesArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeBytes,m}
	}
	off := m.offsets[fieldNum]
	return m.GetMessageArrayIteratorInOffset(off)
}

func (m *InternalMessage) GetStringArrayIteratorInOffset(off Offset) *Iterator {
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeStringArray]
	off = alignDynamicFieldContentOffset(off, TypeStringArray)
	return &Iterator{
		cursor: off,
		endCursor: off+contentSize,
		fieldType: TypeString,
		m: m,
	}
}

func (m *InternalMessage) GetStringArrayIterator(fieldNum int) *Iterator {
	if !m.lazyCalcOffsets() || fieldNum >= len(m.offsets) {
		return &Iterator{0,0,TypeString,m}
	}
	off := m.offsets[fieldNum]
	return m.GetMessageArrayIteratorInOffset(off)
}