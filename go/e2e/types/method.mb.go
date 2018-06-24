package types

import "github.com/orbs-network/membuffers/go"

/*
message Method {
	string name = 1;
	repeated MethodCallArgument arg = 2;
}
*/

// reader

type Method struct {
	message membuffers.Message
}

var m_Method_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeMessageArray}
var m_Method_Unions = [][]membuffers.FieldType{{}}

func MethodReader(buf []byte) *Method {
	x := &Method{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Method_Scheme, m_Method_Unions)
	return x
}

func (x *Method) IsValid() bool {
	return x.message.IsValid()
}

func (x *Method) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *Method) Name() string {
	return x.message.GetString(0)
}

func (x *Method) RawName() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Method) MutateName(v string) error {
	return x.message.SetString(0, v)
}

func (x *Method) ArgIterator() *MethodArgIterator {
	return &MethodArgIterator{iterator: x.message.GetMessageArrayIterator(1)}
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
	return x.message.RawBufferForField(1, 0)
}

// builder

type MethodBuilder struct {
	builder membuffers.Builder
	Name    string
	Arg     []*MethodCallArgumentBuilder
}

func (w *MethodBuilder) arg() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Arg))
	for i, v := range w.Arg {
		res[i] = v
	}
	return res
}

func (w *MethodBuilder) Write(buf []byte) {
	if w == nil {
		return
	}
	w.builder.Reset()
	w.builder.WriteString(buf, w.Name)
	w.builder.WriteMessageArray(buf, w.arg())
}

func (w *MethodBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MethodBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/*
message MethodCallArgument {
	oneof type {
		uint32 num = 1;
		string str = 2;
		bytes data = 3;
	}
}
*/

// reader

type MethodCallArgument struct {
	message membuffers.Message
}

var m_MethodCallArgument_Scheme = []membuffers.FieldType{membuffers.TypeUnion}
var m_MethodCallArgument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32,membuffers.TypeString,membuffers.TypeBytes}}

func MethodCallArgumentReader(buf []byte) *MethodCallArgument {
	x := &MethodCallArgument{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_MethodCallArgument_Scheme, m_MethodCallArgument_Unions)
	return x
}

func (x *MethodCallArgument) IsValid() bool {
	return x.message.IsValid()
}

func (x *MethodCallArgument) Raw() []byte {
	return x.message.RawBuffer()
}

type MethodCallArgumentType uint16

const (
	MethodCallArgumentTypeNum  MethodCallArgumentType = 0
	MethodCallArgumentTypeStr  MethodCallArgumentType = 1
	MethodCallArgumentTypeData MethodCallArgumentType = 2
)

func (x *MethodCallArgument) Type() MethodCallArgumentType {
	return MethodCallArgumentType(x.message.GetUint16(0))
}

func (x *MethodCallArgument) IsTypeNum() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *MethodCallArgument) TypeNum() uint32 {
	_, off := x.message.IsUnionIndex(0, 0, 0)
	return x.message.GetUint32InOffset(off)
}

func (x *MethodCallArgument) MutateTypeNum(v uint32) error {
	is, off := x.message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.SetUint32InOffset(off, v)
	return nil
}

func (x *MethodCallArgument) IsTypeStr() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *MethodCallArgument) TypeStr() string {
	_, off := x.message.IsUnionIndex(0, 0, 1)
	return x.message.GetStringInOffset(off)
}

func (x *MethodCallArgument) MutateTypeStr(v string) error {
	is, off := x.message.IsUnionIndex(0, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.GetStringInOffset(off)
	return nil
}

func (x *MethodCallArgument) IsTypeData() bool {
	is, _ := x.message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *MethodCallArgument) TypeData() []byte {
	_, off := x.message.IsUnionIndex(0, 0, 2)
	return x.message.GetBytesInOffset(off)
}

func (x *MethodCallArgument) MutateTypeData(v []byte) error {
	is, off := x.message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x.message.GetBytesInOffset(off)
	return nil
}

func (x *MethodCallArgument) RawType() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type MethodCallArgumentBuilder struct {
	builder membuffers.Builder
	Num     uint32
	Str     string
	Data    []byte
	Type    MethodCallArgumentType
}

func (w *MethodCallArgumentBuilder) Write(buf []byte) {
	if w == nil {
		return
	}
	w.builder.Reset()
	w.builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case MethodCallArgumentTypeNum:
		w.builder.WriteUint32(buf, w.Num)
	case MethodCallArgumentTypeStr:
		w.builder.WriteString(buf, w.Str)
	case MethodCallArgumentTypeData:
		w.builder.WriteBytes(buf, w.Data)
	}
}

func (w *MethodCallArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *MethodCallArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}