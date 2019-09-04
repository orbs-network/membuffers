package e2e

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go/e2e/types"
	"testing"
)

func TestReadIsBackwardsCompatible_Primitives(t *testing.T) {
	foo := (&types.OnlyPrimitivesV1Builder{
		A: 42,
		B: 17,
	}).Build()
	foo2 := types.OnlyPrimitivesV2Reader(foo.Raw())

	if foo2.A() != foo.A() || foo2.B() != foo.B() {
		reject(t, foo, foo2)
	}
}

func TestReadIsBackwardsCompatible_WithBytes(t *testing.T) {
	foo := (&types.WithBytesV1Builder{
		A: 42,
		B: []byte{0x01, 0x02},
	}).Build()
	foo2 := types.WithBytesV2Reader(foo.Raw())

	if foo2.A() != foo.A() || !bytes.Equal(foo2.B(), foo.B()) {
		reject(t, foo, foo2)
	}
}

func TestReadIsBackwardsCompatible_WithInlineBytes(t *testing.T) {
	foo := (&types.WithInlineV1Builder{
		A: 42,
		B: []byte{0x01, 0x02},
	}).Build()
	foo2 := types.WithInlineV2Reader(foo.Raw())

	if foo2.A() != foo.A() || !bytes.Equal(foo2.B(), foo.B()) {
		reject(t, foo, foo2)
	}
}

func TestReadIsBackwardsCompatible_WithInlineBytes_MultiplesOf4(t *testing.T) {
	foo := (&types.WithInlineV1Builder{
		A: 42,
		B: []byte{0x01, 0x02, 0x03, 0x04},
	}).Build()
	foo2 := types.WithInlineV2Reader(foo.Raw())

	if foo2.A() != foo.A() || !bytes.Equal(foo2.B(), foo.B()) {
		reject(t, foo, foo2)
	}
}

func TestReadIsBackwardsCompatible_Nested(t *testing.T) {
	foo := (&types.NestedWithPrimitivesV1Builder{
		A: 42,
		B: &types.OnlyPrimitivesV1Builder{A: 1, B: 2},
	}).Build()
	foo2 := types.NestedWithPrimitivesV2Reader(foo.Raw())

	if foo2.A() != foo.A() || foo2.B().A() != foo.B().A() || foo2.B().B() != foo.B().B() {
		reject(t, foo, foo2)
	}
}

func TestReadIsBackwardsCompatible_NestedMiddle(t *testing.T) {
	foo1 := (&types.NestedWithPrimitivesInMiddleV1Builder{
		A: 42,
		B: &types.OnlyPrimitivesV1Builder{A: 1, B: 2},
	}).Build()
	foo2 := types.NestedWithPrimitivesInMiddleV2Reader(foo1.Raw())

	if foo2.A() != foo1.A() || foo2.B().A() != foo1.B().A() || foo2.B().B() != foo1.B().B() {
		reject(t, foo1, foo2)
	}
}

func TestReadIsBackwardsCompatible_NestedWithExtra(t *testing.T) {
	foo1 := (&types.NestedWithPrimitivesV1Builder{
		A: 42,
		B: &types.OnlyPrimitivesV1Builder{A: 1, B: 2},
	}).Build()
	foo2 := types.NestedWithPrimitivesWithExtraFieldV2Reader(foo1.Raw())

	if foo2.A() != foo1.A() || foo2.B().A() != foo1.B().A() || foo2.B().B() != foo1.B().B() {
		reject(t, foo1, foo2)
	}
}

func TestReadIsBackwardsCompatible_WithSingleBytes(t *testing.T) {
	foo1 := (&types.WithSingleByteV1Builder{
		A: 42,
		B: types.One,
		C: types.Two,
	}).Build()
	foo2 := types.WithSingleByteV2Reader(foo1.Raw())

	if foo2.A() != foo1.A() || foo2.B() != foo1.B() || foo2.C() != foo1.C() {
		reject(t, foo1, foo2)
	}
}

func reject(t *testing.T, o1 fmt.Stringer, o2 fmt.Stringer) {
	t.Fatalf("Parsing of an old buffer by a new schema failed, expected %s but got %s", o1.String(), o2.String())
}

func TestReadIsForwardsCompatible_Primitives(t *testing.T) {
	foo2 := (&types.OnlyPrimitivesV2Builder{
		A: 42,
		B: 17,
		C: 77,
	}).Build()
	foo1 := types.OnlyPrimitivesV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || foo2.B() != foo1.B() {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_WithBytes(t *testing.T) {
	foo2 := (&types.WithBytesV2Builder{
		A: 42,
		B: []byte{0x01, 0x02},
		C: 77,
	}).Build()
	foo1 := types.WithBytesV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || !bytes.Equal(foo2.B(), foo1.B()) {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_WithInlineBytes(t *testing.T) {
	foo2 := (&types.WithInlineV2Builder{
		A: 42,
		B: []byte{0x01, 0x02},
		C: 77,
	}).Build()
	foo1 := types.WithInlineV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || !bytes.Equal(foo2.B(), foo1.B()) {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_WithInlineBytes_MultiplesOf4(t *testing.T) {
	foo2 := (&types.WithInlineV2Builder{
		A: 42,
		B: []byte{0x01, 0x02, 0x03, 0x04},
		C: 77,
	}).Build()
	foo1 := types.WithInlineV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || !bytes.Equal(foo2.B(), foo1.B()) {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_Nested(t *testing.T) {
	foo2 := (&types.NestedWithPrimitivesV2Builder{
		A: 42,
		B: &types.OnlyPrimitivesV2Builder{A: 1, B: 2, C: 3},
	}).Build()
	foo1 := types.NestedWithPrimitivesV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || foo2.B().A() != foo1.B().A() || foo2.B().B() != foo1.B().B() {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_NestedMiddle(t *testing.T) {
	foo2 := (&types.NestedWithPrimitivesInMiddleV2Builder{
		A: 42,
		B: &types.OnlyPrimitivesV2Builder{A: 1, B: 2, C: 3},
		C: 77,
	}).Build()
	foo1 := types.NestedWithPrimitivesInMiddleV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || foo2.B().A() != foo1.B().A() || foo2.B().B() != foo1.B().B() || foo2.C() != foo1.C() {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_NestedWithExtra(t *testing.T) {
	foo2 := (&types.NestedWithPrimitivesWithExtraFieldV2Builder{
		A: 42,
		B: &types.OnlyPrimitivesV2Builder{A: 1, B: 2, C: 3},
		C: 77,
	}).Build()
	foo1 := types.NestedWithPrimitivesV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || foo2.B().A() != foo1.B().A() || foo2.B().B() != foo1.B().B() {
		rejectForward(t, foo1, foo2)
	}
}

func TestReadIsForwardsCompatible_WithSingleBytes(t *testing.T) {
	foo2 := (&types.WithSingleByteV2Builder{
		A: 42,
		B: types.One,
		C: types.Two,
		D: 77,
	}).Build()
	foo1 := types.WithSingleByteV1Reader(foo2.Raw())

	if foo2.A() != foo1.A() || foo2.B() != foo1.B() || foo2.C() != foo1.C() {
		rejectForward(t, foo1, foo2)
	}
}

func rejectForward(t *testing.T, oldFormat fmt.Stringer, newFormat fmt.Stringer) {
	t.Fatalf("Parsing of an new buffer by a old schema failed, expected all fields of %s to be equal to their counterparts in %s", oldFormat.String(), newFormat.String())
}
