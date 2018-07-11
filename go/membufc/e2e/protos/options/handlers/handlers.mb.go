// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package handlers

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
// service ServicesINeedFromOthersHandler

type ServicesINeedFromOthersHandler interface {
	SomeMethodINeedFromOthers(input *SomeMessage) (*SomeMessage, error)
}

/////////////////////////////////////////////////////////////////////////////
// service ServicesIProvideToOthersHandler

type ServicesIProvideToOthersHandler interface {
	SomeMethodIProvideToOthers(input *SomeMessage) (*SomeMessage, error)
}

/////////////////////////////////////////////////////////////////////////////
// message SomeMessage

// reader

type SomeMessage struct {
	// Str string

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *SomeMessage) String() string {
	return fmt.Sprintf("{Str:%s,}", x.StringStr())
}

var _SomeMessage_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _SomeMessage_Unions = [][]membuffers.FieldType{}

func SomeMessageReader(buf []byte) *SomeMessage {
	x := &SomeMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SomeMessage_Scheme, _SomeMessage_Unions)
	return x
}

func (x *SomeMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *SomeMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SomeMessage) Str() string {
	return x._message.GetString(0)
}

func (x *SomeMessage) RawStr() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SomeMessage) MutateStr(v string) error {
	return x._message.SetString(0, v)
}

func (x *SomeMessage) StringStr() string {
	return fmt.Sprintf(x.Str())
}

// builder

type SomeMessageBuilder struct {
	Str string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *SomeMessageBuilder) Write(buf []byte) (err error) {
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

func (w *SomeMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SomeMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SomeMessageBuilder) Build() *SomeMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SomeMessageReader(buf)
}

