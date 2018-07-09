// AUTO GENERATED FILE (by membufc proto compiler v0.0.13)
package dep2

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message SamePackageDependencyMessage

// reader

type SamePackageDependencyMessage struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _SamePackageDependencyMessage_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _SamePackageDependencyMessage_Unions = [][]membuffers.FieldType{}

func SamePackageDependencyMessageReader(buf []byte) *SamePackageDependencyMessage {
	x := &SamePackageDependencyMessage{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _SamePackageDependencyMessage_Scheme, _SamePackageDependencyMessage_Unions)
	return x
}

func (x *SamePackageDependencyMessage) IsValid() bool {
	return x._message.IsValid()
}

func (x *SamePackageDependencyMessage) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *SamePackageDependencyMessage) Field() string {
	return x._message.GetString(0)
}

func (x *SamePackageDependencyMessage) RawField() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *SamePackageDependencyMessage) MutateField(v string) error {
	return x._message.SetString(0, v)
}

// builder

type SamePackageDependencyMessageBuilder struct {
	Field string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *SamePackageDependencyMessageBuilder) Write(buf []byte) (err error) {
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

func (w *SamePackageDependencyMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *SamePackageDependencyMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *SamePackageDependencyMessageBuilder) Build() *SamePackageDependencyMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SamePackageDependencyMessageReader(buf)
}

