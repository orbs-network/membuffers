package membuffers

type MessageWriter interface {
	Reset()
	Write(buf []byte)
	GetSize() Offset
}

type _MessageWriter struct {
	_Size Offset
}

func (w *_MessageWriter) Reset() {
	w._Size = 0
}

func (w *_MessageWriter) CalcRequiredRawBufferSize() uint32 {
	w.Reset()
	w.Write(nil)
	return uint32(w._Size)
}

func (w *_MessageWriter) Write(buf []byte) {
	// override me
}

func (w *_MessageWriter) GetSize() Offset {
	return w._Size
}

func (w *_MessageWriter) WriteUint8(buf []byte, v uint8) {
	w._Size = alignOffsetToType(w._Size, TypeUint8)
	if buf != nil {
		WriteUint8(buf[w._Size:], v)
	}
	w._Size += FieldSizes[TypeUint8]
}

func (w *_MessageWriter) WriteUint16(buf []byte, v uint16) {
	w._Size = alignOffsetToType(w._Size, TypeUint16)
	if buf != nil {
		WriteUint16(buf[w._Size:], v)
	}
	w._Size += FieldSizes[TypeUint16]
}

func (w *_MessageWriter) WriteUint32(buf []byte, v uint32) {
	w._Size = alignOffsetToType(w._Size, TypeUint32)
	if buf != nil {
		WriteUint32(buf[w._Size:], v)
	}
	w._Size += FieldSizes[TypeUint32]
}

func (w *_MessageWriter) WriteUint64(buf []byte, v uint64) {
	w._Size = alignOffsetToType(w._Size, TypeUint64)
	if buf != nil {
		WriteUint64(buf[w._Size:], v)
	}
	w._Size += FieldSizes[TypeUint64]
}

func (w *_MessageWriter) WriteBytes(buf []byte, v []byte) {
	w._Size = alignOffsetToType(w._Size, TypeBytes)
	if buf != nil {
		WriteOffset(buf[w._Size:], Offset(len(v)))
	}
	w._Size += FieldSizes[TypeBytes]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeBytes)
	if buf != nil {
		copy(buf[w._Size:], v)
	}
	w._Size += Offset(len(v))
}

func (w *_MessageWriter) WriteString(buf []byte, v string) {
	w.WriteBytes(buf, []byte(v))
}

func (w *_MessageWriter) WriteUnionIndex(buf []byte, unionIndex int) {
	w._Size = alignOffsetToType(w._Size, TypeUnion)
	if buf != nil {
		WriteUnionType(buf[w._Size:], unionIndex)
	}
	w._Size += FieldSizes[TypeUnion]
}

func (w *_MessageWriter) WriteUint8Array(buf []byte, v []uint8) {
	w._Size = alignOffsetToType(w._Size, TypeUint8Array)
	if buf != nil {
		WriteOffset(buf[w._Size:], Offset(len(v)) * FieldSizes[TypeUint8])
	}
	w._Size += FieldSizes[TypeUint8Array]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeUint8Array)
	for _, vv := range v {
		w.WriteUint8(buf, vv)
	}
}

func (w *_MessageWriter) WriteUint16Array(buf []byte, v []uint16) {
	w._Size = alignOffsetToType(w._Size, TypeUint16Array)
	if buf != nil {
		WriteOffset(buf[w._Size:], Offset(len(v)) * FieldSizes[TypeUint16])
	}
	w._Size += FieldSizes[TypeUint16Array]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeUint16Array)
	for _, vv := range v {
		w.WriteUint16(buf, vv)
	}
}

func (w *_MessageWriter) WriteUint32Array(buf []byte, v []uint32) {
	w._Size = alignOffsetToType(w._Size, TypeUint32Array)
	if buf != nil {
		WriteOffset(buf[w._Size:], Offset(len(v)) * FieldSizes[TypeUint32])
	}
	w._Size += FieldSizes[TypeUint32Array]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeUint32Array)
	for _, vv := range v {
		w.WriteUint32(buf, vv)
	}
}

func (w *_MessageWriter) WriteUint64Array(buf []byte, v []uint64) {
	w._Size = alignOffsetToType(w._Size, TypeUint64Array)
	if buf != nil {
		WriteOffset(buf[w._Size:], Offset(len(v)) * FieldSizes[TypeUint64])
	}
	w._Size += FieldSizes[TypeUint64Array]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeUint64Array)
	for _, vv := range v {
		w.WriteUint64(buf, vv)
	}
}

func (w *_MessageWriter) WriteBytesArray(buf []byte, v [][]byte) {
	w._Size = alignOffsetToType(w._Size, TypeBytesArray)
	sizePlaceholderOffset := w._Size
	w._Size += FieldSizes[TypeBytesArray]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeBytesArray)
	contentSizeStartOffset := w._Size
	for _, vv := range v {
		w.WriteBytes(buf, vv)
	}
	contentSize := w._Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}

func (w *_MessageWriter) WriteStringArray(buf []byte, v []string) {
	w._Size = alignOffsetToType(w._Size, TypeStringArray)
	sizePlaceholderOffset := w._Size
	w._Size += FieldSizes[TypeStringArray]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeStringArray)
	contentSizeStartOffset := w._Size
	for _, vv := range v {
		w.WriteString(buf, vv)
	}
	contentSize := w._Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}

func (w *_MessageWriter) WriteMessage(buf []byte, v MessageWriter) {
	w._Size = alignOffsetToType(w._Size, TypeMessage)
	sizePlaceholderOffset := w._Size
	w._Size += FieldSizes[TypeMessage]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeMessage)
	v.Reset()
	if buf != nil {
		v.Write(buf[w._Size:])
	} else {
		v.Write(nil)
	}
	contentSize := v.GetSize()
	w._Size += contentSize
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}

func (w *_MessageWriter) WriteMessageArray(buf []byte, v []MessageWriter) {
	w._Size = alignOffsetToType(w._Size, TypeMessageArray)
	sizePlaceholderOffset := w._Size
	w._Size += FieldSizes[TypeMessageArray]
	w._Size = alignDynamicFieldContentOffset(w._Size, TypeMessageArray)
	contentSizeStartOffset := w._Size
	for _, vv := range v {
		w.WriteMessage(buf, vv)
	}
	contentSize := w._Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}
