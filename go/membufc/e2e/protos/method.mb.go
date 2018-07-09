// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package types

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Method

// reader

type Method struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _Method_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeMessageArray,}
var _Method_Unions = [][]membuffers.FieldType{}

func MethodReader(buf []byte) *Method {
	x := &Method{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Method_Scheme, _Method_Unions)
	return x
}

func (x *Method) IsValid() bool {
	return x._message.IsValid()
}

func (x *Method) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Method) Name() string {
	return x._message.GetString(0)
}

func (x *Method) RawName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Method) MutateName(v string) error {
	return x._message.SetString(0, v)
}

func (x *Method) ArgIterator() *MethodArgIterator {
	return &MethodArgIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type MethodArgIterator struct {
	iterator *membuffers.Iterator
}

func (i *MethodArgIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *MethodArgIterator) NextArg() *MethodCallArgument {
	b, s := i.iterator.NextMessage()
	return MethodCallArgumentReader(b[:s])
}

func (x *Method) RawArgArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type MethodBuilder struct {
	Name string
	Arg []*MethodCallArgumentBuilder

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MethodBuilder) arrayOfArg() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Arg))
	for i, v := range w.Arg {
		res[i] = v
	}
	return res
}

func (w *MethodBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Name)
	err = w._builder.WriteMessageArray(buf, w.arrayOfArg())
	if err != nil {
		return
	}
	return nil
}

func (w *MethodBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MethodBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MethodBuilder) Build() *Method {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MethodReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message MethodCallArgument

// reader

type MethodCallArgument struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _MethodCallArgument_Scheme = []membuffers.FieldType{membuffers.TypeUnion,}
var _MethodCallArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeString,membuffers.TypeBytes,}}

func MethodCallArgumentReader(buf []byte) *MethodCallArgument {
	x := &MethodCallArgument{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MethodCallArgument_Scheme, _MethodCallArgument_Unions)
	return x
}

func (x *MethodCallArgument) IsValid() bool {
	return x._message.IsValid()
}

func (x *MethodCallArgument) Raw() []byte {
	return x._message.RawBuffer()
}

type MethodCallArgumentType uint16

const (
	MethodCallArgumentTypeNum MethodCallArgumentType = 0
	MethodCallArgumentTypeStr MethodCallArgumentType = 1
	MethodCallArgumentTypeData MethodCallArgumentType = 2
)

func (x *MethodCallArgument) Type() MethodCallArgumentType {
	return MethodCallArgumentType(x._message.GetUint16(0))
}

func (x *MethodCallArgument) IsTypeNum() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *MethodCallArgument) TypeNum() uint32 {
	_, off := x._message.IsUnionIndex(0, 0, 0)
	return x._message.GetUint32InOffset(off)
}

func (x *MethodCallArgument) MutateTypeNum(v uint32) error {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeStr() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *MethodCallArgument) TypeStr() string {
	_, off := x._message.IsUnionIndex(0, 0, 1)
	return x._message.GetStringInOffset(off)
}

func (x *MethodCallArgument) MutateTypeStr(v string) error {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetStringInOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeData() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *MethodCallArgument) TypeData() []byte {
	_, off := x._message.IsUnionIndex(0, 0, 2)
	return x._message.GetBytesInOffset(off)
}

func (x *MethodCallArgument) MutateTypeData(v []byte) error {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, v)
	return nil
}

func (x *MethodCallArgument) RawType() []byte {
	return x._message.RawBufferForField(0, 0)
}

// builder

type MethodCallArgumentBuilder struct {
	Type MethodCallArgumentType
	Num uint32
	Str string
	Data []byte

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MethodCallArgumentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MethodCallArgumentTypeNum:
		w._builder.WriteUint32(buf, w.Num)
	case MethodCallArgumentTypeStr:
		w._builder.WriteString(buf, w.Str)
	case MethodCallArgumentTypeData:
		w._builder.WriteBytes(buf, w.Data)
	}
	return nil
}

func (w *MethodCallArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MethodCallArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MethodCallArgumentBuilder) Build() *MethodCallArgument {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MethodCallArgumentReader(buf)
}

