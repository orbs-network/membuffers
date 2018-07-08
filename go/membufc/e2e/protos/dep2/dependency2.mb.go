// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package dep2

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message SamePackageDependencyMessage

// reader

type SamePackageDependencyMessage struct {
	message membuffers.Message
}

var _SamePackageDependencyMessage_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _SamePackageDependencyMessage_Unions = [][]membuffers.FieldType{}

func SamePackageDependencyMessageReader(buf []byte) *SamePackageDependencyMessage {
	x := &SamePackageDependencyMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _SamePackageDependencyMessage_Scheme, _SamePackageDependencyMessage_Unions)
	return x
}

func (x *SamePackageDependencyMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *SamePackageDependencyMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SamePackageDependencyMessage) Field() string {
	return x.message.GetString(0)
}

func (x *SamePackageDependencyMessage) RawField() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SamePackageDependencyMessage) MutateField(v string) error {
	return x.message.SetString(0, v)
}

// builder

type SamePackageDependencyMessageBuilder struct {
	builder membuffers.Builder
	Field string
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Field)
	return nil
}

func (w *SamePackageDependencyMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SamePackageDependencyMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SamePackageDependencyMessageBuilder) Build() *SamePackageDependencyMessage {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SamePackageDependencyMessageReader(buf)
}

