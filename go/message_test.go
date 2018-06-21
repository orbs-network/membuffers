package membuffers

import (
	"testing"
	"bytes"
)

var rawBuffer = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	expected []byte
}{
	{
		[]byte{0x00,0x00,0x00,0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		[]byte{0x00,0x00,0x00,0x00},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c'},
	},
}

func TestMessageRawBuffer(t *testing.T) {
	for tn, tt := range rawBuffer {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.RawBuffer()
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var rawBufferForField = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	unionNum int
	expected []byte
}{
	{
		[]byte{0x00,0x00,0x00,0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		0,
		[]byte{},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c', 0x00, 0x44,0x33,0x22,0x11},
		[]FieldType{TypeString,TypeUint32},
		[][]FieldType{{}},
		0,
		0,
		[]byte{'a','b','c'},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c', 0x00, 0x44,0x33,0x22,0x11},
		[]FieldType{TypeString,TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x44,0x33,0x22,0x11},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c', 0x00, 0x06,0x00,0x00,0x00, 0x55,0x66, 0x77,0x88, 0x99,0xaa, 0x00,0x00, 0x44,0x33,0x22,0x11},
		[]FieldType{TypeString,TypeUint16Array,TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x55,0x66,0x77,0x88,0x99,0xaa},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c', 0x00, 0x06,0x00,0x00,0x00, 0x55,0x66, 0x77,0x88, 0x99,0xaa, 0x00,0x00, 0x44,0x33,0x22,0x11},
		[]FieldType{TypeString,TypeUint16Array,TypeUint32},
		[][]FieldType{{}},
		2,
		0,
		[]byte{0x44,0x33,0x22,0x11},
	},
	{
		[]byte{0x11,0x12, 0x00,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05},
		[]FieldType{TypeUint16,TypeMessage},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x01,0x02,0x03,0x04,0x05},
	},
	{
		[]byte{0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		0,
		[]byte{0x01,0x02,0x03,0x04,0x05},
	},
}

func TestMessageRawBufferForField(t *testing.T) {
	for tn, tt := range rawBufferForField {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.RawBufferForField(tt.fieldNum, tt.unionNum)
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint32 = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected uint32
}{
	{
		[]byte{0x44,0x33,0x22,0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		0x11223344,
	},
	{
		[]byte{0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeUint32,TypeUint32},
		[][]FieldType{{}},
		1,
		0x55667788,
	},
	{
		[]byte{0x01, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeUint8,TypeUint32,TypeUint32},
		[][]FieldType{{}},
		1,
		0x11223344,
	},
	{
		[]byte{0x01,0x01,0x01, 0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeUint8,TypeUint8,TypeUint8,TypeUint32,TypeUint32},
		[][]FieldType{{}},
		3,
		0x11223344,
	},
	{
		[]byte{0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeMessage,TypeUint32,TypeUint32},
		[][]FieldType{{}},
		1,
		0x11223344,
	},
	{
		[]byte{0x44,0x33,0x22,0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		1,
		0,
	},
	{
		[]byte{0x44,0x33,0x22},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		0,
	},
}

func TestMessageReadUint32(t *testing.T) {
	for tn, tt := range readUint32 {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint32(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint8 = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected uint8
}{
	{
		[]byte{0x44,0x33,0x22,0x11,0x88,0x99},
		[]FieldType{TypeUint32,TypeUint8,TypeUint8},
		[][]FieldType{{}},
		2,
		0x99,
	},
}

func TestMessageReadUint8(t *testing.T) {
	for tn, tt := range readUint8 {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint8(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint16 = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected uint16
}{
	{
		[]byte{0x44,0x33,0x22,0x11,0x88,0x00,0x99,0xaa},
		[]FieldType{TypeUint32,TypeUint8,TypeUint16},
		[][]FieldType{{}},
		2,
		0xaa99,
	},
}

func TestMessageReadUint16(t *testing.T) {
	for tn, tt := range readUint16 {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint16(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint64 = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected uint64
}{
	{
		[]byte{0x44,0x33,0x22,0x11, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11},
		[]FieldType{TypeUint32,TypeUint64},
		[][]FieldType{{}},
		1,
		0x1122334455667788,
	},
	{
		[]byte{0x44,0x00,0x00,0x00, 0x88,0x77,0x66,0x55,0x44,0x33,0x22,0x11},
		[]FieldType{TypeUint8,TypeUint64},
		[][]FieldType{{}},
		1,
		0x1122334455667788,
	},
}

func TestMessageReadUint64(t *testing.T) {
	for tn, tt := range readUint64 {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint64(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readBytes = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected []byte
}{
	{
		[]byte{0x00,0x00,0x00,0x00},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x22,0x23,0x24},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x22,0x23,0x24},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x22,0x23,0x24, 0x00, 0x02,0x00,0x00,0x00, 0x77,0x88},
		[]FieldType{TypeBytes,TypeBytes},
		[][]FieldType{{}},
		1,
		[]byte{0x77,0x88},
	},
	{
		[]byte{0x01, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeUint8,TypeBytes,TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x11,0x22,0x33},
	},
	{
		[]byte{0x01,0x01,0x01, 0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeUint8,TypeUint8,TypeUint8,TypeBytes,TypeUint32},
		[][]FieldType{{}},
		3,
		[]byte{0x11,0x22,0x33},
	},
	{
		[]byte{0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04,0x05, 0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x11,0x22,0x33, 0x00, 0x88,0x77,0x66,0x55},
		[]FieldType{TypeMessage,TypeBytes,TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x11,0x22,0x33},
	},
	{
		[]byte{0x44,0x33,0x22,0x11},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x44,0x33,0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x11,0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
}

func TestMessageReadBytes(t *testing.T) {
	for tn, tt := range readBytes {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetBytes(tt.fieldNum)
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected %v but got %v in test #%d", tt.expected, s, tn)
		}
	}
}

var readString = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expected string
}{
	{
		[]byte{0x00,0x00,0x00,0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"abc",
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b','c', 0x00, 0x02,0x00,0x00,0x00, 'd','e'},
		[]FieldType{TypeString,TypeString},
		[][]FieldType{{}},
		1,
		"de",
	},
	{
		[]byte{0x44,0x33,0x22,0x11},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x44,0x33,0x22},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 'a','b'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
}

func TestMessageReadString(t *testing.T) {
	for tn, tt := range readString {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetString(tt.fieldNum)
		if s != tt.expected {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var readMessage = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	fieldNum int
	expectedBuf []byte
	expectedSize Offset
}{
	{
		[]byte{0x00,0x00,0x00,0x00},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x01,0x02,0x03},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{0x01,0x02,0x03},
		3,
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x01,0x02,0x03, 0x00, 0x02,0x00,0x00,0x00, 0x04,0x05},
		[]FieldType{TypeMessage,TypeMessage},
		[][]FieldType{{}},
		1,
		[]byte{0x04,0x05},
		2,
	},
	{
		[]byte{0x44,0x33,0x22,0x11},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x44,0x33,0x22},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x03,0x00,0x00,0x00, 0x01,0x02},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
}

func TestMessageReadMessage(t *testing.T) {
	for tn, tt := range readMessage {
		m := Message{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		b, s := m.GetMessage(tt.fieldNum)
		if !bytes.Equal(b, tt.expectedBuf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expectedBuf, b, tn)
		}
		if s != tt.expectedSize {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expectedSize, s, tn)
		}
	}
}
