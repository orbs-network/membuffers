package main

import (
	"github.com/tallstoat/pbparser"
	"io"
	"fmt"
	"sort"
)

func convertFieldNameToGoCase(fieldName string) string {
	return ToCamel(fieldName)
}

func compileProtoFile(w io.Writer, file pbparser.ProtoFile) {
	addHeader(w, file.PackageName)
	for _, m := range file.Messages {
		addMessage(w, m)
	}
}

func fixFields(m pbparser.MessageElement) {
	sort.Slice(m.Fields, func(i, j int) bool {
		return m.Fields[i].Tag < m.Fields[j].Tag
	})
	for i, _ := range m.Fields {
		m.Fields[i].Tag = i
	}
}

func addHeader(w io.Writer, packageName string) {
	fmt.Fprintf(w, `// AUTO GENERATED FILE (by membufc proto compiler)
package %s
`, packageName)
	fmt.Fprintf(w, `
import (
	"github.com/orbs-network/membuffers/go"
)
`)
}

func addMessage(w io.Writer, m pbparser.MessageElement) {
	fixFields(m)
	fmt.Fprintf(w, `
/////////////////////////////////////////////////////////////////////////////
// message %s

// reader

type %s struct {
	message membuffers.Message
}

`, m.Name, m.Name)
	addMessageScheme(w, m.Name, m.Fields)
	addMessageUnions(w, m.Name, m.Fields)
	addMessageReaderStart(w, m.Name, m.Fields)
	for _, field := range m.Fields {
		addMessageReaderField(w, m.Name, field)
	}
}

func addMessageScheme(w io.Writer, name string, fields []pbparser.FieldElement) {
	fmt.Fprintf(w, "var m_%s_Scheme = []membuffers.FieldType{", name)
	for _, field := range fields {
		if field.Label == "" || field.Label == "optional" || field.Label == "required" {
			if field.Type.Category() == pbparser.ScalarDataTypeCategory {
				types := map[string]string{
					"bytes": "TypeBytes",
					"string": "TypeString",
					"uint8": "TypeUint8",
					"uint16": "TypeUint16",
					"uint32": "TypeUint32",
					"uint64": "TypeUint64",
				}
				fmt.Fprintf(w, "membuffers.%s,", types[field.Type.Name()])
			}
			if field.Type.Category() == pbparser.NamedDataTypeCategory {
				fmt.Fprintf(w, "membuffers.TypeMessage,")
			}
		}
		if field.Label == "repeated" {
			if field.Type.Category() == pbparser.ScalarDataTypeCategory {
				types := map[string]string{
					"bytes": "TypeBytesArray",
					"string": "TypeStringArray",
					"uint8": "TypeUint8Array",
					"uint16": "TypeUint16Array",
					"uint32": "TypeUint32Array",
					"uint64": "TypeUint64Array",
				}
				fmt.Fprintf(w, "membuffers.%s,", types[field.Type.Name()])
			}
			if field.Type.Category() == pbparser.NamedDataTypeCategory {
				fmt.Fprintf(w, "membuffers.TypeMessageArray,")
			}
		}
	}
	fmt.Fprintf(w, "}\n")
}

func addMessageUnions(w io.Writer, name string, fields []pbparser.FieldElement) {
	fmt.Fprintf(w, "var m_%s_Unions = [][]membuffers.FieldType{{}}\n", name)
}

func addMessageReaderStart(w io.Writer, name string, fields []pbparser.FieldElement) {
	fmt.Fprintf(w, `
func %sReader(buf []byte) *%s {
	x := &%s{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_%s_Scheme, m_%s_Unions)
	return x
}

func (x *%s) IsValid() bool {
	return x.message.IsValid()
}

func (x *%s) Raw() []byte {
	return x.message.RawBuffer()
}
`, name, name, name, name, name, name, name)
}

func addMessageReaderField(w io.Writer, name string, field pbparser.FieldElement) {
	if field.Label == "" || field.Label == "optional" || field.Label == "required" {
		if field.Type.Category() == pbparser.ScalarDataTypeCategory {
			goTypes := map[string]string{
				"bytes":  "[]byte",
				"string": "string",
				"uint8":  "uint8",
				"uint16": "uint16",
				"uint32": "uint32",
				"uint64": "uint64",
			}
			accessor := map[string]string{
				"bytes":  "Bytes",
				"string": "String",
				"uint8":  "Uint8",
				"uint16": "Uint16",
				"uint32": "Uint32",
				"uint64": "Uint64",
			}
			fmt.Fprintf(w, `
func (x *%s) %s() %s {
	return x.message.Get%s(%d)
}

func (x *%s) Raw%s() []byte {
	return x.message.RawBufferForField(%d, 0)
}

func (x *%s) Mutate%s(v %s) error {
	return x.message.Set%s(%d, v)
}
`, name, convertFieldNameToGoCase(field.Name), goTypes[field.Type.Name()], accessor[field.Type.Name()], field.Tag, name, convertFieldNameToGoCase(field.Name), field.Tag, name, convertFieldNameToGoCase(field.Name), goTypes[field.Type.Name()], accessor[field.Type.Name()], field.Tag)
		}
		if field.Type.Category() == pbparser.NamedDataTypeCategory {
			fmt.Fprintf(w, `
func (x *%s) %s() *%s {
	b, s := x.message.GetMessage(%d)
	return %sReader(b[:s])
}

func (x *%s) Raw%s() []byte {
	return x.message.RawBufferForField(%d, 0)
}
`, name, convertFieldNameToGoCase(field.Name), field.Type.Name(), field.Tag, field.Type.Name(), name, convertFieldNameToGoCase(field.Name), field.Tag)
		}
	}
}