// AUTO GENERATED FILE (by membufc proto compiler)
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
	message membuffers.Message
}

var m_WriteKeyInput_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeUint32,}
var m_WriteKeyInput_Unions = [][]membuffers.FieldType{}

func WriteKeyInputReader(buf []byte) *WriteKeyInput {
	x := &WriteKeyInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_WriteKeyInput_Scheme, m_WriteKeyInput_Unions)
	return x
}

func (x *WriteKeyInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *WriteKeyInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *WriteKeyInput) Key() string {
	return x.message.GetString(0)
}

func (x *WriteKeyInput) RawKey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *WriteKeyInput) MutateKey(v string) error {
	return x.message.SetString(0, v)
}

func (x *WriteKeyInput) Value() uint32 {
	return x.message.GetUint32(1)
}

func (x *WriteKeyInput) RawValue() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *WriteKeyInput) MutateValue(v uint32) error {
	return x.message.SetUint32(1, v)
}

// builder

type WriteKeyInputBuilder struct {
	builder membuffers.Builder
	Key string
	Value uint32
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Key)
	w.builder.WriteUint32(buf, w.Value)
	return nil
}

func (w *WriteKeyInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *WriteKeyInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_WriteKeyOutput_Scheme = []membuffers.FieldType{}
var m_WriteKeyOutput_Unions = [][]membuffers.FieldType{}

func WriteKeyOutputReader(buf []byte) *WriteKeyOutput {
	x := &WriteKeyOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_WriteKeyOutput_Scheme, m_WriteKeyOutput_Unions)
	return x
}

func (x *WriteKeyOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *WriteKeyOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type WriteKeyOutputBuilder struct {
	builder membuffers.Builder
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
	w.builder.Reset()
	return nil
}

func (w *WriteKeyOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *WriteKeyOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_ReadKeyInput_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var m_ReadKeyInput_Unions = [][]membuffers.FieldType{}

func ReadKeyInputReader(buf []byte) *ReadKeyInput {
	x := &ReadKeyInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReadKeyInput_Scheme, m_ReadKeyInput_Unions)
	return x
}

func (x *ReadKeyInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReadKeyInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ReadKeyInput) Key() string {
	return x.message.GetString(0)
}

func (x *ReadKeyInput) RawKey() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ReadKeyInput) MutateKey(v string) error {
	return x.message.SetString(0, v)
}

// builder

type ReadKeyInputBuilder struct {
	builder membuffers.Builder
	Key string
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Key)
	return nil
}

func (w *ReadKeyInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReadKeyInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
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
	message membuffers.Message
}

var m_ReadKeyOutput_Scheme = []membuffers.FieldType{membuffers.TypeUint32,}
var m_ReadKeyOutput_Unions = [][]membuffers.FieldType{}

func ReadKeyOutputReader(buf []byte) *ReadKeyOutput {
	x := &ReadKeyOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ReadKeyOutput_Scheme, m_ReadKeyOutput_Unions)
	return x
}

func (x *ReadKeyOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ReadKeyOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ReadKeyOutput) Value() uint32 {
	return x.message.GetUint32(0)
}

func (x *ReadKeyOutput) RawValue() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ReadKeyOutput) MutateValue(v uint32) error {
	return x.message.SetUint32(0, v)
}

// builder

type ReadKeyOutputBuilder struct {
	builder membuffers.Builder
	Value uint32
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
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.Value)
	return nil
}

func (w *ReadKeyOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ReadKeyOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ReadKeyOutputBuilder) Build() *ReadKeyOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ReadKeyOutputReader(buf)
}

