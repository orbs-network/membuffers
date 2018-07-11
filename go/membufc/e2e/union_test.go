package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
)

func TestComplexUnionMessage(t *testing.T) {
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

func TestComplexUnionEnum(t *testing.T) {
	cu := (&types.ComplexUnionBuilder{
		Option: types.COMPLEX_UNION_OPTION_ENU,
		Enu: types.EXAMPLE_ENUM_OPTION_B,
	}).Build()
	if !cu.IsOptionEnu() || cu.Enu() != types.EXAMPLE_ENUM_OPTION_B {
		t.Fatalf("Enum inside ComplexUnion is not as expected")
	}
}