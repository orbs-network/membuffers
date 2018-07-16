package membuffers

type Iterator struct {
	cursor Offset
	endCursor Offset
	fieldType FieldType
	m *InternalMessage
}

func (i *Iterator) HasNext() bool {
	return i.cursor < i.endCursor
}

func (i *Iterator) NextUint8() uint8 {
	if i.cursor+FieldSizes[TypeUint8] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint8InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint8]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint8Array)
	return res
}

func (i *Iterator) NextUint16() uint16 {
	if i.cursor+FieldSizes[TypeUint16] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint16InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint16]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint16Array)
	return res
}

func (i *Iterator) NextUint32() uint32 {
	if i.cursor+FieldSizes[TypeUint32] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint32InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint32]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint32Array)
	return res
}

func (i *Iterator) NextUint64() uint64 {
	if i.cursor+FieldSizes[TypeUint64] > i.endCursor {
		i.cursor = i.endCursor
		return 0
	}
	res := i.m.GetUint64InOffset(i.cursor)
	i.cursor += FieldSizes[TypeUint64]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeUint64Array)
	return res
}

func (i *Iterator) NextMessage() (buf []byte, size Offset) {
	if i.cursor+FieldSizes[TypeMessage] > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}, 0
	}
	resSize := i.m.GetOffsetInOffset(i.cursor)
	i.cursor += FieldSizes[TypeMessage]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeMessage)
	if i.cursor+resSize > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}, 0
	}
	resBuf := i.m.bytes[i.cursor:i.cursor+resSize]
	i.cursor += resSize
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeMessageArray)
	return resBuf, resSize
}

func (i *Iterator) NextBytes() []byte {
	if i.cursor+FieldSizes[TypeBytes] > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}
	}
	resSize := i.m.GetOffsetInOffset(i.cursor)
	i.cursor += FieldSizes[TypeBytes]
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeBytes)
	if i.cursor+resSize > i.endCursor {
		i.cursor = i.endCursor
		return []byte{}
	}
	resBuf := i.m.bytes[i.cursor:i.cursor+resSize]
	i.cursor += resSize
	i.cursor = alignDynamicFieldContentOffset(i.cursor, TypeBytesArray)
	return resBuf
}

func (i *Iterator) NextString() string {
	b := i.NextBytes()
	return byteSliceToString(b)
}