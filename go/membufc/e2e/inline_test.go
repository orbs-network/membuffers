package e2e

import (
	"testing"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos"
	"bytes"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/crypto"
)

func TestReadWriteWithInline(t *testing.T) {
	// write
	builder := &types.FileRecordBuilder{
		Data: []byte{0x01,0x02,0x03},
		Hash: crypto.Sha256{0x04,0x05,0x06},
		AnotherHash: []crypto.Md5{{0x07,0x08},{0x09}},
	}
	buf := make([]byte, builder.CalcRequiredSize())
	err := builder.Write(buf)
	if err != nil {
		t.Fatalf("error while writing in builder")
	}

	// read
	record := types.FileRecordReader(buf)
	if !bytes.Equal(record.Hash(), crypto.Sha256{0x04,0x05,0x06}) {
		t.Fatalf("FileRecord.Hash is not as expected")
	}
	iter := record.AnotherHashIterator()
	if !bytes.Equal(iter.NextAnotherHash(), crypto.Md5{0x07,0x08}) {
		t.Fatalf("FileRecord.AnotherHash[0] is not as expected")
	}
	if !bytes.Equal(iter.NextAnotherHash(), crypto.Md5{0x09}) {
		t.Fatalf("FileRecord.AnotherHash[0] is not as expected")
	}
}