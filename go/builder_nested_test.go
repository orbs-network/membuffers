package membuffers

import (
	"bytes"
	"testing"
)

// an example nested builder

type ExampleMessageBuilder struct {
	Field1 uint8
	Field2 uint32

	// internal
	builder               InternalBuilder
	overrideWithRawBuffer []byte
}

func (w *ExampleMessageBuilder) Write(buf []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &ErrBufferOverrun{}
		}
	}()
	if w.overrideWithRawBuffer != nil {
		return w.builder.WriteOverrideWithRawBuffer(buf, w.overrideWithRawBuffer)
	}
	w.builder.Reset()
	w.builder.WriteUint8(buf, w.Field1)
	w.builder.WriteUint32(buf, w.Field2)
	return nil
}

func exampleMessageBuilderFromRaw(raw []byte) *ExampleMessageBuilder {
	return &ExampleMessageBuilder{overrideWithRawBuffer: raw}
}

func (w *ExampleMessageBuilder) HexDump(prefix string, offsetFromStart Offset) (err error) {
	return nil
}

func (w *ExampleMessageBuilder) GetSize() Offset {
	return w.builder.GetSize()
}

func (w *ExampleMessageBuilder) CalcRequiredSize() Offset {
	w.Write(nil)
	return w.builder.GetSize()
}

// tests start here

func TestBuilderMessage(t *testing.T) {
	w := InternalBuilder{}
	v := ExampleMessageBuilder{Field1: 0x17, Field2: 0x33}
	w.WriteMessage(nil, &v)
	if w.size != 12 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteMessage(buf, &v)
	expected := []byte{0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderMessageArray(t *testing.T) {
	w := InternalBuilder{}
	v := []MessageWriter{&ExampleMessageBuilder{Field1: 0x17, Field2: 0x33}, &ExampleMessageBuilder{Field1: 0x17, Field2: 0x33}}
	w.WriteMessageArray(nil, v)
	if w.size != 28 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteMessageArray(buf, v)
	expected := []byte{0x18, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x33, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderMessageWithRawOverride(t *testing.T) {
	// generate a raw buffer on the side with the builder result as raw
	src := &ExampleMessageBuilder{Field1: 0x19, Field2: 0x88}
	src.Write(nil)
	rawBufForBuilder := make([]byte, src.GetSize())
	src.Write(rawBufForBuilder)
	src = nil

	// start the test
	w := InternalBuilder{}
	v := exampleMessageBuilderFromRaw(rawBufForBuilder)
	w.WriteMessage(nil, v)
	if w.size != 12 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteMessage(buf, v)
	expected := []byte{0x08, 0x00, 0x00, 0x00, 0x19, 0x00, 0x00, 0x00, 0x88, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}
