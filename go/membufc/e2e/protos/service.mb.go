// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package types

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// service StateStorage

type StateStorage interface {
	WriteKey(input *WriteKeyInput) (*WriteKeyOutput, error)
	ReadKey(input *ReadKeyInput) (*ReadKeyOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message WriteKeyInput

// reader

type WriteKeyInput struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _WriteKeyInput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeUint32,}
var _WriteKeyInput_Unions = [][]membuffers.FieldType{}

func WriteKeyInputReader(buf []byte) *WriteKeyInput {
	x := &WriteKeyInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _WriteKeyInput_Scheme, _WriteKeyInput_Unions)
	return x
}

func (x *WriteKeyInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *WriteKeyInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *WriteKeyInput) Key() string {
	return x._message.GetString(0)
}

func (x *WriteKeyInput) RawKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *WriteKeyInput) MutateKey(v string) error {
	return x._message.SetString(0, v)
}

func (x *WriteKeyInput) Value() uint32 {
	return x._message.GetUint32(1)
}

func (x *WriteKeyInput) RawValue() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *WriteKeyInput) MutateValue(v uint32) error {
	return x._message.SetUint32(1, v)
}

// builder

type WriteKeyInputBuilder struct {
	Key string
	Value uint32

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *WriteKeyInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Key)
	w._builder.WriteUint32(buf, w.Value)
	return nil
}

func (w *WriteKeyInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *WriteKeyInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *WriteKeyInputBuilder) Build() *WriteKeyInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return WriteKeyInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message WriteKeyOutput

// reader

type WriteKeyOutput struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _WriteKeyOutput_Scheme = []membuffers.FieldType{}
var _WriteKeyOutput_Unions = [][]membuffers.FieldType{}

func WriteKeyOutputReader(buf []byte) *WriteKeyOutput {
	x := &WriteKeyOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _WriteKeyOutput_Scheme, _WriteKeyOutput_Unions)
	return x
}

func (x *WriteKeyOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *WriteKeyOutput) Raw() []byte {
	return x._message.RawBuffer()
}

// builder

type WriteKeyOutputBuilder struct {

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *WriteKeyOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	return nil
}

func (w *WriteKeyOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *WriteKeyOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *WriteKeyOutputBuilder) Build() *WriteKeyOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return WriteKeyOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeyInput

// reader

type ReadKeyInput struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ReadKeyInput_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _ReadKeyInput_Unions = [][]membuffers.FieldType{}

func ReadKeyInputReader(buf []byte) *ReadKeyInput {
	x := &ReadKeyInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ReadKeyInput_Scheme, _ReadKeyInput_Unions)
	return x
}

func (x *ReadKeyInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *ReadKeyInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ReadKeyInput) Key() string {
	return x._message.GetString(0)
}

func (x *ReadKeyInput) RawKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ReadKeyInput) MutateKey(v string) error {
	return x._message.SetString(0, v)
}

// builder

type ReadKeyInputBuilder struct {
	Key string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ReadKeyInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Key)
	return nil
}

func (w *ReadKeyInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ReadKeyInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ReadKeyInputBuilder) Build() *ReadKeyInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReadKeyInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ReadKeyOutput

// reader

type ReadKeyOutput struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ReadKeyOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,}
var _ReadKeyOutput_Unions = [][]membuffers.FieldType{}

func ReadKeyOutputReader(buf []byte) *ReadKeyOutput {
	x := &ReadKeyOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ReadKeyOutput_Scheme, _ReadKeyOutput_Unions)
	return x
}

func (x *ReadKeyOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *ReadKeyOutput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ReadKeyOutput) Value() uint32 {
	return x._message.GetUint32(0)
}

func (x *ReadKeyOutput) RawValue() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ReadKeyOutput) MutateValue(v uint32) error {
	return x._message.SetUint32(0, v)
}

// builder

type ReadKeyOutputBuilder struct {
	Value uint32

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *ReadKeyOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, w.Value)
	return nil
}

func (w *ReadKeyOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ReadKeyOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ReadKeyOutputBuilder) Build() *ReadKeyOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReadKeyOutputReader(buf)
}

