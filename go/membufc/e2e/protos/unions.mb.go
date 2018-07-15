// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package types

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
// message ExampleMessage

// reader

type ExampleMessage struct {
	// Str string

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ExampleMessage) String() string {
	return fmt.Sprintf("{Str:%s,}", x.StringStr())
}

var _ExampleMessage_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _ExampleMessage_Unions = [][]membuffers.FieldType{}

func ExampleMessageReader(buf []byte) *ExampleMessage {
	x := &ExampleMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ExampleMessage_Scheme, _ExampleMessage_Unions)
	return x
}

func (x *ExampleMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *ExampleMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ExampleMessage) Str() string {
	return x._message.GetString(0)
}

func (x *ExampleMessage) RawStr() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ExampleMessage) MutateStr(v string) error {
	return x._message.SetString(0, v)
}

func (x *ExampleMessage) StringStr() string {
	return fmt.Sprintf(x.Str())
}

// builder

type ExampleMessageBuilder struct {
	Str string

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *ExampleMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Str)
	return nil
}

func (w *ExampleMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ExampleMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ExampleMessageBuilder) Build() *ExampleMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ExampleMessageReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ComplexUnion

// reader

type ComplexUnion struct {
	// Option ComplexUnionOption

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ComplexUnion) String() string {
	return fmt.Sprintf("{Option:%s,}", x.StringOption())
}

var _ComplexUnion_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var _ComplexUnion_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeMessage,membuffers.TypeUint16,}}

func ComplexUnionReader(buf []byte) *ComplexUnion {
	x := &ComplexUnion{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ComplexUnion_Scheme, _ComplexUnion_Unions)
	return x
}

func (x *ComplexUnion) IsValid() bool {
	return x._message.IsValid()
}

func (x *ComplexUnion) Raw() []byte {
	return x._message.RawBuffer()
}

type ComplexUnionOption uint16

const (
	COMPLEX_UNION_OPTION_NUM ComplexUnionOption = 0
	COMPLEX_UNION_OPTION_MSG ComplexUnionOption = 1
	COMPLEX_UNION_OPTION_ENU ComplexUnionOption = 2
)

func (x *ComplexUnion) Option() ComplexUnionOption {
	return ComplexUnionOption(x._message.GetUint16(0))
}

func (x *ComplexUnion) IsOptionNum() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *ComplexUnion) Num() uint32 {
	_, off := x._message.IsUnionIndex(0, 0, 0)
	return x._message.GetUint32InOffset(off)
}

func (x *ComplexUnion) StringNum() string {
	return fmt.Sprintf("%x", x.Num())
}

func (x *ComplexUnion) MutateNum(v uint32) error {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *ComplexUnion) IsOptionMsg() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *ComplexUnion) Msg() *ExampleMessage {
	_, off := x._message.IsUnionIndex(0, 0, 1)
	b, s := x._message.GetMessageInOffset(off)
	return ExampleMessageReader(b[:s])
}

func (x *ComplexUnion) StringMsg() string {
	return x.Msg().String()
}

func (x *ComplexUnion) IsOptionEnu() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *ComplexUnion) Enu() ExampleEnum {
	_, off := x._message.IsUnionIndex(0, 0, 2)
	return ExampleEnum(x._message.GetUint16InOffset(off))
}

func (x *ComplexUnion) StringEnu() string {
	return x.Enu().String()
}

func (x *ComplexUnion) MutateEnu(v ExampleEnum) error {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint16InOffset(off, uint16(v))
	return nil
}

func (x *ComplexUnion) RawOption() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ComplexUnion) StringOption() string {
	switch x.Option() {
	case COMPLEX_UNION_OPTION_NUM:
		return "(Num)" + x.StringNum()
	case COMPLEX_UNION_OPTION_MSG:
		return "(Msg)" + x.StringMsg()
	case COMPLEX_UNION_OPTION_ENU:
		return "(Enu)" + x.StringEnu()
	}
	return "(Unknown)"
}

// builder

type ComplexUnionBuilder struct {
	Option ComplexUnionOption
	Num uint32
	Msg *ExampleMessageBuilder
	Enu ExampleEnum

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *ComplexUnionBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Option))
	switch w.Option {
	case COMPLEX_UNION_OPTION_NUM:
		w._builder.WriteUint32(buf, w.Num)
	case COMPLEX_UNION_OPTION_MSG:
		w._builder.WriteMessage(buf, w.Msg)
	case COMPLEX_UNION_OPTION_ENU:
		w._builder.WriteUint16(buf, uint16(w.Enu))
	}
	return nil
}

func (w *ComplexUnionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ComplexUnionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ComplexUnionBuilder) Build() *ComplexUnion {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ComplexUnionReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ExampleEnum uint16

const (
	EXAMPLE_ENUM_OPTION_A ExampleEnum = 0
	EXAMPLE_ENUM_OPTION_B ExampleEnum = 1
	EXAMPLE_ENUM_OPTION_C ExampleEnum = 2
)

func (n ExampleEnum) String() string {
	switch n {
	case EXAMPLE_ENUM_OPTION_A:
		return "EXAMPLE_ENUM_OPTION_A"
	case EXAMPLE_ENUM_OPTION_B:
		return "EXAMPLE_ENUM_OPTION_B"
	case EXAMPLE_ENUM_OPTION_C:
		return "EXAMPLE_ENUM_OPTION_C"
	}
	return "UNKNOWN"
}

