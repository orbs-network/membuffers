// AUTO GENERATED FILE (by membufc proto compiler v0.0.15)
package types

import (
	"github.com/orbs-network/membuffers/go"
	"fmt"
	"github.com/orbs-network/membuffers/go/membufc/e2e/protos/crypto"
)

/////////////////////////////////////////////////////////////////////////////
// message FileRecord

// reader

type FileRecord struct {
	// Data []byte
	// Hash crypto.Sha256
	// AnotherHash []crypto.Md5
	// BlockHeight crypto.BlockHeight

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *FileRecord) String() string {
	return fmt.Sprintf("{Data:%s,Hash:%s,AnotherHash:%s,BlockHeight:%s,}", x.StringData(), x.StringHash(), x.StringAnotherHash(), x.StringBlockHeight())
}

var _FileRecord_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytesArray,membuffers.TypeUint64,}
var _FileRecord_Unions = [][]membuffers.FieldType{}

func FileRecordReader(buf []byte) *FileRecord {
	x := &FileRecord{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _FileRecord_Scheme, _FileRecord_Unions)
	return x
}

func (x *FileRecord) IsValid() bool {
	return x._message.IsValid()
}

func (x *FileRecord) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *FileRecord) Data() []byte {
	return x._message.GetBytes(0)
}

func (x *FileRecord) RawData() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *FileRecord) MutateData(v []byte) error {
	return x._message.SetBytes(0, v)
}

func (x *FileRecord) StringData() string {
	return fmt.Sprintf("%x", x.Data())
}

func (x *FileRecord) Hash() crypto.Sha256 {
	return crypto.Sha256(x._message.GetBytes(1))
}

func (x *FileRecord) RawHash() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *FileRecord) MutateHash(v crypto.Sha256) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *FileRecord) StringHash() string {
	return fmt.Sprintf("%x", x.Hash())
}

func (x *FileRecord) AnotherHashIterator() *FileRecordAnotherHashIterator {
	return &FileRecordAnotherHashIterator{iterator: x._message.GetBytesArrayIterator(2)}
}

type FileRecordAnotherHashIterator struct {
	iterator *membuffers.Iterator
}

func (i *FileRecordAnotherHashIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *FileRecordAnotherHashIterator) NextAnotherHash() crypto.Md5 {
	return crypto.Md5(i.iterator.NextBytes())
}

func (x *FileRecord) RawAnotherHashArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *FileRecord) StringAnotherHash() (res string) {
	res = "["
	for i := x.AnotherHashIterator(); i.HasNext(); {
		res += fmt.Sprintf("%x", i.NextAnotherHash()) + ","
	}
	res += "]"
	return
}

func (x *FileRecord) BlockHeight() crypto.BlockHeight {
	return crypto.BlockHeight(x._message.GetUint64(3))
}

func (x *FileRecord) RawBlockHeight() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *FileRecord) MutateBlockHeight(v crypto.BlockHeight) error {
	return x._message.SetUint64(3, uint64(v))
}

func (x *FileRecord) StringBlockHeight() string {
	return fmt.Sprintf("%x", x.BlockHeight())
}

// builder

type FileRecordBuilder struct {
	Data []byte
	Hash crypto.Sha256
	AnotherHash []crypto.Md5
	BlockHeight crypto.BlockHeight

	// internal
	// implements membuffers.Builder
	_builder membuffers.InternalBuilder
}

func (w *FileRecordBuilder) arrayOfAnotherHash() [][]byte {
	res := make([][]byte, len(w.AnotherHash))
	for i, v := range w.AnotherHash {
		res[i] = v
	}
	return res
}

func (w *FileRecordBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteBytes(buf, w.Data)
	w._builder.WriteBytes(buf, []byte(w.Hash))
	w._builder.WriteBytesArray(buf, w.arrayOfAnotherHash())
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	return nil
}

func (w *FileRecordBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *FileRecordBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *FileRecordBuilder) Build() *FileRecord {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return FileRecordReader(buf)
}

