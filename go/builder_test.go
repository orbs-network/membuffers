package membuffers

import (
	"testing"
	"bytes"
)

func TestBuilderUint8(t *testing.T) {
	w := Builder{}
	v := uint8(0x17)
	w.WriteUint8(nil, v)
	if w.Size != 1 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint8(buf, v)
	expected := []byte{0x17}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint16(t *testing.T) {
	w := Builder{}
	v := uint16(0x17)
	w.WriteUint16(nil, v)
	if w.Size != 2 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint16(buf, v)
	expected := []byte{0x17,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint32(t *testing.T) {
	w := Builder{}
	v := uint32(0x17)
	w.WriteUint32(nil, v)
	if w.Size != 4 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint32(buf, v)
	expected := []byte{0x17,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint64(t *testing.T) {
	w := Builder{}
	v := uint64(0x17)
	w.WriteUint64(nil, v)
	if w.Size != 8 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint64(buf, v)
	expected := []byte{0x17,0x00,0x00,0x00,0x00,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderBytes(t *testing.T) {
	w := Builder{}
	v := []byte{0x01,0x02,0x03}
	w.WriteBytes(nil, v)
	if w.Size != 7 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	w.WriteBytes(nil, v)
	if w.Size != 15 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteBytes(buf, v)
	expected := []byte{0x03,0x00,0x00,0x00, 0x01,0x02,0x03}
	if !bytes.Equal(buf[:7], expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
	w.WriteBytes(buf, v)
	expected2 := []byte{0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,}
	if !bytes.Equal(buf, expected2) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected2, buf)
	}
}

func TestBuilderString(t *testing.T) {
	w := Builder{}
	v := "hello"
	w.WriteString(nil, v)
	if w.Size != 9 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	w.WriteString(nil, v)
	if w.Size != 21 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteString(buf, v)
	expected := []byte{0x05,0x00,0x00,0x00, 'h','e','l','l','o'}
	if !bytes.Equal(buf[:9], expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
	w.WriteString(buf, v)
	expected2 := []byte{0x05,0x00,0x00,0x00, 'h','e','l','l','o',0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 'h','e','l','l','o'}
	if !bytes.Equal(buf, expected2) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected2, buf)
	}
}

func TestBuilderUnionIndex(t *testing.T) {
	w := Builder{}
	v := uint16(0x01)
	w.WriteUnionIndex(nil, v)
	if w.Size != 2 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUnionIndex(buf, v)
	expected := []byte{0x01,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint8Array(t *testing.T) {
	w := Builder{}
	v := []uint8{0x01,0x02,0x03}
	w.WriteUint8Array(nil, v)
	if w.Size != 7 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint8Array(buf, v)
	expected := []byte{0x03,0x00,0x00,0x00, 0x01,0x02,0x03}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint16Array(t *testing.T) {
	w := Builder{}
	v := []uint16{0x01,0x02,0x03}
	w.WriteUint16Array(nil, v)
	if w.Size != 10 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint16Array(buf, v)
	expected := []byte{0x06,0x00,0x00,0x00, 0x01,0x00, 0x02,0x00, 0x03,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint32Array(t *testing.T) {
	w := Builder{}
	v := []uint32{0x01,0x02,0x03}
	w.WriteUint32Array(nil, v)
	if w.Size != 16 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint32Array(buf, v)
	expected := []byte{0x0c,0x00,0x00,0x00, 0x01,0x00,0x00,0x00, 0x02,0x00,0x00,0x00, 0x03,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint64Array(t *testing.T) {
	w := Builder{}
	v := []uint64{0x01,0x02,0x03}
	w.WriteUint64Array(nil, v)
	if w.Size != 28 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteUint64Array(buf, v)
	expected := []byte{0x18,0x00,0x00,0x00, 0x01,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x02,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x03,0x00,0x00,0x00,0x00,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderBytesArray(t *testing.T) {
	w := Builder{}
	v := [][]byte{{0x01,0x02,0x03},{0x04,0x05}}
	w.WriteBytesArray(nil, v)
	if w.Size != 18 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteBytesArray(buf, v)
	expected := []byte{0x0e,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x02,0x00,0x00,0x00, 0x04,0x05}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderStringArray(t *testing.T) {
	w := Builder{}
	v := []string{"jay","lo"}
	w.WriteStringArray(nil, v)
	if w.Size != 18 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteStringArray(buf, v)
	expected := []byte{0x0e,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 'j','a','y',0x00, 0x02,0x00,0x00,0x00, 'l','o'}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

type ExampleMessageBuilder struct {
	builder Builder
}
func (w *ExampleMessageBuilder) Write(buf []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = &ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteUint8(buf, 0x17)
	w.builder.WriteUint32(buf, 0x033)
	return nil
}
func (w *ExampleMessageBuilder) GetSize() Offset {
	return w.builder.GetSize()
}
func (w *ExampleMessageBuilder) CalcRequiredSize() Offset {
	w.Write(nil)
	return w.builder.GetSize()
}

func TestBuilderMessage(t *testing.T) {
	w := Builder{}
	v := ExampleMessageBuilder{}
	w.WriteMessage(nil, &v)
	if w.Size != 12 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteMessage(buf, &v)
	expected := []byte{0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderMessageArray(t *testing.T) {
	w := Builder{}
	v := []MessageBuilder{&ExampleMessageBuilder{}, &ExampleMessageBuilder{}}
	w.WriteMessageArray(nil, v)
	if w.Size != 28 {
		t.Fatalf("instead of expected size got %v", w.Size)
	}
	buf := make([]byte, w.Size)
	w.Reset()
	w.WriteMessageArray(buf, v)
	expected := []byte{0x18,0x00,0x00,0x00, 0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00, 0x08,0x00,0x00,0x00, 0x17,0x00,0x00,0x00, 0x33,0x00,0x00,0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}
