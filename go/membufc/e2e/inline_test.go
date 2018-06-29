package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"bytes"
)

func TestReadWriteWithInline(t *testing.T) {
	// write
	builder := &types.FileRecordBuilder{
		Data: []byte{0x01,0x02,0x03},
		Hash: []byte{0x04,0x05,0x06},
		AnotherHash: [][]byte{{0x07,0x08},{0x09}},
	}
	buf := make([]byte, builder.CalcRequiredSize())
	err := builder.Write(buf)
	if err != nil {
		t.Fatalf("error while writing in builder")
	}

	// read
	record := types.FileRecordReader(buf)
	if !bytes.Equal(record.Hash(), []byte{0x04,0x05,0x06}) {
		t.Fatalf("FileRecord.Hash is not as expected")
	}
}