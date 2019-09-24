// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package main

import "github.com/orbs-network/pbparser"

func isExtendedType(name string) bool {
	return name == "membuffers.bytes32" || name == "membuffers.bytes20"
}

func getExtendedType(messageName string, field pbparser.FieldElement, isArray bool) MessageField {
	if field.Type.Name() == "membuffers.bytes32" {
		return MessageField{
			FieldName:    convertFieldNameToGoCase(field.Name),
			FieldGoType:  "[32]byte",
			IsMessage:    false,
			IsArray:      isArray,
			IsUnion:      false,
			TypeAccessor: "Bytes32",
			FieldIndex:   field.Tag,
			MessageName:  messageName,
		}
	} else if field.Type.Name() == "membuffers.bytes20" {
		return MessageField{
			FieldName:    convertFieldNameToGoCase(field.Name),
			FieldGoType:  "[20]byte",
			IsMessage:    false,
			IsArray:      isArray,
			IsUnion:      false,
			TypeAccessor: "Bytes20",
			FieldIndex:   field.Tag,
			MessageName:  messageName,
		}
	}
	return MessageField{}
}
