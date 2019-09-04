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

func reject(t *testing.T, o1 fmt.Stringer, o2 fmt.Stringer) {
	t.Fatalf("Parsing of an old buffer by a new schema failed, expected %s but got %s", o1.String(), o2.String())
}

