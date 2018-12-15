package e2e

import (
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"testing"
)

func TestHexDumpSender(t *testing.T) {
	builder := &types.TransactionSenderBuilder{
		Name:   "johnny",
		Friend: []string{"billy", "jeff", "alex"},
	}

	fmt.Println()
	fmt.Printf("*** Hex dump:\n\n")
	builder.HexDump("", 0)
	fmt.Println()

	fmt.Printf("*** Raw:\n\n")
	raw := builder.Build().Raw()
	membuffers.HexDumpRawInLines(raw, 0x20)
	fmt.Println()
}

func TestHexDumpTransaction(t *testing.T) {
	builder := &types.TransactionBuilder{
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

	fmt.Println()
	fmt.Printf("*** Hex dump:\n\n")
	builder.HexDump("", 0)
	fmt.Println()

	fmt.Printf("*** Raw:\n\n")
	raw := builder.Build().Raw()
	membuffers.HexDumpRawInLines(raw, 0x20)
	fmt.Println()
}

func TestHexDumpMethod(t *testing.T) {
	builder := &types.MethodBuilder{
		Name: "MyMethod",
		Arg: []*types.MethodCallArgumentBuilder{
			&types.MethodCallArgumentBuilder{
				Type: types.METHOD_CALL_ARGUMENT_TYPE_NUM,
				Num:  0x17,
			},
			&types.MethodCallArgumentBuilder{
				Type: types.METHOD_CALL_ARGUMENT_TYPE_STR,
				Str:  "flower",
			},
			&types.MethodCallArgumentBuilder{
				Type: types.METHOD_CALL_ARGUMENT_TYPE_DATA,
				Data: []byte{0x01, 0x02, 0x03},
			},
		},
	}

	fmt.Println()
	fmt.Printf("*** Hex dump:\n\n")
	builder.HexDump("", 0)
	fmt.Println()

	fmt.Printf("*** Raw:\n\n")
	raw := builder.Build().Raw()
	membuffers.HexDumpRawInLines(raw, 0x20)
	fmt.Println()
}
