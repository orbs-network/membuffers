package main

import (
	"github.com/orbs-network/pbparser"
	"io"
	"sort"
	"text/template"
	"fmt"
	"os"
	"path"
	"strings"
	"github.com/gobuffalo/packr"
)

var box = packr.NewBox("./templates/go")
var inlineTypes = make(map[string]string)

func templateByBoxName(name string) *template.Template {
	s, err := box.MustString(name)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	t, err := template.New(name).Parse(s)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	return t
}

func convertFieldNameToGoCase(fieldName string) string {
	return ToCamel(fieldName)
}

func compileMockFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	addMockHeader(w, &file, dependencyData)
	for _, s := range file.Services {
		addMockService(w, s, &file)
	}
}

func compileProtoFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	addHeader(w, &file, dependencyData)
	for _, s := range file.Services {
		addService(w, s, &file)
	}
	for _, m := range file.Messages {
		addMessage(w, file.PackageName, m, &file)
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

func addInlineFromImports(file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	for _, dep := range dependencyData {
		importedFile, err := parseImportedFile(dep.path)
		if err != nil {
			fmt.Println("ERROR:", "imported file cannot be parsed:", dep.path, "\n", err)
			os.Exit(1)
		}
		for _, option := range importedFile.Options {
			if option.Name == "inline" && option.Value == "true" {
				for _, m := range importedFile.Messages {
					if len(m.Options) == 1 && m.Options[0].Name == "inline_type" {
						if importedFile.PackageName != file.PackageName {
							inlineTypes[importedFile.PackageName + "." + m.Name] = m.Options[0].Value
						} else {
							inlineTypes[m.Name] = m.Options[0].Value
						}
					}
				}
			}
		}
	}
}

func addEnumsFromImports(file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	for _, dep := range dependencyData {
		importedFile, err := parseImportedFile(dep.path)
		if err != nil {
			fmt.Println("ERROR:", "imported file cannot be parsed:", dep.path, "\n", err)
			os.Exit(1)
		}
		for i, enum := range importedFile.Enums {
			importedFile.Enums[i].Documentation = "imported"
			importedPackageName := path.Base(path.Dir(dep.path))
			if importedPackageName != file.PackageName {
				importedFile.Enums[i].Name = importedPackageName + "." + enum.Name
			} else {
				importedFile.Enums[i].Name = enum.Name
			}

		}
		file.Enums = append(file.Enums, importedFile.Enums...)
	}
}

func addMockHeader(w io.Writer, file *pbparser.ProtoFile, dependencyData map[string]dependencyData) {
	t := templateByBoxName("MockFileHeader.template")
	t.Execute(w, struct {
		PackageName string
	}{
		PackageName: file.PackageName,
	})
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
	addInlineFromImports(file, dependencyData)
	addEnumsFromImports(file, dependencyData)
	imports := []string{}
	for _, dependency := range file.Dependencies {
		if isInlineFileByPath(dependencyData[dependency].path) {
			continue
		}
		relative := dependencyData[dependency].relative
		packageImport := path.Dir(path.Clean(goPackage + "/" + relative + "/" + dependency))
		if packageImport != goPackage {
			imports = append(imports, packageImport)
		}
	}
	t := templateByBoxName("MessageFileHeader.template")
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
	t := templateByBoxName("MessageFileEnums.template")
	t.Execute(w, struct {
		Enums []Enum
	}{
		Enums: messageEnums,
	})
}

type ServiceMethod struct{
	Name string
	Input string
	Output string
}

func addMockService(w io.Writer, s pbparser.ServiceElement, file *pbparser.ProtoFile) {
	methods := []ServiceMethod{}
	for _, rpc := range s.RPCs {
		method := ServiceMethod{
			Name: rpc.Name,
			Input: rpc.RequestType.Name(),
			Output: rpc.ResponseType.Name(),
		}
		methods = append(methods, method)
	}
	t := templateByBoxName("MockService.template")
	t.Execute(w, struct {
		ServiceName string
		Methods []ServiceMethod
	}{
		ServiceName: s.Name,
		Methods: methods,
	})
}

func addService(w io.Writer, s pbparser.ServiceElement, file *pbparser.ProtoFile) {
	methods := []ServiceMethod{}
	for _, rpc := range s.RPCs {
		method := ServiceMethod{
			Name: rpc.Name,
			Input: rpc.RequestType.Name(),
			Output: rpc.ResponseType.Name(),
		}
		methods = append(methods, method)
	}
	t := templateByBoxName("MessageService.template")
	t.Execute(w, struct {
		ServiceName string
		Methods []ServiceMethod
	}{
		ServiceName: s.Name,
		Methods: methods,
	})
}

func addMessage(w io.Writer, packageName string, m pbparser.MessageElement, file *pbparser.ProtoFile) {
	normalizeFieldsAndOneOfs(&m)
	_, enumNameToIndex := getFileEnums(file.Enums)
	t := templateByBoxName("MessageHeader.template")
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: m.Name,
	})
	addMessageScheme(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageUnions(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageReaderHeader(w, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	for _, field := range m.Fields {
		addMessageReaderField(w, packageName, m.Name, field, m.OneOfs, enumNameToIndex)
	}
	addMessageBuilder(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
}

func addMessageScheme(w io.Writer, packageName string, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageFields := []MessageField{}
	for _, field := range fields {
		messageField := getMessageField(packageName, messageName, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
	}
	t := templateByBoxName("MessageScheme.template")
	t.Execute(w, struct {
		MessageName string
		MessageFields []MessageField
	}{
		MessageName: messageName,
		MessageFields: messageFields,
	})
}

func addMessageUnions(w io.Writer, packageName string, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	unionByIndex, _ := getMessageUnions(packageName, messageName, unions, enumNameToIndex)
	t := templateByBoxName("MessageUnions.template")
	t.Execute(w, struct {
		MessageName string
		UnionByIndex [][]MessageField
	}{
		MessageName: messageName,
		UnionByIndex: unionByIndex,
	})
}

func addMessageReaderHeader(w io.Writer, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	t := templateByBoxName("MessageReaderHeader.template")
	t.Execute(w, struct {
		MessageName string
	}{
		MessageName: messageName,
	})
}

func addMessageReaderField(w io.Writer, packageName string, messageName string, field pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageField := getMessageField(packageName, messageName, field, enumNameToIndex)
	if messageField.IsUnion {
		unionByIndex, unionNameToIndex := getMessageUnions(packageName, messageName, unions, enumNameToIndex)
		t := templateByBoxName("MessageReaderUnionField.template")
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
		t := templateByBoxName("MessageReaderMutableField.template")
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
		t := templateByBoxName("MessageReaderMessageField.template")
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
		t := templateByBoxName("MessageReaderMutableArrayField.template")
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
		t := templateByBoxName("MessageReaderMessageArrayField.template")
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

func addMessageBuilder(w io.Writer, packageName string, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	messageFields := []MessageField{}
	unionByIndex, unionNameToIndex := getMessageUnions(packageName, messageName, unions, enumNameToIndex)
	for _, field := range fields {
		messageField := getMessageField(packageName, messageName, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
	}
	t := templateByBoxName("MessageBuilder.template")
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

func getMessageUnions(packageName string, messageName string, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) ([][]MessageField, map[string]int) {
	unionByIndex := [][]MessageField{}
	unionNameToIndex := make(map[string]int)
	for _, oneOf := range unions {
		messageFields := []MessageField{}
		for _, field := range oneOf.Fields {
			messageFields = append(messageFields, getMessageField(packageName, messageName, field, enumNameToIndex))
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

func getMessageField(packageName string, messageName string, field pbparser.FieldElement, enumNameToIndex map[string]int) (messageField MessageField) {
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
	if inlineType, isInline := inlineTypes[field.Type.Name()]; isInline {
		newType, err := pbparser.NewScalarDataType(inlineType)
		if err == nil {
			field.Type = newType
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
			_, ok := accessor[field.Type.Name()]
			if !ok {
				fmt.Println("ERROR: unsupported primitive type:", field.Type.Name())
				os.Exit(1)
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
				FieldGoType:  removeLocalPackagePrefix(packageName, field.Type.Name()),
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
				FieldGoType:  removeLocalPackagePrefix(packageName, field.Type.Name()),
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

func removeLocalPackagePrefix(localPackageName string, fieldGoType string) string {
	return strings.TrimPrefix(fieldGoType, localPackageName + ".")
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
		if enum.Documentation != "imported" {
			enumByIndex = append(enumByIndex, Enum{
				Name: enum.Name,
				Values: values,
			})
		}
	}
	return enumByIndex, enumNameToIndex
}

func isInlineFileByPath(path string) bool {
	file, err := parseImportedFile(path)
	if err != nil {
		fmt.Println("ERROR:", "imported file cannot be parsed:", path, "\n", err)
		os.Exit(1)
	}
	return isInlineFile(&file)
}

type Enum struct{
	Name string
	Values []EnumValue
}

type EnumValue struct{
	Name string
	Value int
}

func parseImportedFile(path string) (pbparser.ProtoFile, error) {
	in, err := os.Open(path)
	if err != nil {
		return pbparser.ProtoFile{}, err
	}
	p := importProvider{protoFile: path, moduleToRelative: nil}
	return pbparser.Parse(in, &p)
}