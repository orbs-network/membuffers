package main

import (
	"fmt"
	"os"
	"github.com/orbs-network/pbparser"
	"io"
)

func shouldSerializeAllMessages(file *pbparser.ProtoFile) bool {
	for _, option := range file.Options {
		if option.Name == "serialize_messages" && option.Value == "false" {
			return false
		}
	}
	return true
}

func shouldSerializeMessage(m pbparser.MessageElement) bool {
	for _, option := range m.Options {
		if option.Name == "serialize_message" && option.Value == "false" {
			return false
		}
	}
	return true
}

func doesFileContainHandlers(path string, handlers []NameWithAndWithoutImport) bool {
	file, err := parseImportedFile(path)
	if err != nil {
		fmt.Println("ERROR:", "imported file cannot be parsed:", path, "\n", err)
		os.Exit(1)
	}
	for _, handler := range handlers {
		for _, service := range file.Services {
			if handler.CleanName == service.Name {
				return true
			}
		}
	}
	return false
}

func addMessageNonSerializable(w io.Writer, packageName string, m pbparser.MessageElement, file *pbparser.ProtoFile) {
	messageFields := []MessageField{}
	messageName := m.Name
	fields := m.Fields
	unions := m.OneOfs
	normalizeFieldsAndOneOfs(&m)
	_, enumNameToIndex := getFileEnums(file.Enums)
	unionByIndex, unionNameToIndex := getMessageUnions(packageName, messageName, unions, enumNameToIndex)
	for _, field := range fields {
		messageField := getMessageField(packageName, messageName, field, enumNameToIndex)
		messageFields = append(messageFields, messageField)
	}
	t := templateByBoxName("MessageNonSerializable.template")
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
