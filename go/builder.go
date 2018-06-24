package membuffers

type MessageBuilder interface {
	Write(buf []byte) (err error)
	GetSize() Offset
	CalcRequiredSize() Offset
}

type Builder struct {
	Size Offset
}

func (w *Builder) Reset() {
	w.Size = 0
}

// override me
func (w *Builder) Write(buf []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &ErrBufferOverrun{}
		}
	}()
	w.Reset()
	return nil
}

func (w *Builder) CalcRequiredSize() Offset {
	w.Write(nil)
	return w.GetSize()
}

func (w *Builder) GetSize() Offset {
	return w.Size
}

func (w *Builder) WriteUint8(buf []byte, v uint8) {
	w.Size = alignOffsetToType(w.Size, TypeUint8)
	if buf != nil {
		WriteUint8(buf[w.Size:], v)
	}
	w.Size += FieldSizes[TypeUint8]
}

func (w *Builder) WriteUint16(buf []byte, v uint16) {
	w.Size = alignOffsetToType(w.Size, TypeUint16)
	if buf != nil {
		WriteUint16(buf[w.Size:], v)
	}
	w.Size += FieldSizes[TypeUint16]
}

func (w *Builder) WriteUint32(buf []byte, v uint32) {
	w.Size = alignOffsetToType(w.Size, TypeUint32)
	if buf != nil {
		WriteUint32(buf[w.Size:], v)
	}
	w.Size += FieldSizes[TypeUint32]
}

func (w *Builder) WriteUint64(buf []byte, v uint64) {
	w.Size = alignOffsetToType(w.Size, TypeUint64)
	if buf != nil {
		WriteUint64(buf[w.Size:], v)
	}
	w.Size += FieldSizes[TypeUint64]
}

func (w *Builder) WriteBytes(buf []byte, v []byte) {
	w.Size = alignOffsetToType(w.Size, TypeBytes)
	if buf != nil {
		WriteOffset(buf[w.Size:], Offset(len(v)))
	}
	w.Size += FieldSizes[TypeBytes]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeBytes)
	if v != nil {
		if buf != nil {
			copy(buf[w.Size:], v)
		}
		w.Size += Offset(len(v))
	}
}

func (w *Builder) WriteString(buf []byte, v string) {
	w.WriteBytes(buf, []byte(v))
}

func (w *Builder) WriteUnionIndex(buf []byte, unionIndex uint16) {
	w.Size = alignOffsetToType(w.Size, TypeUnion)
	if buf != nil {
		WriteUnionType(buf[w.Size:], unionIndex)
	}
	w.Size += FieldSizes[TypeUnion]
}

func (w *Builder) WriteUint8Array(buf []byte, v []uint8) {
	w.Size = alignOffsetToType(w.Size, TypeUint8Array)
	if buf != nil {
		WriteOffset(buf[w.Size:], Offset(len(v)) * FieldSizes[TypeUint8])
	}
	w.Size += FieldSizes[TypeUint8Array]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeUint8Array)
	for _, vv := range v {
		w.WriteUint8(buf, vv)
	}
}

func (w *Builder) WriteUint16Array(buf []byte, v []uint16) {
	w.Size = alignOffsetToType(w.Size, TypeUint16Array)
	if buf != nil {
		WriteOffset(buf[w.Size:], Offset(len(v)) * FieldSizes[TypeUint16])
	}
	w.Size += FieldSizes[TypeUint16Array]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeUint16Array)
	for _, vv := range v {
		w.WriteUint16(buf, vv)
	}
}

func (w *Builder) WriteUint32Array(buf []byte, v []uint32) {
	w.Size = alignOffsetToType(w.Size, TypeUint32Array)
	if buf != nil {
		WriteOffset(buf[w.Size:], Offset(len(v)) * FieldSizes[TypeUint32])
	}
	w.Size += FieldSizes[TypeUint32Array]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeUint32Array)
	for _, vv := range v {
		w.WriteUint32(buf, vv)
	}
}

func (w *Builder) WriteUint64Array(buf []byte, v []uint64) {
	w.Size = alignOffsetToType(w.Size, TypeUint64Array)
	if buf != nil {
		WriteOffset(buf[w.Size:], Offset(len(v)) * FieldSizes[TypeUint64])
	}
	w.Size += FieldSizes[TypeUint64Array]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeUint64Array)
	for _, vv := range v {
		w.WriteUint64(buf, vv)
	}
}

func (w *Builder) WriteBytesArray(buf []byte, v [][]byte) {
	w.Size = alignOffsetToType(w.Size, TypeBytesArray)
	sizePlaceholderOffset := w.Size
	w.Size += FieldSizes[TypeBytesArray]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeBytesArray)
	contentSizeStartOffset := w.Size
	for _, vv := range v {
		w.WriteBytes(buf, vv)
	}
	contentSize := w.Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}

func (w *Builder) WriteStringArray(buf []byte, v []string) {
	w.Size = alignOffsetToType(w.Size, TypeStringArray)
	sizePlaceholderOffset := w.Size
	w.Size += FieldSizes[TypeStringArray]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeStringArray)
	contentSizeStartOffset := w.Size
	for _, vv := range v {
		w.WriteString(buf, vv)
	}
	contentSize := w.Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
}

func (w *Builder) WriteMessage(buf []byte, v MessageBuilder) (err error) {
	w.Size = alignOffsetToType(w.Size, TypeMessage)
	sizePlaceholderOffset := w.Size
	w.Size += FieldSizes[TypeMessage]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeMessage)
	if buf != nil {
		err = v.Write(buf[w.Size:])
	} else {
		err = v.Write(nil)
	}
	contentSize := v.GetSize()
	w.Size += contentSize
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
	return
}

func (w *Builder) WriteMessageArray(buf []byte, v []MessageBuilder) (err error) {
	w.Size = alignOffsetToType(w.Size, TypeMessageArray)
	sizePlaceholderOffset := w.Size
	w.Size += FieldSizes[TypeMessageArray]
	w.Size = alignDynamicFieldContentOffset(w.Size, TypeMessageArray)
	contentSizeStartOffset := w.Size
	for _, vv := range v {
		err = w.WriteMessage(buf, vv)
		if err != nil {
			return
		}
	}
	contentSize := w.Size - contentSizeStartOffset
	if buf != nil {
		WriteOffset(buf[sizePlaceholderOffset:], contentSize)
	}
	return nil
}
