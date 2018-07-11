// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package dep1

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
)

/////////////////////////////////////////////////////////////////////////////
// message DependencyMessage

// reader

type DependencyMessage struct {
	// Field uint32

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

func (x *DependencyMessage) String() string {
	return fmt.Sprintf("{Field:%s,}", x.StringField())
}

var _DependencyMessage_Scheme = []membuffers.FieldType{membuffers.TypeUint32,}
var _DependencyMessage_Unions = [][]membuffers.FieldType{}

func DependencyMessageReader(buf []byte) *DependencyMessage {
	x := &DependencyMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _DependencyMessage_Scheme, _DependencyMessage_Unions)
	return x
}

func (x *DependencyMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *DependencyMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *DependencyMessage) Field() uint32 {
	return x._message.GetUint32(0)
}

func (x *DependencyMessage) RawField() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *DependencyMessage) MutateField(v uint32) error {
	return x._message.SetUint32(0, v)
}

func (x *DependencyMessage) StringField() string {
	return fmt.Sprintf("%x", x.Field())
}

// builder

type DependencyMessageBuilder struct {
	Field uint32

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *DependencyMessageBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, w.Field)
	return nil
}

func (w *DependencyMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *DependencyMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *DependencyMessageBuilder) Build() *DependencyMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return DependencyMessageReader(buf)
}

