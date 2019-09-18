// Copyright 2019 the membuffers authors
// This file is part of the membuffers library in the Orbs project.
//
// This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
// The above notice should be included in all copies or substantial portions of the software.

package e2e

import (
	"bytes"
	membuffers "github.com/orbs-network/membuffers/go"
	types "github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"github.com/stretchr/testify/require"
	"testing"
)

func getByte32Obj() [32]byte {
	return [32]byte{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
		0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf}
}

func getByte32Array() [][32]byte {
	return [][32]byte{
		{0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1, 0x1,
			0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf},
		{0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf,
			0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd, 0xe, 0xe, 0xf},
	}
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32(t *testing.T) {
	// prepare
	data := getByte32Obj()

	// build
	msgBuilder := &types.WithFixedBytes32Builder{Foo: data}
	msg := msgBuilder.Build()
	t.Log(msg.String())

	// check
	require.EqualValues(t, data, msg.Foo())
	//originalRaw := msg.Raw()
	require.True(t, bytes.Equal(msg.Raw(), data[:]), "raw message should not include size")

	// overwrite
	msg2 := types.WithFixedBytes32BuilderFromRaw(msg.Raw()).Build()
	require.EqualValues(t, data, msg2.Foo())
	require.True(t, bytes.Equal(msg2.Raw(), data[:]), "raw message should be equal to the original value")

	// mutate
	other := [32]byte{0xfe, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab,
		0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xab, 0xef}
	require.NoError(t, msg.MutateFoo(other))
	t.Log(msg.String())

	// check mutate
	require.EqualValues(t, other, msg.Foo())
	require.True(t, bytes.Equal(msg.Raw(), other[:]), "raw message should be equal to the mutated value")

	// check version with overwrite didn't change after mutate
	require.EqualValues(t, data, msg2.Foo())

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32AndUint32(t *testing.T) {
	bar := uint32(1977)
	data := getByte32Obj()
	msg := (&types.WithFixedBytes32AndUint32Builder{Foo: data, Bar: bar}).Build()

	t.Log(msg.String())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, bar, msg.Bar())

	// check hexdump
	require.NoError(t, (&types.WithFixedBytes32AndUint32Builder{Foo: data, Bar: bar}).HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32InMiddle(t *testing.T) {
	baz := uint32(1789)
	bar := "1977"
	data := getByte32Obj()
	msg := (&types.WithFixedBytes32InMiddleBuilder{Baz: baz, Foo: data, Bar: bar}).Build()

	t.Log(msg.String())
	require.EqualValues(t, baz, msg.Baz())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, bar, msg.Bar())
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32Array(t *testing.T) {
	data := getByte32Array()
	msgBuilder := &types.WithRepeatedFixedBytes32Builder{Foo: data}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	oneArray := make([]byte, 4+len(data)*32)
	membuffers.WriteOffset(oneArray, membuffers.Offset(len(data)*32))
	for i, v := range data {
		require.EqualValues(t, data[i], itr.NextFoo())
		copy(oneArray[4+i*32:4+(i+1)*32], v[:])
	}
	require.True(t, bytes.Equal(msg.Raw(), oneArray[:]), "raw message should be equal to the mutated value")

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32ArrayAndOthers(t *testing.T) {
	data := getByte32Array()
	msgBuilder := &types.WithRepeatedFixedBytes32AndOthersBuilder{Foo: data, Bar: [][]byte{{0x01, 0x02}, {0x02, 0x03}}, Baz: 1997}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	for i, _ := range data {
		require.EqualValues(t, data[i], itr.NextFoo())
	}

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}
