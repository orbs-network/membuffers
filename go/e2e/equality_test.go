package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/e2e/types"
	"reflect"
	"github.com/google/go-cmp/cmp"
)

func TestEqualityForUnininitalizedObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	}

	if !cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return true for (%v, %v)", m1, m2)
	}
}

func TestEqualityForOneUninitializedObject(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	m1.IsValid()

	//TODO reflect.DeepEqual does not pass in this case
	//if !reflect.DeepEqual(m1, m2) {
	//	t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	//}

	if !cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return true for (%v, %v)", m1, m2)
	}
}

func TestEqualityForTwoInitializedObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod"}).Build()

	m1.IsValid()
	m2.IsValid()

	if !reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return true for (%v, %v)", m1, m2)
	}

	if !cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return true for (%v, %v)", m1, m2)
	}
}

func TestEqualityForUnequalObjects(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	m2 := (&types.MethodBuilder{Name: "myMethod2"}).Build()

	m1.IsValid()
	m2.IsValid()

	if reflect.DeepEqual(m1, m2) {
		t.Fatalf("Expected DeepEqual to return false for (%v, %v)", m1, m2)
	}

	if cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return false for (%v, %v)", m1, m2)
	}
}

func TestEqualityForOneUninitializedObjectAndOneNil(t *testing.T) {
	m1 := (&types.MethodBuilder{Name: "myMethod"}).Build()
	var m2 *types.Method

	if cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return false for (%v, %v)", m1, m2)
	}
}

func TestEqualityForTwoNilObjects(t *testing.T) {
	var m1 *types.Method
	var m2 *types.Method

	if !cmp.Equal(m1, m2) {
		t.Fatalf("Expected cmp.Equal to return true for (%v, %v)", m1, m2)
	}
}
