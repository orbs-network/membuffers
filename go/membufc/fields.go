package main

import (
	"strings"
	"sort"
	"github.com/orbs-network/pbparser"
	"fmt"
	"os"
)

type MessageField struct{
	FieldName string
	FieldGoType string
	IsMessage bool
	IsArray bool
	IsUnion bool
	IsEnum bool
	IsInline bool
	InlineUnderlyingGoType string
	TypeAccessor string
	FieldIndex int
	MessageName string
}

type NameWithAndWithoutImport struct {
	CleanName string
	ImportName string
	MockImportName string
}

func convertFieldNameToGoCase(fieldName string) string {
	return ToCamel(fieldName)
}

func convertFieldNameToGoCaseWithPackage(fieldName string) string {
	parts := strings.Split(fieldName, ".")
	parts[len(parts) - 1] = convertFieldNameToGoCase(parts[len(parts) - 1])
	return strings.Join(parts, ".")
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
	if inlineType, isInline := globalInlineTypes[field.Type.Name()]; isInline {
		originalFieldType := field.Type.Name()
		newType, err := pbparser.NewScalarDataType(inlineType)
		if err == nil {
			field.Type = newType
			defer func() {
				messageField.IsInline = true
				messageField.InlineUnderlyingGoType = messageField.FieldGoType
				messageField.FieldGoType = convertFieldNameToGoCaseWithPackage(originalFieldType)
			}()
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

func getNameWithAndWithoutImport(name string) NameWithAndWithoutImport {
	parts := strings.Split(name, ".")
	packagePrefix := ""
	if len(parts) == 2 {
		packagePrefix = parts[0] + "."
	}
	return NameWithAndWithoutImport{
		CleanName: parts[len(parts)-1],
		ImportName: name,
		MockImportName: packagePrefix + "Mock" + parts[len(parts)-1],
	}
}
