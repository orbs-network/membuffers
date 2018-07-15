// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package dep2

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/dep1"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/dep1/dep11"
)

/////////////////////////////////////////////////////////////////////////////
// message Dependent

// reader

type Dependent struct {
	// A dep1.DependencyMessage
	// B dep11.DependencyEnum
	// C SamePackageDependencyMessage

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Dependent) String() string {
	return fmt.Sprintf("{A:%s,B:%s,C:%s,}", x.StringA(), x.StringB(), x.StringC())
}

var _Dependent_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeMessage,}
var _Dependent_Unions = [][]membuffers.FieldType{}

func DependentReader(buf []byte) *Dependent {
	x := &Dependent{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Dependent_Scheme, _Dependent_Unions)
	return x
}

func (x *Dependent) IsValid() bool {
	return x._message.IsValid()
}

func (x *Dependent) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Dependent) A() *dep1.DependencyMessage {
	b, s := x._message.GetMessage(0)
	return dep1.DependencyMessageReader(b[:s])
}

func (x *Dependent) RawA() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Dependent) StringA() string {
	return x.A().String()
}

func (x *Dependent) B() dep11.DependencyEnum {
	return dep11.DependencyEnum(x._message.GetUint16(1))
}

func (x *Dependent) RawB() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Dependent) MutateB(v dep11.DependencyEnum) error {
	return x._message.SetUint16(1, uint16(v))
}

func (x *Dependent) StringB() string {
	return x.B().String()
}

func (x *Dependent) C() *SamePackageDependencyMessage {
	b, s := x._message.GetMessage(2)
	return SamePackageDependencyMessageReader(b[:s])
}

func (x *Dependent) RawC() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Dependent) StringC() string {
	return x.C().String()
}

// builder

type DependentBuilder struct {
	A *dep1.DependencyMessageBuilder
	B dep11.DependencyEnum
	C *SamePackageDependencyMessageBuilder

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *DependentBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.A)
	if err != nil {
		return
	}
	w._builder.WriteUint16(buf, uint16(w.B))
	err = w._builder.WriteMessage(buf, w.C)
	if err != nil {
		return
	}
	return nil
}

func (w *DependentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *DependentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *DependentBuilder) Build() *Dependent {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return DependentReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

