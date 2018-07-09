package membuffers

import (
	"testing"
	"reflect"
)

var uint32Iterator = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	expected []uint32
}{
	{
		[]byte{0x0c,0x00,0x00,0x00, 0x13,0x00,0x00,0x00, 0x14,0x00,0x00,0x00, 0x15,0x00,0x00,0x00},
		[]FieldType{TypeUint32Array},
		[][]FieldType{{}},
		[]uint32{0x13,0x14,0x15},
	},
	{
		[]byte{0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33},
		[]FieldType{TypeUint32Array},
		[][]FieldType{{}},
		[]uint32{},
	},
	{
		[]byte{0x07,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33},
		[]FieldType{TypeUint32Array},
		[][]FieldType{{}},
		[]uint32{0x05, 0x00},
	},
}

func TestIteratorUint32(t *testing.T) {
	for tn, tt := range uint32Iterator {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		res := []uint32{}
		for i := m.GetUint32ArrayIterator(0); i.HasNext(); {
			res = append(res, i.NextUint32())
		}
		if !reflect.DeepEqual(tt.expected, res) {
			t.Fatalf("expected %v but got %v in test #%d", tt.expected, res, tn)
		}
	}
}

func TestIteratorUint8(t *testing.T) {
	buf := []byte{0x03,0x00,0x00,0x00, 0x13, 0x14, 0x15}
	scheme := []FieldType{TypeUint8Array}
	unions :=	[][]FieldType{{}}
	expected := []uint8{0x13,0x14,0x15}
	m := InternalMessage{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []uint8
	for i := m.GetUint8ArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextUint8())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}

func TestIteratorUint16(t *testing.T) {
	buf := []byte{0x06,0x00,0x00,0x00, 0x13,0x00, 0x14,0x00, 0x15,0x00}
	scheme := []FieldType{TypeUint16Array}
	unions :=	[][]FieldType{{}}
	expected := []uint16{0x13,0x14,0x15}
	m := InternalMessage{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []uint16
	for i := m.GetUint16ArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextUint16())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}

func TestIteratorUint64(t *testing.T) {
	buf := []byte{0x18,0x00,0x00,0x00, 0x13,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x14,0x00,0x00,0x00,0x00,0x00,0x00,0x00, 0x15,0x00,0x00,0x00,0x00,0x00,0x00,0x00}
	scheme := []FieldType{TypeUint64Array}
	unions :=	[][]FieldType{{}}
	expected := []uint64{0x13,0x14,0x15}
	m := InternalMessage{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []uint64
	for i := m.GetUint64ArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextUint64())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}

var messageIterator = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	expectedSizes []Offset
}{
	{
		[]byte{0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03},
		[]FieldType{TypeMessageArray},
		[][]FieldType{{}},
		[]Offset{0x03,0x05,0x03},
	},
	{
		[]byte{0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeMessageArray},
		[][]FieldType{{}},
		[]Offset{0},
	},
	{
		[]byte{0x08,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeMessageArray},
		[][]FieldType{{}},
		[]Offset{0},
	},
	{
		[]byte{0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeMessageArray},
		[][]FieldType{{}},
		[]Offset{},
	},
	{
		[]byte{0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44,0x55},
		[]FieldType{TypeMessageArray},
		[][]FieldType{{}},
		[]Offset{4,0},
	},
}

func TestIteratorMessage(t *testing.T) {
	for tn, tt := range messageIterator {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		res := []Offset{}
		for i := m.GetMessageArrayIterator(0); i.HasNext(); {
			_, size := i.NextMessage()
			res = append(res, size)
		}
		if !reflect.DeepEqual(tt.expectedSizes, res) {
			t.Fatalf("expected %v but got %v in test #%d", tt.expectedSizes, res, tn)
		}
	}
}

var bytesIterator = []struct{
	buf []byte
	scheme []FieldType
	unions [][]FieldType
	expected [][]byte
}{
	{
		[]byte{0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03},
		[]FieldType{TypeBytesArray},
		[][]FieldType{{}},
		[][]byte{{0x01,0x02,0x03},{0x01,0x02,0x03,0x04,0x05},{0x01,0x02,0x03}},
	},
	{
		[]byte{0x08,0x00,0x00,0x00, 0x88,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeBytesArray},
		[][]FieldType{{}},
		[][]byte{{}},
	},
	{
		[]byte{0x08,0x00,0x00,0x00, 0x05,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeBytesArray},
		[][]FieldType{{}},
		[][]byte{{}},
	},
	{
		[]byte{0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44},
		[]FieldType{TypeBytesArray},
		[][]FieldType{{}},
		[][]byte{},
	},
	{
		[]byte{0x09,0x00,0x00,0x00, 0x04,0x00,0x00,0x00, 0x11,0x22,0x33,0x44,0x55},
		[]FieldType{TypeBytesArray},
		[][]FieldType{{}},
		[][]byte{{0x11,0x22,0x33,0x44},{}},
	},
}

func TestIteratorBytes(t *testing.T) {
	for tn, tt := range bytesIterator {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		res := [][]byte{}
		for i := m.GetBytesArrayIterator(0); i.HasNext(); {
			res = append(res, i.NextBytes())
		}
		if !reflect.DeepEqual(tt.expected, res) {
			t.Fatalf("expected %v but got %v in test #%d", tt.expected, res, tn)
		}
	}
}

func TestIteratorString(t *testing.T) {
	buf := []byte{0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 'a','b','c',0x00, 0x05,0x00,0x00,0x00, 'h','e','l','l', 'o',0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 'd','e','f'}
	scheme := []FieldType{TypeStringArray}
	unions :=	[][]FieldType{{}}
	expected := []string{"abc","hello","def"}
	m := InternalMessage{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	res := []string{}
	for i := m.GetStringArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextString())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}