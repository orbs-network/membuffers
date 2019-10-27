// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package tests

import (
	"github.com/orbs-network/membuffers/tests/types"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func makeBytes32ObjectWithByte(b byte) [32]byte {
	var obj = [32]byte{}
	for i := 0; i < 32; i++ {
		obj[i] = b
	}
	return obj
}

func TestMembuffers_ExtendedFieldTypes_JSreadsBytes32(t *testing.T) {
	obj := (&types.WithFixedBytes32AndUint32Builder{A: 24041977, B: makeBytes32ObjectWithByte(0xab)}).Build()
	t.Log(obj.Raw())
	jsonFromJs := readInJs(t, obj.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeBytes32], []);
const obj = {};
obj.A = m.getUint32(0)
obj.B = Array.from(m.getBytes32(1))`)

	require.EqualValues(t, obj.A(), jsonFromJs["A"], "field A missing in object read in JS version")
	array := jsonFromJs["B"].([]interface{})
	require.Len(t, jsonFromJs["B"], 32, "field B must be len 32 ")
	for i, b := range obj.B() {
		require.EqualValues(t, b, array[i])
	}
}

func makeBytes20ObjectWithByte(b byte) [20]byte {
	var obj = [20]byte{}
	for i := 0; i < 20; i++ {
		obj[i] = b
	}
	return obj
}

func TestMembuffers_ExtendedFieldTypes_JSreadsBytes20(t *testing.T) {
	obj := (&types.WithFixedBytes20AndUint32Builder{A: 24041977, B: makeBytes20ObjectWithByte(0xab)}).Build()
	t.Log(obj.Raw())
	jsonFromJs := readInJs(t, obj.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeBytes20], []);
const obj = {};
obj.A = m.getUint32(0)
obj.B = Array.from(m.getBytes20(1))`)

	require.EqualValues(t, obj.A(), jsonFromJs["A"], "field A missing in object read in JS version")
	array := jsonFromJs["B"].([]interface{})
	require.Len(t, jsonFromJs["B"], 20, "field B must be len 20 ")
	for i, b := range obj.B() {
		require.EqualValues(t, b, array[i])
	}
}

func TestMembuffers_ExtendedFieldTypes_JSreadsUint256(t *testing.T) {
	obj := (&types.WithFixedUint256AndUint32Builder{A: 2, B: big.NewInt(8224), C: 2}).Build()
	t.Log(obj.Raw())
	jsonFromJs := readInJs(t, obj.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeUint256, FieldTypes.TypeUint32], []);
const obj = {};
obj.A = m.getUint32(0);
obj.B = m.getUint256(1);
obj.C = m.getUint32(2);`)

	require.EqualValues(t, obj.A(), jsonFromJs["A"], "field A missing in object read in JS version")
	require.EqualValues(t, obj.B().String(), jsonFromJs["B"], "field B missing in object read in JS version") // for test we convert js bigint to string, it has no default stringify
	require.EqualValues(t, obj.C(), jsonFromJs["C"], "field C missing in object read in JS version")
}

func TestMembuffers_ExtendedFieldTypes_JSreadsBool(t *testing.T) {
	obj := (&types.WithBoolsAndOthersBuilder{A: false, B: true, C: 1977, D: true, E: 1889, F: true}).Build()
	t.Log(obj.Raw())
	jsonFromJs := readInJs(t, obj.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeBool, FieldTypes.TypeBool, FieldTypes.TypeUint32, FieldTypes.TypeBool, FieldTypes.TypeUint32, FieldTypes.TypeBool], []);
const obj = {};
obj.A = m.getBool(0);
obj.B = m.getBool(1);
obj.C = m.getUint32(2);
obj.D = m.getBool(3);
obj.E = m.getUint32(4);
obj.F = m.getBool(5);
`)
	require.EqualValues(t, obj.A(), jsonFromJs["A"], "field A missing in object read in JS version")
	require.EqualValues(t, obj.B(), jsonFromJs["B"], "field B missing in object read in JS version")
	require.EqualValues(t, obj.C(), jsonFromJs["C"], "field C missing in object read in JS version")
	require.EqualValues(t, obj.D(), jsonFromJs["D"], "field D missing in object read in JS version")
	require.EqualValues(t, obj.E(), jsonFromJs["E"], "field E missing in object read in JS version")
	require.EqualValues(t, obj.F(), jsonFromJs["F"], "field F missing in object read in JS version")
}
