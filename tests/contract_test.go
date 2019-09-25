// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package tests

import (
	"github.com/orbs-network/membuffers/tests/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMembuffers_BackwardsCompatibility_JSWithNewSchemaReadsOlderSchema_Primitives(t *testing.T) {
	v1 := (&types.SomeObjectV1Builder{A: 42}).Build()
	jsonFromJs := readInJs(t, v1.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeUint32], []);
const obj = {};
obj.A = m.getUint32(0)
obj.B = m.getUint32(1)`)

	require.EqualValues(t, v1.A(), jsonFromJs["A"], "field missing in object read in JS version")
	require.EqualValues(t, 0, jsonFromJs["B"], "did not get zero value from missing field read in JS version")
}

func TestMembuffers_BackwardsCompatibility_JSWithOldSchemaReadsNewerSchema_Primitives(t *testing.T) {

	v2 := (&types.SomeObjectV2Builder{A: 42, B: 17}).Build()
	jsonFromJs := readInJs(t, v2.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32], []);
const obj = {};
obj.A = m.getUint32(0)`)

	require.EqualValues(t, v2.A(), jsonFromJs["A"], "field missing in object read in JS version")
}

func TestMembuffers_BackwardsCompatibility_JSWithNewSchemaReadsOlderSchema_WithBytes(t *testing.T) {
	v1 := (&types.WithBytesV1Builder{A: 42, B: []byte{0x01, 0x02}}).Build()
	jsonFromJs := readInJs(t, v1.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeBytes, FieldTypes.TypeUint32], []);
const obj = {};
obj.A = m.getUint32(0)
obj.B = Array.from(m.getBytes(1))`)

	require.EqualValues(t, v1.A(), jsonFromJs["A"], "field missing in object read in JS version")
	b := jsonFromJs["B"].([]interface{})
	require.Equal(t, len(v1.B()), len(b), "inconsistent array lengths in object read from JS version")
	require.EqualValues(t, v1.B()[0], b[0], "inconsistent array value in object read from JS version")
	require.EqualValues(t, v1.B()[1], b[1], "inconsistent array value in object read from JS version")
}

func TestMembuffers_BackwardsCompatibility_JSWithOldSchemaReadsNewerSchema_WithBytes(t *testing.T) {
	v2 := (&types.WithBytesV2Builder{A: 42, B: []byte{0x01, 0x02}}).Build()
	jsonFromJs := readInJs(t, v2.Raw(), `
const m = new InternalMessage(buf, buf.byteLength, [FieldTypes.TypeUint32, FieldTypes.TypeBytes], []);
const obj = {};
obj.A = m.getUint32(0)
obj.B = Array.from(m.getBytes(1))`)

	require.EqualValues(t, v2.A(), jsonFromJs["A"], "field missing in object read in JS version")
	b := jsonFromJs["B"].([]interface{})
	require.Equal(t, len(v2.B()), len(b), "inconsistent array lengths in object read from JS version")
	require.EqualValues(t, v2.B()[0], b[0], "inconsistent array value in object read from JS version")
	require.EqualValues(t, v2.B()[1], b[1], "inconsistent array value in object read from JS version")
}
