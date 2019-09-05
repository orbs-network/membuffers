package e2e

import (
	"github.com/orbs-network/membuffers/go/e2e/types"
	"testing"
)

func TestReadEmptyField_String(t *testing.T) {
	raw := (&types.WithStringBuilder{I: 42, S:"arthur"}).Build().Raw()
	read := types.WithStringReader(raw[:4])
	if read.I() != 42 {
		t.Fatalf("failed reading truncated raw")
	}
	if read.S() != "" {
		t.Fatalf("missing string field should've default to empty string")
	}
}

func TestReadEmptyField_Int(t *testing.T) {
	raw := (&types.WithIntBuilder{I: 42, J:17}).Build().Raw()
	read := types.WithIntReader(raw[:4])
	if read.I() != 42 {
		t.Fatalf("failed reading truncated raw")
	}
	if read.J() != 0 {
		t.Fatalf("missing int field should've default to zero")
	}
}

