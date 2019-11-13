// Copyright 2018 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package e2e

import (
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestComplexUnion_Message(t *testing.T) {
	cu := (&types.ComplexUnionBuilder{
		Option: types.COMPLEX_UNION_OPTION_MSG,
		Msg: &types.ExampleMessageBuilder{
			Str: "hello",
		},
	}).Build()
	if !cu.IsOptionMsg() || cu.Msg().Str() != "hello" {
		t.Fatalf("Message inside ComplexUnion is not as expected")
	}
}

func TestComplexUnion_Enum(t *testing.T) {
	cu := (&types.ComplexUnionBuilder{
		Option: types.COMPLEX_UNION_OPTION_ENU,
		Enu:    types.EXAMPLE_ENUM_OPTION_B,
	}).Build()
	if !cu.IsOptionEnu() || cu.Enu() != types.EXAMPLE_ENUM_OPTION_B {
		t.Fatalf("Enum inside ComplexUnion is not as expected")
	}
}

func TestComplexUnion_WithRepeated(t *testing.T) {
	bog := []uint32{10, 20, 30, 40}
	msgBuilder := &types.UnionWithUint32ArrayBuilder{
		Option: types.UNION_WITH_UINT_32_ARRAY_OPTION_BOG,
		Bog:    bog,
	}
	msg := msgBuilder.Build()

	raw := msg.Raw()
	t.Log(msg.String())
	t.Log(raw)

	require.Len(t, raw, 24) // no offset in end
	bogOutput := msg.BogCopiedToNative()

	require.EqualValues(t, bog, bogOutput)
	require.EqualValues(t, len(bog), len(bogOutput))

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))

	msgBuilder = &types.UnionWithUint32ArrayBuilder{
		Option: types.UNION_WITH_UINT_32_ARRAY_OPTION_BAG,
		Bag:    [][]byte{{0x1, 0x2, 0x3}, {0xaa, 0xbb}},
	}
	msg = msgBuilder.Build()
	t.Log(msg.String())
	require.NoError(t, msgBuilder.HexDump("msg ", 0))

}
