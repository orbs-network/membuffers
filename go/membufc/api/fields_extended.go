// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package api

import "github.com/orbs-network/pbparser"

const Bytes20 = "membuffers.bytes20"
const Bytes32 = "membuffers.bytes32"
const Uint256 = "membuffers.uint256"

func isExtendedType(name string) bool {
	return name == Bytes32 || name == Bytes20 || name == Uint256
}

func isExtendedTypeBigInt(name string) bool {
	return name == Uint256
}

func getExtendedType(messageName string, field pbparser.FieldElement, isArray bool) MessageField {
	m := MessageField{
		FieldName:   convertFieldNameToGoCase(field.Name),
		IsMessage:   false,
		IsArray:     isArray,
		IsUnion:     false,
		FieldIndex:  field.Tag,
		MessageName: messageName,
	}
	switch field.Type.Name() {
	case Bytes32:
		m.FieldGoType = "[32]byte"
		m.TypeAccessor = "Bytes32"
		return m
	case Bytes20:
		m.FieldGoType = "[20]byte"
		m.TypeAccessor = "Bytes20"
		return m
	case Uint256:
		m.FieldGoType = "*big.Int"
		m.TypeAccessor = "Uint256"
		return m
	}
	return MessageField{}
}
