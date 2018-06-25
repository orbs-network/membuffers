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

func normalizeFieldsAndOneOfs(m *pbparser.MessageElement) {
	for _, oneOf := range m.OneOfs {
		if len(oneOf.Fields) == 0 {
			continue
		}
		m.Fields = append(m.Fields, pbparser.FieldElement{
			Label: "oneof",
			Tag: oneOf.Fields[0].Tag,
			Name: oneOf.Name,
		})
	}
	sort.Slice(m.Fields, func(i, j int) bool {
		return m.Fields[i].Tag < m.Fields[j].Tag
	})
	sort.Slice(m.OneOfs, func(i, j int) bool {
		return m.OneOfs[i].Fields[0].Tag < m.OneOfs[j].Fields[0].Tag
	})
	for i, _ := range m.Fields {
		m.Fields[i].Tag = i
	}
	for i, _ := range m.OneOfs {
		for j, _ := range m.OneOfs[i].Fields {
			m.OneOfs[i].Fields[j].Tag = j
		}
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
	normalizeFieldsAndOneOfs(&m)
	t := template.Must(template.ParseFiles("templates/go/MessageHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: m.Name,
	})
	addMessageScheme(w, m.Name, m.Fields, m.OneOfs)
	addMessageUnions(w, m.Name, m.Fields, m.OneOfs)
	addMessageReaderHeader(w, m.Name, m.Fields, m.OneOfs)
	for _, field := range m.Fields {
		addMessageReaderField(w, m.Name, field, m.OneOfs)
	}
	addMessageBuilder(w, m.Name, m.Fields, m.OneOfs)
}

func addMessageScheme(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement) {
	fieldTypes := []string{}
	for _, field := range fields {
		if field.Label == "oneof" {
			fieldTypes = append(fieldTypes, "TypeUnion")
		}
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

func addMessageUnions(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement) {
	unionByIndex, _ := getMessageUnions(messageName, unions)
	t := template.Must(template.ParseFiles("templates/go/MessageUnions.template"))
	t.Execute(w, struct {
		MessageName string
		UnionByIndex [][]MessageField
	}{
		MessageName: messageName,
		UnionByIndex: unionByIndex,
	})
}

func addMessageReaderHeader(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement) {
	t := template.Must(template.ParseFiles("templates/go/MessageReaderHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: messageName,
	})
}

func addMessageReaderField(w io.Writer, messageName string, field pbparser.FieldElement, unions []pbparser.OneOfElement) {
	messageField := getMessageField(messageName, field)
	if messageField.IsUnion {
		unionByIndex, unionNameToIndex := getMessageUnions(messageName, unions)
		t := template.Must(template.ParseFiles("templates/go/MessageReaderUnionField.template"))
		t.Execute(w, struct {
			MessageName string
			UnionName string
			FieldIndex int
			UnionNum int
			UnionFields []MessageField
		}{
			MessageName:  messageName,
			UnionName: messageField.FieldName,
			FieldIndex: messageField.FieldIndex,
 			UnionNum: unionNameToIndex[messageField.FieldName],
 			UnionFields: unionByIndex[unionNameToIndex[messageField.FieldName]],
		})
		return
	}
	if !messageField.IsMessage && !messageField.IsArray {
		t := template.Must(template.ParseFiles("templates/go/MessageReaderMutableField.template"))
		t.Execute(w, struct {
			MessageName string
			MessageField MessageField
		}{
			MessageName:  messageName,
			MessageField: messageField,
		})
		return
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
		return
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
		return
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
		return
	}
}

func addMessageBuilder(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement) {
	messageFields := []MessageField{}
	unionByIndex, unionNameToIndex := getMessageUnions(messageName, unions)
	for _, field := range fields {
		messageFields = append(messageFields, getMessageField(messageName, field))
	}
	t := template.Must(template.ParseFiles("templates/go/MessageBuilder.template"))
	t.Execute(w, struct {
		MessageName string
		MessageFields []MessageField
		UnionByIndex [][]MessageField
		UnionNameToIndex map[string]int
	}{
		MessageName:   messageName,
		MessageFields: messageFields,
		UnionByIndex:  unionByIndex,
		UnionNameToIndex: unionNameToIndex,
	})
}

func getMessageUnions(messageName string, unions []pbparser.OneOfElement) ([][]MessageField, map[string]int) {
	unionByIndex := [][]MessageField{}
	unionNameToIndex := make(map[string]int)
	for _, oneOf := range unions {
		messageFields := []MessageField{}
		for _, field := range oneOf.Fields {
			messageFields = append(messageFields, getMessageField(messageName, field))
		}
		unionNameToIndex[oneOf.Name] = len(unionByIndex)
		unionByIndex = append(unionByIndex, messageFields)
	}
	return unionByIndex, unionNameToIndex
}

type MessageField struct{
	FieldName string
	FieldGoType string
	IsMessage bool
	IsArray bool
	IsUnion bool
	TypeAccessor string
	FieldIndex int
	MessageName string
}

func getMessageField(messageName string, field pbparser.FieldElement) MessageField {
	if field.Label == "oneof" {
		return MessageField{
			FieldName:    convertFieldNameToGoCase(field.Name),
			FieldGoType:  messageName + convertFieldNameToGoCase(field.Name),
			IsMessage:    false,
			IsArray:      false,
			IsUnion:      true,
			TypeAccessor: "Union",
			FieldIndex:   field.Tag,
			MessageName:  messageName,
		}
	}
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
				IsUnion:      false,
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
				IsUnion:      false,
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
				IsUnion:      false,
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
				IsUnion:      false,
				TypeAccessor: "Message",
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
		}
	}
	return MessageField{}
}
