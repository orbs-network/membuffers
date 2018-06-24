package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/e2e/types"
	"reflect"
)

func TestMutateFields(t *testing.T) {
	// write
	builder := &types.TransactionBuilder{
		Data: &types.TransactionDataBuilder{
			ProtocolVersion: 0x01,
			VirtualChain: 0x11223344,
			TimeStamp: 0x445566778899,
		},
		Signature: []byte{0x00,0x00,0x00,0x00,0x00,0x00,0x00},
	}
	buf := make([]byte, builder.CalcRequiredSize())
	err := builder.Write(buf)
	if err != nil {
		t.Fatalf("error while writing in builder")
	}

	// read
	transaction := types.TransactionReader(buf)
	transaction.Data().MutateProtocolVersion(0x22)
	if transaction.Data().ProtocolVersion() != 0x22 {
		t.Fatalf("ProtocolVersion: instead of expected got %v", transaction.Data().ProtocolVersion())
	}
	transaction.Data().MutateVirtualChain(0x33)
	if transaction.Data().VirtualChain() != 0x33 {
		t.Fatalf("VirtualChain: instead of expected got %v", transaction.Data().VirtualChain())
	}
	transaction.Data().MutateTimeStamp(0x44)
	if transaction.Data().TimeStamp() != 0x44 {
		t.Fatalf("TimeStamp: instead of expected got %v", transaction.Data().TimeStamp())
	}
	transaction.MutateSignature([]byte{0x55,0x55,0x55,0x55,0x55,0x55,0x55})
	if !reflect.DeepEqual(transaction.Signature(), []byte{0x55,0x55,0x55,0x55,0x55,0x55,0x55}) {
		t.Fatalf("Signature: instead of expected got %v", transaction.Signature())
	}
	err = transaction.MutateSignature([]byte{0x66,0x66,0x66})
	if !reflect.DeepEqual(transaction.Signature(), []byte{0x55,0x55,0x55,0x55,0x55,0x55,0x55}) {
		t.Fatalf("Signature: instead of expected got %v", transaction.Signature())
	}
	if err == nil {
		t.Fatalf("Signature: expected mutation to return error")
	}
}