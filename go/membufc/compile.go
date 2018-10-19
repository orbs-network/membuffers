package main

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"github.com/orbs-network/pbparser"
	"io"
	"os"
	"path"
	"text/template"
)

var box = packr.NewBox("./templates/go")

func templateByBoxName(name string) *template.Template {
	s, err := box.MustString(name)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"capsnake": ToSnake,
	}
	t, err := template.New(name).Funcs(funcMap).Parse(s)
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		os.Exit(1)
	}
	return t
}

func compileProtoFile(w io.Writer, file pbparser.ProtoFile, dependencyData map[string]dependencyData, compilerVersion string, languageGoCtx bool) {
	serializeAllMessages := shouldSerializeAllMessages(&file)
	addHeader(w, &file, dependencyData, serializeAllMessages, compilerVersion, languageGoCtx)
	for _, s := range file.Services {
		addService(w, file.PackageName, s, &file, languageGoCtx)
	}
	for _, m := range file.Messages {
		if !serializeAllMessages || !shouldSerializeMessage(m) {
			addMessageNonSerializable(w, file.PackageName, m, &file)
		} else {
			addMessage(w, file.PackageName, m, &file)
		}
	}
	addEnums(w, file.Enums)
}

func addHeader(w io.Writer, file *pbparser.ProtoFile, dependencyData map[string]dependencyData, serializeServiceArgs bool, compilerVersion string, languageGoCtx bool) {
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
		relative := dependencyData[dependency].relative
		packageImport := path.Dir(path.Clean(goPackage + "/" + relative + "/" + dependency))
		if packageImport != goPackage {
			imports = append(imports, packageImport)
		}
	}
	t := templateByBoxName("MessageFileHeader.template")
	t.Execute(w, struct {
		PackageName       string
		Imports           []string
		HasMembuffers     bool
		HasMessages       bool
		HasServiceMethods bool
		CompilerVersion   string
		LanguageGoCtx     bool
	}{
		PackageName:       file.PackageName,
		Imports:           unique(imports),
		HasMembuffers:     len(file.Messages) > 0 && serializeServiceArgs,
		HasMessages:       len(file.Messages) > 0,
		HasServiceMethods: fileHasServiceMethods(file),
		CompilerVersion:   compilerVersion,
		LanguageGoCtx:     languageGoCtx,
	})
}

type ServiceMethod struct {
	Name   string
	Input  string
	Output string
}

func addService(w io.Writer, packageName string, s pbparser.ServiceElement, file *pbparser.ProtoFile, languageGoCtx bool) {
	methods := []ServiceMethod{}
	for _, rpc := range s.RPCs {
		method := ServiceMethod{
			Name:   rpc.Name,
			Input:  removeLocalPackagePrefix(packageName, rpc.RequestType.Name()),
			Output: removeLocalPackagePrefix(packageName, rpc.ResponseType.Name()),
		}
		methods = append(methods, method)
	}
	registerHandlers := []NameWithAndWithoutImport{}
	implementHandlers := []NameWithAndWithoutImport{}
	for _, option := range s.Options {
		if option.Name == "register_handler" {
			registerHandlers = append(registerHandlers, getNameWithAndWithoutImport(option.Value))
		}
		if option.Name == "implement_handler" {
			implementHandlers = append(implementHandlers, getNameWithAndWithoutImport(option.Value))
		}
	}
	t := templateByBoxName("MessageService.template")
	t.Execute(w, struct {
		ServiceName       string
		Methods           []ServiceMethod
		RegisterHandlers  []NameWithAndWithoutImport
		ImplementHandlers []NameWithAndWithoutImport
		LanguageGoCtx     bool
	}{
		ServiceName:       s.Name,
		Methods:           methods,
		RegisterHandlers:  registerHandlers,
		ImplementHandlers: implementHandlers,
		LanguageGoCtx:     languageGoCtx,
	})
}

func addMessage(w io.Writer, packageName string, m pbparser.MessageElement, file *pbparser.ProtoFile) {
	normalizeFieldsAndOneOfs(&m)
	_, enumNameToIndex := getFileEnums(file.Enums)
	messageFields := []MessageField{}
	for _, field := range m.Fields {
		messageField := getMessageField(packageName, m.Name, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
	}
	t := templateByBoxName("MessageHeader.template")
	t.Execute(w, struct {
		MessageName   string
		MessageFields []MessageField
	}{
		MessageName:   m.Name,
		MessageFields: messageFields,
	})
	addMessageScheme(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageUnions(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
	addMessageReaderHeader(w, packageName, m.Name, m.Fields, m.OneOfs, enumNameToIndex)
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
		MessageName   string
		MessageFields []MessageField
	}{
		MessageName:   messageName,
		MessageFields: messageFields,
	})
}

func addMessageUnions(w io.Writer, packageName string, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
	unionByIndex, _ := getMessageUnions(packageName, messageName, unions, enumNameToIndex)
	t := templateByBoxName("MessageUnions.template")
	t.Execute(w, struct {
		MessageName  string
		UnionByIndex [][]MessageField
	}{
		MessageName:  messageName,
		UnionByIndex: unionByIndex,
	})
}

func addMessageReaderHeader(w io.Writer, packageName string, messageName string, fields []pbparser.FieldElement, unions []pbparser.OneOfElement, enumNameToIndex map[string]int) {
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
			UnionName   string
			FieldIndex  int
			UnionNum    int
			UnionFields []MessageField
		}{
			MessageName: messageName,
			UnionName:   messageField.FieldName,
			FieldIndex:  messageField.FieldIndex,
			UnionNum:    unionNameToIndex[messageField.FieldName],
			UnionFields: unionByIndex[unionNameToIndex[messageField.FieldName]],
		})
		return
	}
	if !messageField.IsMessage && !messageField.IsArray {
		t := templateByBoxName("MessageReaderMutableField.template")
		t.Execute(w, struct {
			MessageName  string
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
			MessageName  string
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
			MessageName  string
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
			MessageName  string
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
		MessageName      string
		MessageFields    []MessageField
		UnionByIndex     [][]MessageField
		UnionNameToIndex map[string]int
	}{
		MessageName:      messageName,
		MessageFields:    messageFields,
		UnionByIndex:     unionByIndex,
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

func parseImportedFile(path string) (pbparser.ProtoFile, error) {
	in, err := os.Open(path)
	if err != nil {
		return pbparser.ProtoFile{}, err
	}
	p := importProvider{protoFile: path, moduleToRelative: nil}
	return pbparser.Parse(in, &p)
}
