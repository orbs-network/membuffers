package main

import (
	"github.com/tallstoat/pbparser"
	"io"
	"sort"
	"text/template"
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
	t := template.Must(template.ParseFiles("templates/go/MessageFileHeader.template"))
	t.Execute(w, struct {
		PackageName string
	}{
		PackageName: packageName,
	})
}

func addMessage(w io.Writer, m pbparser.MessageElement) {
	fixFields(m)
	t := template.Must(template.ParseFiles("templates/go/MessageHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: m.Name,
	})
	addMessageScheme(w, m.Name, m.Fields)
	addMessageUnions(w, m.Name, m.Fields)
	addMessageReaderHeader(w, m.Name, m.Fields)
	for _, field := range m.Fields {
		addMessageReaderField(w, m.Name, field)
	}
	addMessageBuilder(w, m.Name, m.Fields)
}

func addMessageScheme(w io.Writer, messageName string, fields []pbparser.FieldElement) {
	fieldTypes := []string{}
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
				fieldTypes = append(fieldTypes, types[field.Type.Name()])
			}
			if field.Type.Category() == pbparser.NamedDataTypeCategory {
				fieldTypes = append(fieldTypes, "TypeMessage")
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
				fieldTypes = append(fieldTypes, types[field.Type.Name()])
			}
			if field.Type.Category() == pbparser.NamedDataTypeCategory {
				fieldTypes = append(fieldTypes, "TypeMessageArray")
			}
		}
	}
	t := template.Must(template.ParseFiles("templates/go/MessageScheme.template"))
	t.Execute(w, struct {
		MessageName string
		FieldTypes []string
	}{
		MessageName: messageName,
		FieldTypes:  fieldTypes,
	})
}

func addMessageUnions(w io.Writer, messageName string, fields []pbparser.FieldElement) {
	t := template.Must(template.ParseFiles("templates/go/MessageUnions.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: messageName,
	})
}

func addMessageReaderHeader(w io.Writer, messageName string, fields []pbparser.FieldElement) {
	t := template.Must(template.ParseFiles("templates/go/MessageReaderHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: messageName,
	})
}

func addMessageReaderField(w io.Writer, messageName string, field pbparser.FieldElement) {
	messageField := getMessageField(messageName, field)
	if !messageField.IsMessage && !messageField.IsArray {
		t := template.Must(template.ParseFiles("templates/go/MessageReaderMutableField.template"))
		t.Execute(w, struct {
			MessageName string
			MessageField MessageField
		}{
			MessageName:  messageName,
			MessageField: messageField,
		})
	}
	if messageField.IsMessage && !messageField.IsArray {
		t := template.Must(template.ParseFiles("templates/go/MessageReaderMessageField.template"))
		t.Execute(w, struct {
			MessageName string
			MessageField MessageField
		}{
			MessageName:  messageName,
			MessageField: messageField,
		})
	}
	if !messageField.IsMessage && messageField.IsArray {
		t := template.Must(template.ParseFiles("templates/go/MessageReaderMutableArrayField.template"))
		t.Execute(w, struct {
			MessageName string
			MessageField MessageField
		}{
			MessageName:  messageName,
			MessageField: messageField,
		})
	}
	if messageField.IsMessage && messageField.IsArray {
		t := template.Must(template.ParseFiles("templates/go/MessageReaderMessageArrayField.template"))
		t.Execute(w, struct {
			MessageName string
			MessageField MessageField
		}{
			MessageName:  messageName,
			MessageField: messageField,
		})
	}
}

func addMessageBuilder(w io.Writer, messageName string, fields []pbparser.FieldElement) {
	messageFields := []MessageField{}
	for _, field := range fields {
		messageFields = append(messageFields, getMessageField(messageName, field))
	}
	t := template.Must(template.ParseFiles("templates/go/MessageBuilder.template"))
	t.Execute(w, struct {
		MessageName string
		MessageFields []MessageField
	}{
		MessageName:   messageName,
		MessageFields: messageFields,
	})
}

type MessageField struct{
	FieldName string
	FieldGoType string
	IsMessage bool
	IsArray bool
	TypeAccessor string
	FieldIndex int
	MessageName string
}

func getMessageField(messageName string, field pbparser.FieldElement) MessageField {
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
			return MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  goTypes[field.Type.Name()],
				IsMessage:    false,
				IsArray:      false,
				TypeAccessor: accessor[field.Type.Name()],
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
		}
		if field.Type.Category() == pbparser.NamedDataTypeCategory {
			return MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  field.Type.Name(),
				IsMessage:    true,
				IsArray:      false,
				TypeAccessor: "Message",
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
		}
	}
	if field.Label == "repeated" {
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
			return MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  goTypes[field.Type.Name()],
				IsMessage:    false,
				IsArray:      true,
				TypeAccessor: accessor[field.Type.Name()],
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
		}
		if field.Type.Category() == pbparser.NamedDataTypeCategory {
			return MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  field.Type.Name(),
				IsMessage:    true,
				IsArray:      true,
				TypeAccessor: "Message",
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
		}
	}
	return MessageField{}
}
