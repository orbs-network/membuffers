// AUTO GENERATED FILE (by membufc proto compiler)
package dep1

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message DependencyMessage

// reader

type DependencyMessage struct {
	message membuffers.Message
}

var m_DependencyMessage_Scheme = []membuffers.FieldType{membuffers.TypeUint32,}
var m_DependencyMessage_Unions = [][]membuffers.FieldType{}

func DependencyMessageReader(buf []byte) *DependencyMessage {
	x := &DependencyMessage{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_DependencyMessage_Scheme, m_DependencyMessage_Unions)
	return x
}

func (x *DependencyMessage) IsValid() bool {
	return x.message.IsValid()
}

func (x *DependencyMessage) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *DependencyMessage) Field() uint32 {
	return x.message.GetUint32(0)
}

func (x *DependencyMessage) RawField() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *DependencyMessage) MutateField(v uint32) error {
	return x.message.SetUint32(0, v)
}

// builder

type DependencyMessageBuilder struct {
	builder membuffers.Builder
	Field uint32
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
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.Field)
	return nil
}

func (w *DependencyMessageBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *DependencyMessageBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

