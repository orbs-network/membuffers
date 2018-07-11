// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package types

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
// message ExampleContainer (non serializable)

type ExampleContainer struct {
	Message1 *MessageInContainer
	Container1 *NestedContainer
	Containers2 []*NestedContainer
}

/////////////////////////////////////////////////////////////////////////////
// message MessageInContainer

// reader

type MessageInContainer struct {
	// Field string

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *MessageInContainer) String() string {
	return fmt.Sprintf("{Field:%s,}", x.StringField())
}

var _MessageInContainer_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _MessageInContainer_Unions = [][]membuffers.FieldType{}

func MessageInContainerReader(buf []byte) *MessageInContainer {
	x := &MessageInContainer{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _MessageInContainer_Scheme, _MessageInContainer_Unions)
	return x
}

func (x *MessageInContainer) IsValid() bool {
	return x._message.IsValid()
}

func (x *MessageInContainer) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *MessageInContainer) Field() string {
	return x._message.GetString(0)
}

func (x *MessageInContainer) RawField() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *MessageInContainer) MutateField(v string) error {
	return x._message.SetString(0, v)
}

func (x *MessageInContainer) StringField() string {
	return fmt.Sprintf(x.Field())
}

// builder

type MessageInContainerBuilder struct {
	Field string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *MessageInContainerBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Field)
	return nil
}

func (w *MessageInContainerBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *MessageInContainerBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *MessageInContainerBuilder) Build() *MessageInContainer {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return MessageInContainerReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message NestedContainer (non serializable)

type NestedContainer struct {
	Name string
}

