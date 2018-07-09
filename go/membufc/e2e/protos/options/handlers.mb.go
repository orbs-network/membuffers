// AUTO GENERATED FILE (by membufc proto compiler v0.0.12)
package options

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/options/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service ExampleService

type ExampleService interface {
	handlers.ServicesINeedFromOthersHandler
	ExampleMethod(input *ExampleMethodInput) (*ExampleMethodOutput, error)
	RegisterServicesIProvideToOthersHandler(handler handlers.ServicesIProvideToOthersHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message ExampleMethodInput

// reader

type ExampleMethodInput struct {
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ExampleMethodInput_Scheme = []membuffers.FieldType{membuffers.TypeString,}
var _ExampleMethodInput_Unions = [][]membuffers.FieldType{}

func ExampleMethodInputReader(buf []byte) *ExampleMethodInput {
	x := &ExampleMethodInput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ExampleMethodInput_Scheme, _ExampleMethodInput_Unions)
	return x
}

func (x *ExampleMethodInput) IsValid() bool {
	return x._message.IsValid()
}

func (x *ExampleMethodInput) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ExampleMethodInput) Arg() string {
	return x._message.GetString(0)
}

func (x *ExampleMethodInput) RawArg() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ExampleMethodInput) MutateArg(v string) error {
	return x._message.SetString(0, v)
}

// builder

type ExampleMethodInputBuilder struct {
	Arg string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	w._builder.WriteString(buf, w.Arg)
	return nil
}

func (w *ExampleMethodInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ExampleMethodInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
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
	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _ExampleMethodOutput_Scheme = []membuffers.FieldType{}
var _ExampleMethodOutput_Unions = [][]membuffers.FieldType{}

func ExampleMethodOutputReader(buf []byte) *ExampleMethodOutput {
	x := &ExampleMethodOutput{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ExampleMethodOutput_Scheme, _ExampleMethodOutput_Unions)
	return x
}

func (x *ExampleMethodOutput) IsValid() bool {
	return x._message.IsValid()
}

func (x *ExampleMethodOutput) Raw() []byte {
	return x._message.RawBuffer()
}

// builder

type ExampleMethodOutputBuilder struct {

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
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
	w._builder.Reset()
	return nil
}

func (w *ExampleMethodOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ExampleMethodOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ExampleMethodOutputBuilder) Build() *ExampleMethodOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ExampleMethodOutputReader(buf)
}

