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

func AlignOffsetToType(off Offset, fieldType FieldType) Offset {
	fieldSize := FieldAlignment[fieldType]
	return (off + fieldSize - 1) / fieldSize * fieldSize
}

func AlignDynamicFieldContentOffset(off Offset, fieldType FieldType) Offset {
	contentAlignment := FieldDynamicContentAlignment[fieldType]
	return (off + contentAlignment - 1) / contentAlignment * contentAlignment
}

func (m *Message) LazyCalcOffsets() bool {
	if m.Offsets != nil {
		return true
	}
	res := make(map[int]Offset)
	var off Offset = 0
	var unionNum = 0
	for fieldNum, fieldType := range m.Scheme {
		// write the current offset
		off = AlignOffsetToType(off, fieldType)
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
			off = AlignOffsetToType(off, fieldType)
		}
		if FieldDynamic[fieldType] {
			if off + FieldSizes[fieldType] > m.Size {
				return false
			}
			contentSize := GetOffset(m.Bytes[off:])
			off += FieldSizes[fieldType]
			off = AlignDynamicFieldContentOffset(off, fieldType)
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
	return m.LazyCalcOffsets()
}

func (m *Message) RawBuffer() []byte {
	return m.Bytes[:m.Size]
}

func (m *Message) RawBufferForField(fieldNum int, unionNum int) []byte {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) || fieldNum >= len(m.Scheme) {
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
		off = AlignOffsetToType(off, fieldType)
	}
	if FieldDynamic[fieldType] {
		contentSize := GetOffset(m.Bytes[off:])
		off += FieldSizes[fieldType]
		off = AlignDynamicFieldContentOffset(off, fieldType)
		return m.Bytes[off:off+contentSize]
	} else {
		return m.Bytes[off:off+FieldSizes[fieldType]]
	}
}

func (m *Message) GetUint8(fieldNum int) uint8 {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return GetUint8(m.Bytes[off:])
}

func (m *Message) GetUint16(fieldNum int) uint16 {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return GetUint16(m.Bytes[off:])
}

func (m *Message) GetUint32(fieldNum int) uint32 {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return GetUint32(m.Bytes[off:])
}

func (m *Message) GetUint64(fieldNum int) uint64 {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return 0
	}
	off := m.Offsets[fieldNum]
	return GetUint64(m.Bytes[off:])
}

func (m *Message) GetBytes(fieldNum int) []byte {
	if !m.LazyCalcOffsets() || fieldNum >= len(m.Offsets) {
		return []byte{}
	}
	off := m.Offsets[fieldNum]
	contentSize := GetOffset(m.Bytes[off:])
	off += FieldSizes[TypeBytes]
	off = AlignDynamicFieldContentOffset(off, TypeBytes)
	return m.Bytes[off:off+contentSize]
}

func (m *Message) GetString(fieldNum int) string {
	b := m.GetBytes(fieldNum)
	if len(b) == 0 {
		return ""
	}
	return byteSliceToString(b)
}
