// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package api

import (
	"fmt"
	"github.com/orbs-network/pbparser"
	"io"
	"os"
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

func doesFileContainBigInt(file *pbparser.ProtoFile) bool {
	for _, m := range file.Messages {
		for _, field := range m.Fields {
			if isExtendedTypeBigInt(field.Type.Name()) {
				return true
			}
		}
		for _, union := range m.OneOfs {
			for _, field := range union.Fields {
				if isExtendedTypeBigInt(field.Type.Name()) {
					return true
				}
			}
		}
	}
	return false
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
