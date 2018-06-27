package main

import (
	"github.com/tallstoat/pbparser"
	"io"
	"sort"
	"text/template"
	"fmt"
	"os"
	"path"
	"strings"
)

func convertFieldNameToGoCase(fieldName string) string {
	return ToCamel(fieldName)
}

func compileProtoFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	addHeader(w, &file, dependencyData)
	for _, m := range file.Messages {
		addMessage(w, m, &file)
	}
	addEnums(w, file.Enums)
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

func addEnumsFromImports(file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	for _, dep := range dependencyData {
		importedFile, err := pbparser.ParseFile(dep.path)
		if err != nil {
			fmt.Println("ERROR:", "imported file cannot be parsed: %s", dep.path)
			os.Exit(1)
		}
		for i, enum := range importedFile.Enums {
			importedFile.Enums[i].Name = path.Base(path.Dir(dep.path)) + "." + enum.Name
		}
		file.Enums = append(file.Enums, importedFile.Enums...)
	}
}

func addHeader(w io.Writer, file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	var goPackage string
	for _, option := range file.Options {
		if option.Name == "go_package" {
			goPackage = option.Value
		}
	}
	if len(file.Dependencies) > 0 && len(goPackage) == 0 {
		fmt.Println("ERROR:", "option go_package not provided, required when we have imports")
		os.Exit(1)
	}
	addEnumsFromImports(file, dependencyData)
	imports := []string{}
	for _, dependency := range file.Dependencies {
		relative := dependencyData[dependency].relative
		imports = append(imports, path.Dir(path.Clean(goPackage + "/" + relative + "/" + dependency)))
	}
	t := template.Must(template.ParseFiles("templates/go/MessageFileHeader.template"))
	t.Execute(w, struct {
		PackageName string
		Imports []string
		HasMessages bool
	}{
		PackageName: file.PackageName,
		Imports: imports,
		HasMessages: len(file.Messages) > 0,
	})
}

func addEnums(w io.Writer, enums []pbparser.EnumElement) {
	if len(enums) == 0 {
		return
	}
	messageEnums, _ := getFileEnums(enums)
	t := template.Must(template.ParseFiles("templates/go/MessageFileEnums.template"))
	t.Execute(w, struct {
		Enums []Enum
	}{
		Enums: messageEnums,
	})
}

func addMessage(w io.Writer, m pbparser.MessageElement, file *pbparser.ProtoFile) {
	normalizeFieldsAndOneOfs(&m)
	_, enumNameToIndex := getFileEnums(file.Enums)
	t := template.Must(template.ParseFiles("templates/go/MessageHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: m.Name,
	})
	addMessageScheme(w, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageUnions(w, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageReaderHeader(w, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	for _, field := range m.Fields {
		addMessageReaderField(w, m.Name, field, m.OneOfs, enumNameToIndex)
	}
	addMessageBuilder(w, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
}

func addMessageScheme(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageFields := []MessageField{}
	for _, field := range fields {
		messageField := getMessageField(messageName, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
	}
	t := template.Must(template.ParseFiles("templates/go/MessageScheme.template"))
	t.Execute(w, struct {
		MessageName string
		MessageFields []MessageField
	}{
		MessageName: messageName,
		MessageFields: messageFields,
	})
}

func addMessageUnions(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	unionByIndex, _ := getMessageUnions(messageName, unions, enumNameToIndex)
	t := template.Must(template.ParseFiles("templates/go/MessageUnions.template"))
	t.Execute(w, struct {
		MessageName string
		UnionByIndex [][]MessageField
	}{
		MessageName: messageName,
		UnionByIndex: unionByIndex,
	})
}

func addMessageReaderHeader(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	t := template.Must(template.ParseFiles("templates/go/MessageReaderHeader.template"))
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: messageName,
	})
}

func addMessageReaderField(w io.Writer, messageName string, field pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageField := getMessageField(messageName, field, enumNameToIndex)
	if messageField.IsUnion {
		unionByIndex, unionNameToIndex := getMessageUnions(messageName, unions, enumNameToIndex)
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

func addMessageBuilder(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageFields := []MessageField{}
	unionByIndex, unionNameToIndex := getMessageUnions(messageName, unions, enumNameToIndex)
	for _, field := range fields {
		messageField := getMessageField(messageName, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
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

func getMessageUnions(messageName string, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) ([][]MessageField, map[string]int) {
	unionByIndex := [][]MessageField{}
	unionNameToIndex := make(map[string]int)
	for _, oneOf := range unions {
		messageFields := []MessageField{}
		for _, field := range oneOf.Fields {
			messageFields = append(messageFields, getMessageField(messageName, field, enumNameToIndex))
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
	IsEnum bool
	TypeAccessor string
	FieldIndex int
	MessageName string
}

func getMessageField(messageName string, field pbparser.FieldElement, enumNameToIndex map[string]int) (messageField MessageField) {
	defer func() {
		if _, ok := enumNameToIndex[messageField.FieldGoType]; ok {
			messageField.IsEnum = true
			messageField.IsMessage = false
			messageField.TypeAccessor = "Uint16"
		}
	}()
	if field.Label == "oneof" {
		messageField = MessageField{
			FieldName:    convertFieldNameToGoCase(field.Name),
			FieldGoType:  messageName + convertFieldNameToGoCase(field.Name),
			IsMessage:    false,
			IsArray:      false,
			IsUnion:      true,
			TypeAccessor: "Union",
			FieldIndex:   field.Tag,
			MessageName:  messageName,
		}
		return
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
			messageField = MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  goTypes[field.Type.Name()],
				IsMessage:    false,
				IsArray:      false,
				IsUnion:      false,
				TypeAccessor: accessor[field.Type.Name()],
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
			return
		}
		if field.Type.Category() == pbparser.NamedDataTypeCategory {
			messageField = MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  field.Type.Name(),
				IsMessage:    true,
				IsArray:      false,
				IsUnion:      false,
				TypeAccessor: "Message",
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
			return
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
			messageField = MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  goTypes[field.Type.Name()],
				IsMessage:    false,
				IsArray:      true,
				IsUnion:      false,
				TypeAccessor: accessor[field.Type.Name()],
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
			return
		}
		if field.Type.Category() == pbparser.NamedDataTypeCategory {
			messageField = MessageField{
				FieldName:    convertFieldNameToGoCase(field.Name),
				FieldGoType:  field.Type.Name(),
				IsMessage:    true,
				IsArray:      true,
				IsUnion:      false,
				TypeAccessor: "Message",
				FieldIndex:   field.Tag,
				MessageName:  messageName,
			}
			return
		}
	}
	return MessageField{}
}

func getFileEnums(enums []pbparser.EnumElement) ([]Enum, map[string]int) {
	enumByIndex := []Enum{}
	enumNameToIndex := make(map[string]int)
	for _, enum := range enums {
		enumNameToIndex[enum.Name] = len(enumByIndex)
		values := []EnumValue{}
		for _, value := range enum.EnumConstants {
			values = append(values, EnumValue{
				Name: value.Name,
				Value: value.Tag,
			})
		}
		// only add here enums from this package
		if !strings.Contains(enum.Name, ".") {
			enumByIndex = append(enumByIndex, Enum{
				Name: enum.Name,
				Values: values,
			})
		}
	}
	return enumByIndex, enumNameToIndex
}

type Enum struct{
	Name string
	Values []EnumValue
}

type EnumValue struct{
	Name string
	Value int
}