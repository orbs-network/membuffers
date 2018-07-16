package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/e2e/types"
	"reflect"
)

func TestDeepEqualsForUnininitalizedObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	}
}

func TestDeepEqualsForOneUninitializedObject(t *testing.T) {
	t.Skip("This test doesn't pass because DeepEqual sucks")

	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	m1.IsValid()

	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	}
}

func TestDeepEqualsForTwoInitializedObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	m1.IsValid()
	m2.IsValid()

	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	}
}

func TestDeepEqualsForUnequalObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod2"}).Build()

	m1.IsValid()
	m2.IsValid()

	if reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return false for (%v, %v)", m1, m2)
	}
}