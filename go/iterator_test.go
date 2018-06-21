package membuffers

import (
	"testing"
	"reflect"
)

func TestIteratorUint32(t *testing.T) {
	buf := []byte{0x0c,0x00,0x00,0x00, 0x13,0x00,0x00,0x00, 0x14,0x00,0x00,0x00, 0x15,0x00,0x00,0x00}
	scheme := []FieldType{TypeUint32Array}
	unions :=	[][]FieldType{{}}
	expected := []uint32{0x13,0x14,0x15}
	m := Message{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []uint32
	for i := m.GetUint32ArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextUint32())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}

func TestIteratorUint8(t *testing.T) {
	buf := []byte{0x03,0x00,0x00,0x00, 0x13, 0x14, 0x15}
	scheme := []FieldType{TypeUint8Array}
	unions :=	[][]FieldType{{}}
	expected := []uint8{0x13,0x14,0x15}
	m := Message{}
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
	m := Message{}
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
	m := Message{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []uint64
	for i := m.GetUint64ArrayIterator(0); i.HasNext(); {
		res = append(res, i.NextUint64())
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}

func TestIteratorMessage(t *testing.T) {
	buf := []byte{0x1b,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03,0x00, 0x05,0x00,0x00,0x00, 0x01,0x02,0x03,0x04, 0x05,0x00,0x00,0x00, 0x03,0x00,0x00,0x00, 0x01,0x02,0x03}
	scheme := []FieldType{TypeMessageArray}
	unions :=	[][]FieldType{{}}
	expected := []Offset{0x03,0x05,0x03}
	m := Message{}
	m.Init(buf, Offset(len(buf)), scheme, unions)
	var res []Offset
	for i := m.GetMessageArrayIterator(0); i.HasNext(); {
		_, size := i.NextMessage()
		res = append(res, size)
	}
	if !reflect.DeepEqual(expected, res) {
		t.Fatalf("expected %v but got %v", expected, res)
	}
}
