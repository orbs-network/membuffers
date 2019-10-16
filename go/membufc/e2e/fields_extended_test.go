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
	"math/big"
	"testing"
	"unsafe"
)

func TestMembuffersFieldsExtended_HandlesFixedBytes32(t *testing.T) {
	// prepare
	data := generateBytes32(-1)

	// build
	msgBuilder := &types.WithFixedBytes32Builder{Foo: data}
	msg := msgBuilder.Build()
	t.Log(msg.String())

	// check
	require.EqualValues(t, data, msg.Foo())
	require.True(t, bytes.Equal(msg.Raw(), data[:]), "raw message should not include size")

	// overwrite
	msg2 := types.WithFixedBytes32BuilderFromRaw(msg.Raw()).Build()
	require.EqualValues(t, data, msg2.Foo())
	require.True(t, bytes.Equal(msg2.Raw(), data[:]), "raw message should be equal to the original value")

	// mutate
	other := generateBytes32(100)
	require.NoError(t, msg.MutateFoo(other))
	t.Log(msg.String())

	// check mutate
	require.EqualValues(t, other, msg.Foo())
	require.True(t, bytes.Equal(msg.Raw(), other[:]), "raw message should be equal to the mutated value")
	require.EqualValues(t, data, msg2.Foo(), "mutate of msg.Foo should not affect msg2.Foo")

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32AndUint32(t *testing.T) {
	bar := uint32(1977)
	data := generateBytes32(-1)
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
	data := generateBytes32(-1)
	msg := (&types.WithFixedBytes32InMiddleBuilder{Baz: baz, Foo: data, Bar: bar}).Build()

	t.Log(msg.String())
	require.EqualValues(t, baz, msg.Baz())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, bar, msg.Bar())
}

func generateExectedBytes32ArrayRaw(data [][32]byte) []byte {
	fakeRaw := make([]byte, 4+len(data)*32)
	membuffers.WriteOffset(fakeRaw, membuffers.Offset(len(data)*32)) // size of actual data not the size of the array
	for i, v := range data {
		copy(fakeRaw[4+i*32:4+(i+1)*32], v[:])
	}
	return fakeRaw
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32Array(t *testing.T) {
	data := generateBytes32Array(3)
	msgBuilder := &types.WithRepeatedFixedBytes32Builder{Foo: data}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	for _, oneBytes32 := range data {
		require.EqualValues(t, oneBytes32, itr.NextFoo())
	}
	require.True(t, bytes.Equal(msg.Raw(), generateExectedBytes32ArrayRaw(data)), "raw message should be equal to the fake raw")

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes32ArrayAndOthers(t *testing.T) {
	data := generateBytes32Array(3)
	msgBuilder := &types.WithRepeatedFixedBytes32AndOthersBuilder{Foo: data, Bar: [][]byte{{0x01, 0x02}, {0x02, 0x03}}, Baz: 1997}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	for i := range data {
		require.EqualValues(t, data[i], itr.NextFoo())
	}

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesFixedBytes20InMiddle(t *testing.T) {
	baz := uint32(1789)
	bar := "1977"
	data := generateBytes20(-1)
	msg := (&types.WithFixedBytes20InMiddleBuilder{Baz: baz, Foo: data, Bar: bar}).Build()

	t.Log(msg.String())
	require.EqualValues(t, baz, msg.Baz())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, bar, msg.Bar())
}

func TestMembuffersFieldsExtended_HandlesFixedBytes20ArrayAndOthers(t *testing.T) {
	data := generateBytes20Array(3)
	msgBuilder := &types.WithRepeatedFixedBytes20AndOthersBuilder{Foo: data, Bar: [][]byte{{0x01, 0x02}, {0x02, 0x03}}, Baz: 1997}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	for i := range data {
		require.EqualValues(t, data[i], itr.NextFoo())
	}

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesEnumAndFixedBytes32(t *testing.T) {
	data := generateBytes32(-1)
	msgBuilder := &types.WithEnumAndFixedBytes32Builder{Foo: data, Bar: types.FIXED_EXAMPLE_ENUM_OPTION_C}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, types.FIXED_EXAMPLE_ENUM_OPTION_C, msg.Bar())

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesUint256(t *testing.T) {
	baz := uint32(1789)
	bar := "1977"
	data := big.NewInt(500000000)
	msgBuilder := types.WithUint256InMiddleBuilder{Baz: baz, Foo: data, Bar: bar}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	require.EqualValues(t, baz, msg.Baz())
	require.EqualValues(t, data, msg.Foo())
	require.EqualValues(t, bar, msg.Bar())

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesUint256ArrayAndOthers(t *testing.T) {
	data := []*big.Int{big.NewInt(500000000), big.NewInt(600000000), big.NewInt(700000000)}
	msgBuilder := &types.WithRepeatedUint256AndOthersBuilder{Foo: data, Bar: 1997}
	msg := msgBuilder.Build()

	t.Log(msg.String())
	t.Log(msg.Raw())
	itr := msg.FooIterator()
	for i := range data {
		require.EqualValues(t, data[i], itr.NextFoo())
	}

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

func TestMembuffersFieldsExtended_HandlesBoolsAndOthers(t *testing.T) {
	foo, bar, qux, thud := false, true, true, true
	baz, quux := uint32(1977), uint32(1889)
	msgBuilder := &types.WithBoolsAndOthersBuilder{Foo: foo, Bar: bar, Baz: baz, Qux: qux, Quux: quux, Thud: thud}
	msg := msgBuilder.Build()

	raw := msg.Raw()
	t.Log(msg.String())
	t.Log(raw)

	require.Len(t, raw, 17)
	require.True(t, foo == msg.Foo())
	require.True(t, bar == msg.Bar())
	require.EqualValues(t, baz, msg.Baz())
	require.True(t, qux == msg.Qux())
	require.EqualValues(t, quux, msg.Quux())
	require.True(t, thud == msg.Thud())

	// check hexdump
	require.NoError(t, msgBuilder.HexDump("msg ", 0))
}

// Helpers
func generateBytes(byteValue int, size int) []byte {
	out := make([]byte, size)
	for i := 0; i < size; i++ {
		if byteValue < 0 {
			out[i] = byte(i + 1)
		} else {
			out[i] = byte(byteValue)
		}
	}
	return out
}

func generateBytes20(byteValue int) [20]byte {
	return *(*[20]byte)(unsafe.Pointer(&generateBytes(byteValue, 20)[0]))
}

func generateBytes20Array(arraySize int) [][20]byte {
	out := make([][20]byte, 0, arraySize)
	for i := 0; i < arraySize; i++ {
		c := *(*[20]byte)(unsafe.Pointer(&generateBytes(i+1, 20)[0]))
		out = append(out, c)
	}
	return out
}

func generateBytes32(byteValue int) [32]byte {
	return *(*[32]byte)(unsafe.Pointer(&generateBytes(byteValue, 32)[0]))
}

func generateBytes32Array(arraySize int) [][32]byte {
	out := make([][32]byte, 0, arraySize)
	for i := 0; i < arraySize; i++ {
		c := *(*[32]byte)(unsafe.Pointer(&generateBytes(i+1, 32)[0]))
		out = append(out, c)
	}
	return out
}
