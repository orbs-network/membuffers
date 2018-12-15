package e2e

import (
	"bytes"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"reflect"
	"testing"
)

func TestWriteTransactionWithBuilderFromRawBuffer(t *testing.T) {
	// do a regular build
	builder1 := &types.TransactionBuilder{
		Data: &types.TransactionDataBuilder{
			ProtocolVersion: 0x01,
			VirtualChain:    0x11223344,
			Sender: []*types.TransactionSenderBuilder{
				&types.TransactionSenderBuilder{
					Name:   "johnny",
					Friend: []string{"billy", "jeff", "alex"},
				},
				&types.TransactionSenderBuilder{
					Name:   "rachel",
					Friend: []string{"jessica", "sara"},
				},
			},
			TimeStamp: 0x445566778899,
		},
		Signature: []byte{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22},
		Type:      types.NETWORK_TYPE_RESERVED,
	}
	txFromRegularBuild := builder1.Build()

	// do a build with transaction data from raw buffer
	builder2 := &types.TransactionBuilder{
		Data:      types.TransactionDataBuilderFromRaw(txFromRegularBuild.Data().Raw()),
		Signature: []byte{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22},
		Type:      types.NETWORK_TYPE_RESERVED,
	}
	txFromRawBufferBuild := builder2.Build()

	// compare
	if !bytes.Equal(txFromRegularBuild.Raw(), txFromRawBufferBuild.Raw()) {
		t.Fatalf("The two builds returned different results: (regular first, from raw second)\n%x\n%x\n", txFromRegularBuild.Raw(), txFromRawBufferBuild.Raw())
	}
	if reflect.ValueOf(txFromRegularBuild.Raw()).Pointer() == reflect.ValueOf(txFromRawBufferBuild.Raw()).Pointer() {
		t.Fatalf("The two slices are the same slice: (regular first, from raw second)\n%p\n%p\n", txFromRegularBuild.Raw(), txFromRawBufferBuild.Raw())
	}
}
