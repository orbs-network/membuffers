// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package membuffers

import (
	"bytes"
	"testing"
)

var rawBuffer = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	expected []byte
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		[]byte{0x00, 0x00, 0x00, 0x00},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
	},
}

func TestMessageRawBuffer(t *testing.T) {
	for tn, tt := range rawBuffer {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.RawBuffer()
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var rawBufferForField = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	unionNum int
	expected []byte
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		0,
		[]byte{},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint32},
		[][]FieldType{{}},
		0,
		0,
		[]byte{'a', 'b', 'c'},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x44, 0x33, 0x22, 0x11},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint16Array, TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x55, 0x66, 0x77, 0x88, 0x99, 0xaa},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint16Array, TypeUint32},
		[][]FieldType{{}},
		2,
		0,
		[]byte{0x44, 0x33, 0x22, 0x11},
	},
	{
		[]byte{0x11, 0x12, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		[]FieldType{TypeUint16, TypeMessage},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x01, 0x02, 0x03, 0x04, 0x05},
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		0,
		[]byte{0x01, 0x02, 0x03, 0x04, 0x05},
	},
	{
		[]byte{0x01, 0x00, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x33},
	},
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x33, 0x44, 0x55, 0x66},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x033, 0x44, 0x55, 0x66},
	},
	{
		[]byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x11, 0x22, 0x33},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{'a', 'b', 'c'},
	},
	{
		[]byte{0x04, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x11, 0x11, 0x22, 0x22, 0x33, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString, TypeUint16Array}},
		0,
		0,
		[]byte{0x11, 0x11, 0x22, 0x22, 0x33, 0x33},
	},
}

func TestMessageRawBufferForField(t *testing.T) {
	for tn, tt := range rawBufferForField {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.RawBufferForField(tt.fieldNum, tt.unionNum)
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var rawBufferWithHeaderForField = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	unionNum int
	expected []byte
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		0,
		[]byte{0x00, 0x00, 0x00, 0x00},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint32},
		[][]FieldType{{}},
		0,
		0,
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x44, 0x33, 0x22, 0x11},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint16Array, TypeUint32},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x06, 0x00, 0x00, 0x00, 0x55, 0x66, 0x77, 0x88, 0x99, 0xaa, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString, TypeUint16Array, TypeUint32},
		[][]FieldType{{}},
		2,
		0,
		[]byte{0x44, 0x33, 0x22, 0x11},
	},
	{
		[]byte{0x11, 0x12, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		[]FieldType{TypeUint16, TypeMessage},
		[][]FieldType{{}},
		1,
		0,
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		0,
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05},
	},
	{
		[]byte{0x01, 0x00, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x01, 0x00, 0x033},
	},
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x33, 0x44, 0x55, 0x66},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x00, 0x00, 0x00, 0x00, 0x033, 0x44, 0x55, 0x66},
	},
	{
		[]byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x02, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString}},
		0,
		0,
		[]byte{0x03, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
	},
	{
		[]byte{0x04, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x11, 0x11, 0x22, 0x22, 0x33, 0x33},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage, TypeString, TypeUint16Array}},
		0,
		0,
		[]byte{0x04, 0x00, 0x00, 0x00, 0x06, 0x00, 0x00, 0x00, 0x11, 0x11, 0x22, 0x22, 0x33, 0x33},
	},
}

func TestMessageRawBufferWithHeaderForField(t *testing.T) {
	for tn, tt := range rawBufferWithHeaderForField {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.RawBufferWithHeaderForField(tt.fieldNum, tt.unionNum)
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint32 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected uint32
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		0x11223344,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		0x55667788,
	},
	{
		[]byte{0x01, 0x00, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		0x11223344,
	},
	{
		[]byte{0x01, 0x01, 0x01, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint8, TypeUint8, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		3,
		0x11223344,
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeMessage, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		0x11223344,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		1,
		0,
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		0,
	},
}

func TestMessageReadUint32(t *testing.T) {
	for tn, tt := range readUint32 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint32(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint8 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected uint8
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x99},
		[]FieldType{TypeUint32, TypeUint8, TypeUint8},
		[][]FieldType{{}},
		2,
		0x99,
	},
}

func TestMessageReadUint8(t *testing.T) {
	for tn, tt := range readUint8 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint8(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint16 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected uint16
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x00, 0x99, 0xaa},
		[]FieldType{TypeUint32, TypeUint8, TypeUint16},
		[][]FieldType{{}},
		2,
		0xaa99,
	},
}

func TestMessageReadUint16(t *testing.T) {
	for tn, tt := range readUint16 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint16(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readUint64 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected uint64
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32, TypeUint64},
		[][]FieldType{{}},
		1,
		0x1122334455667788,
	},
	{
		[]byte{0x44, 0x00, 0x00, 0x00, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint8, TypeUint64},
		[][]FieldType{{}},
		1,
		0x1122334455667788,
	},
}

func TestMessageReadUint64(t *testing.T) {
	for tn, tt := range readUint64 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetUint64(tt.fieldNum)
		if tt.expected != s {
			t.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
		}
	}
}

var readBytes = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x22, 0x23, 0x24},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x22, 0x23, 0x24},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x22, 0x23, 0x24, 0x00, 0x02, 0x00, 0x00, 0x00, 0x77, 0x88},
		[]FieldType{TypeBytes, TypeBytes},
		[][]FieldType{{}},
		1,
		[]byte{0x77, 0x88},
	},
	{
		[]byte{0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x11, 0x22, 0x33},
	},
	{
		[]byte{0x01, 0x01, 0x01, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint8, TypeUint8, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		3,
		[]byte{0x11, 0x22, 0x33},
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeMessage, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x11, 0x22, 0x33},
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x11, 0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{},
	},
}

func TestMessageReadBytes(t *testing.T) {
	for tn, tt := range readBytes {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetBytes(tt.fieldNum)
		if !bytes.Equal(s, tt.expected) {
			t.Fatalf("expected %v but got %v in test #%d", tt.expected, s, tn)
		}
	}
}

var readString = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected string
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"abc",
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c', 0x00, 0x02, 0x00, 0x00, 0x00, 'd', 'e'},
		[]FieldType{TypeString, TypeString},
		[][]FieldType{{}},
		1,
		"de",
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		"",
	},
}

func TestMessageReadString(t *testing.T) {
	for tn, tt := range readString {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		s := m.GetString(tt.fieldNum)
		if s != tt.expected {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
		}
	}
}

var readMessage = []struct {
	buf          []byte
	scheme       []FieldType
	unions       [][]FieldType
	fieldNum     int
	expectedBuf  []byte
	expectedSize Offset
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{0x01, 0x02, 0x03},
		3,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x00, 0x02, 0x00, 0x00, 0x00, 0x04, 0x05},
		[]FieldType{TypeMessage, TypeMessage},
		[][]FieldType{{}},
		1,
		[]byte{0x04, 0x05},
		2,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x01, 0x02},
		[]FieldType{TypeMessage},
		[][]FieldType{{}},
		0,
		[]byte{},
		0,
	},
}

func TestMessageReadMessage(t *testing.T) {
	for tn, tt := range readMessage {
		m := InternalMessage{}
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

var readUnion = []struct {
	buf         []byte
	scheme      []FieldType
	unions      [][]FieldType
	fieldNum    int
	unionNum    int
	unionIndex  uint16
	expectedIs  bool
	expectedOff Offset
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}},
		0,
		0,
		0,
		true,
		4,
	},
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}},
		0,
		0,
		1,
		false,
		4,
	},
	{
		[]byte{0x01, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04},
		[]FieldType{TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}},
		0,
		0,
		1,
		true,
		2,
	},
	{
		[]byte{0x01, 0x00, 0x11, 0x00, 0x00, 0x00, 0x22, 0x23},
		[]FieldType{TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}, {TypeUint16, TypeUint8, TypeUint32}},
		1,
		1,
		0,
		true,
		6,
	},
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44, 0x00, 0x00, 0x22, 0x23},
		[]FieldType{TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}, {TypeUint16, TypeUint8, TypeUint32}},
		1,
		1,
		0,
		true,
		10,
	},
	{
		[]byte{0x00, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x44, 0x02, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04},
		[]FieldType{TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8}, {TypeUint16, TypeUint8, TypeUint32}},
		1,
		1,
		2,
		true,
		12,
	},
	{
		[]byte{0x22, 0x22, 0x22, 0x22, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x01, 0x00, 0x17},
		[]FieldType{TypeUint32, TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage}, {TypeUint16, TypeUint8, TypeUint32}},
		2,
		1,
		1,
		true,
		20,
	},
	{
		[]byte{0x22, 0x22, 0x22, 0x22, 0x07, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x01, 0x00, 0x17},
		[]FieldType{TypeUint32, TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage}, {TypeUint16, TypeUint8, TypeUint32}},
		2,
		1,
		1,
		false,
		0,
	},
	{
		[]byte{0x22, 0x22, 0x22, 0x22, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x01},
		[]FieldType{TypeUint32, TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage}, {TypeUint16, TypeUint8, TypeUint32}},
		2,
		1,
		1,
		false,
		0,
	},
	{
		[]byte{0x22, 0x22, 0x22, 0x22, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x01, 0x00},
		[]FieldType{TypeUint32, TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage}, {TypeUint16, TypeUint8, TypeUint32}},
		2,
		1,
		1,
		false,
		0,
	},
	{
		[]byte{0x22, 0x22, 0x22, 0x22, 0x02, 0x00, 0x00, 0x00, 0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x11},
		[]FieldType{TypeUint32, TypeUnion, TypeUnion},
		[][]FieldType{{TypeUint32, TypeUint8, TypeMessage}, {TypeUint16, TypeUint8, TypeUint32}},
		2,
		1,
		1,
		false,
		0,
	},
}

func TestMessageReadUnion(t *testing.T) {
	for tn, tt := range readUnion {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		is, off := m.IsUnionIndex(tt.fieldNum, tt.unionNum, tt.unionIndex)
		if is != tt.expectedIs {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expectedIs, is, tn)
		}
		if off != tt.expectedOff {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expectedOff, off, tn)
		}
	}
}

var mutateUint32 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		[]byte{0x55, 0x55, 0x55, 0x55},
		false,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x44, 0x33, 0x22, 0x11, 0x55, 0x55, 0x55, 0x55},
		false,
	},
	{
		[]byte{0x01, 0x00, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x01, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x55, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x01, 0x01, 0x01, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint8, TypeUint8, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		3,
		[]byte{0x01, 0x01, 0x01, 0x00, 0x55, 0x55, 0x55, 0x55, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeMessage, TypeUint32, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x55, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x44, 0x33, 0x22, 0x11},
		true,
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeUint32},
		[][]FieldType{{}},
		0,
		[]byte{0x44, 0x33, 0x22},
		true,
	},
}

func TestMessageMutateUint32(t *testing.T) {
	for tn, tt := range mutateUint32 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetUint32(tt.fieldNum, 0x55555555)
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

var mutateUint8 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x99},
		[]FieldType{TypeUint32, TypeUint8, TypeUint8},
		[][]FieldType{{}},
		2,
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x55},
		false,
	},
}

func TestMessageMutateUint8(t *testing.T) {
	for tn, tt := range mutateUint8 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetUint8(tt.fieldNum, 0x55)
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

var mutateUint16 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x00, 0x99, 0xaa},
		[]FieldType{TypeUint32, TypeUint8, TypeUint16},
		[][]FieldType{{}},
		2,
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x00, 0x55, 0x55},
		false,
	},
}

func TestMessageMutateUint16(t *testing.T) {
	for tn, tt := range mutateUint16 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetUint16(tt.fieldNum, 0x5555)
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

var mutateUint64 = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x44, 0x33, 0x22, 0x11, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint32, TypeUint64},
		[][]FieldType{{}},
		1,
		[]byte{0x44, 0x33, 0x22, 0x11, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55},
		false,
	},
	{
		[]byte{0x44, 0x00, 0x00, 0x00, 0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeUint8, TypeUint64},
		[][]FieldType{{}},
		1,
		[]byte{0x44, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55},
		false,
	},
}

func TestMessageMutateUint64(t *testing.T) {
	for tn, tt := range mutateUint64 {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetUint64(tt.fieldNum, 0x5555555555555555)
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

var mutateBytes = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x00, 0x00, 0x00, 0x00},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x00, 0x00, 0x00, 0x00},
		true,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x22, 0x23, 0x24},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x03, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55},
		false,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x22, 0x23, 0x24, 0x00, 0x02, 0x00, 0x00, 0x00, 0x77, 0x88},
		[]FieldType{TypeBytes, TypeBytes},
		[][]FieldType{{}},
		1,
		[]byte{0x03, 0x00, 0x00, 0x00, 0x22, 0x23, 0x24, 0x00, 0x02, 0x00, 0x00, 0x00, 0x77, 0x88},
		true,
	},
	{
		[]byte{0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x01, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x00, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x01, 0x01, 0x01, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeUint8, TypeUint8, TypeUint8, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		3,
		[]byte{0x01, 0x01, 0x01, 0x00, 0x03, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x00, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x11, 0x22, 0x33, 0x00, 0x88, 0x77, 0x66, 0x55},
		[]FieldType{TypeMessage, TypeBytes, TypeUint32},
		[][]FieldType{{}},
		1,
		[]byte{0x05, 0x00, 0x00, 0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x00, 0x00, 0x00, 0x03, 0x00, 0x00, 0x00, 0x55, 0x55, 0x55, 0x00, 0x88, 0x77, 0x66, 0x55},
		false,
	},
	{
		[]byte{0x44, 0x33, 0x22, 0x11},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x44, 0x33, 0x22, 0x11},
		true,
	},
	{
		[]byte{0x44, 0x33, 0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x44, 0x33, 0x22},
		true,
	},
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 0x11, 0x22},
		[]FieldType{TypeBytes},
		[][]FieldType{{}},
		0,
		[]byte{0x03, 0x00, 0x00, 0x00, 0x11, 0x22},
		true,
	},
}

func TestMessageMutateBytes(t *testing.T) {
	for tn, tt := range mutateBytes {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetBytes(tt.fieldNum, []byte{0x55, 0x55, 0x55})
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

var mutateString = []struct {
	buf      []byte
	scheme   []FieldType
	unions   [][]FieldType
	fieldNum int
	expected []byte
	err      bool
}{
	{
		[]byte{0x03, 0x00, 0x00, 0x00, 'a', 'b', 'c'},
		[]FieldType{TypeString},
		[][]FieldType{{}},
		0,
		[]byte{0x03, 0x00, 0x00, 0x00, 'z', 'z', 'z'},
		false,
	},
}

func TestMessageMutateString(t *testing.T) {
	for tn, tt := range mutateString {
		m := InternalMessage{}
		m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
		err := m.SetString(tt.fieldNum, "zzz")
		if !bytes.Equal(tt.expected, tt.buf) {
			t.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, tt.buf, tn)
		}
		if (err != nil) != tt.err {
			t.Fatalf("expected error \"%v\" but got \"%v\" in test #%d", err != nil, tt.err, tn)
		}
	}
}

func BenchmarkUint32Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for tn, tt := range readUint32 {
			m := InternalMessage{}
			m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
			s := m.GetUint32(tt.fieldNum)
			if tt.expected != s {
				b.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
			}
		}
	}
}

func BenchmarkUint64Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for tn, tt := range readUint64 {
			m := InternalMessage{}
			m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
			s := m.GetUint64(tt.fieldNum)
			if tt.expected != s {
				b.Fatalf("expected 0x%x but got 0x%0x in test #%d", tt.expected, s, tn)
			}
		}
	}
}

func BenchmarkSingleUint64Read(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := readUint64[1]
		m := InternalMessage{}
		m.Init(x.buf, Offset(len(x.buf)), x.scheme, x.unions)
		s := m.GetUint64(x.fieldNum)
		if x.expected != s {
			b.Fatalf("expected 0x%x but got 0x%0x", x.expected, s)
		}
	}
}

func BenchmarkStringRead(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for tn, tt := range readString {
			m := InternalMessage{}
			m.Init(tt.buf, Offset(len(tt.buf)), tt.scheme, tt.unions)
			s := m.GetString(tt.fieldNum)
			if s != tt.expected {
				b.Fatalf("expected \"%v\" but got \"%v\" in test #%d", tt.expected, s, tn)
			}
		}
	}
}
