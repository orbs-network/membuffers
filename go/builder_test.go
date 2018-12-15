package membuffers

import (
	"bytes"
	"testing"
)

func TestBuilderUint8(t *testing.T) {
	w := InternalBuilder{}
	v := uint8(0x17)
	w.WriteUint8(nil, v)
	if w.size != 1 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint8(buf, v)
	expected := []byte{0x17}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint16(t *testing.T) {
	w := InternalBuilder{}
	v := uint16(0x17)
	w.WriteUint16(nil, v)
	if w.size != 2 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint16(buf, v)
	expected := []byte{0x17, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint32(t *testing.T) {
	w := InternalBuilder{}
	v := uint32(0x17)
	w.WriteUint32(nil, v)
	if w.size != 4 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint32(buf, v)
	expected := []byte{0x17, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint64(t *testing.T) {
	w := InternalBuilder{}
	v := uint64(0x17)
	w.WriteUint64(nil, v)
	if w.size != 8 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint64(buf, v)
	expected := []byte{0x17, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderBytes(t *testing.T) {
	w := InternalBuilder{}
	v := []byte{0x01, 0x02, 0x03}
	w.WriteBytes(nil, v)
	if w.size != 7 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	w.WriteBytes(nil, v)
	if w.size != 15 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteBytes(buf, v)
	expected := []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03}
	if !bytes.Equal(buf[:7], expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
	w.WriteBytes(buf, v)
	expected2 := []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03}
	if !bytes.Equal(buf, expected2) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected2, buf)
	}
}

func TestBuilderString(t *testing.T) {
	w := InternalBuilder{}
	v := "hello"
	w.WriteString(nil, v)
	if w.size != 9 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	w.WriteString(nil, v)
	if w.size != 21 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteString(buf, v)
	expected := []byte{0x05, 0x00, 0x00, 0x00, 'h', 'e', 'l', 'l', 'o'}
	if !bytes.Equal(buf[:9], expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
	w.WriteString(buf, v)
	expected2 := []byte{0x05, 0x00, 0x00, 0x00, 'h', 'e', 'l', 'l', 'o', 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 'h', 'e', 'l', 'l', 'o'}
	if !bytes.Equal(buf, expected2) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected2, buf)
	}
}

func TestBuilderUnionIndex(t *testing.T) {
	w := InternalBuilder{}
	v := uint16(0x01)
	w.WriteUnionIndex(nil, v)
	if w.size != 2 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUnionIndex(buf, v)
	expected := []byte{0x01, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint8Array(t *testing.T) {
	w := InternalBuilder{}
	v := []uint8{0x01, 0x02, 0x03}
	w.WriteUint8Array(nil, v)
	if w.size != 7 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint8Array(buf, v)
	expected := []byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint16Array(t *testing.T) {
	w := InternalBuilder{}
	v := []uint16{0x01, 0x02, 0x03}
	w.WriteUint16Array(nil, v)
	if w.size != 10 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint16Array(buf, v)
	expected := []byte{0x06, 0x00, 0x00, 0x00, 0x01, 0x00, 0x02, 0x00, 0x03, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint32Array(t *testing.T) {
	w := InternalBuilder{}
	v := []uint32{0x01, 0x02, 0x03}
	w.WriteUint32Array(nil, v)
	if w.size != 16 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint32Array(buf, v)
	expected := []byte{0x0c, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderUint64Array(t *testing.T) {
	w := InternalBuilder{}
	v := []uint64{0x01, 0x02, 0x03}
	w.WriteUint64Array(nil, v)
	if w.size != 28 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteUint64Array(buf, v)
	expected := []byte{0x18, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderBytesArray(t *testing.T) {
	w := InternalBuilder{}
	v := [][]byte{{0x01, 0x02, 0x03}, {0x04, 0x05}}
	w.WriteBytesArray(nil, v)
	if w.size != 18 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteBytesArray(buf, v)
	expected := []byte{0x0e, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x05}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}

func TestBuilderStringArray(t *testing.T) {
	w := InternalBuilder{}
	v := []string{"jay", "lo"}
	w.WriteStringArray(nil, v)
	if w.size != 18 {
		t.Fatalf("instead of expected size got %v", w.size)
	}
	buf := make([]byte, w.size)
	w.Reset()
	w.WriteStringArray(buf, v)
	expected := []byte{0x0e, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 'j', 'a', 'y', 0x00, 0x02, 0x00, 0x00, 0x00, 'l', 'o'}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("expected \"%v\" but got \"%v\"", expected, buf)
	}
}
