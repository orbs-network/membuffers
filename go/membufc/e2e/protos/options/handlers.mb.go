// AUTO GENERATED FILE (by membufc proto compiler)
package options

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ExampleService

type ExampleService interface {
	handlers.ServicesINeedFromOthersHandler
	ExampleMethod(*ExampleMethodInput) (*ExampleMethodOutput, error)
	RegisterServicesIProvideToOthersHandler(handlers.ServicesIProvideToOthersHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message ExampleMethodInput

// reader

type ExampleMethodInput struct {
	message membuffers.Message
}

var m_ExampleMethodInput_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var m_ExampleMethodInput_Unions = [][]membuffers.FieldType{}

func ExampleMethodInputReader(buf []byte) *ExampleMethodInput {
	x := &ExampleMethodInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ExampleMethodInput_Scheme, m_ExampleMethodInput_Unions)
	return x
}

func (x *ExampleMethodInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ExampleMethodInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *ExampleMethodInput) Arg() string {
	return x.message.GetString(0)
}

func (x *ExampleMethodInput) RawArg() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *ExampleMethodInput) MutateArg(v string) error {
	return x.message.SetString(0, v)
}

// builder

type ExampleMethodInputBuilder struct {
	builder membuffers.Builder
	Arg string
}

func (w *ExampleMethodInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteString(buf, w.Arg)
	return nil
}

func (w *ExampleMethodInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ExampleMethodInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ExampleMethodInputBuilder) Build() *ExampleMethodInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ExampleMethodInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message ExampleMethodOutput

// reader

type ExampleMethodOutput struct {
	message membuffers.Message
}

var m_ExampleMethodOutput_Scheme = []membuffers.FieldType{}
var m_ExampleMethodOutput_Unions = [][]membuffers.FieldType{}

func ExampleMethodOutputReader(buf []byte) *ExampleMethodOutput {
	x := &ExampleMethodOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_ExampleMethodOutput_Scheme, m_ExampleMethodOutput_Unions)
	return x
}

func (x *ExampleMethodOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *ExampleMethodOutput) Raw() []byte {
	return x.message.RawBuffer()
}

// builder

type ExampleMethodOutputBuilder struct {
	builder membuffers.Builder
}

func (w *ExampleMethodOutputBuilder) Write(buf []byte) (err error) {
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

func (w *ExampleMethodOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *ExampleMethodOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *ExampleMethodOutputBuilder) Build() *ExampleMethodOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ExampleMethodOutputReader(buf)
}

