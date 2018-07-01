// AUTO GENERATED FILE (by membufc proto compiler)
package types

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message FileRecord

// reader

type FileRecord struct {
	message membuffers.Message
}

var m_FileRecord_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeBytes,membuffers.TypeBytesArray,}
var m_FileRecord_Unions = [][]membuffers.FieldType{}

func FileRecordReader(buf []byte) *FileRecord {
	x := &FileRecord{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_FileRecord_Scheme, m_FileRecord_Unions)
	return x
}

func (x *FileRecord) IsValid() bool {
	return x.message.IsValid()
}

func (x *FileRecord) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *FileRecord) Data() []byte {
	return x.message.GetBytes(0)
}

func (x *FileRecord) RawData() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *FileRecord) MutateData(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *FileRecord) Hash() []byte {
	return x.message.GetBytes(1)
}

func (x *FileRecord) RawHash() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *FileRecord) MutateHash(v []byte) error {
	return x.message.SetBytes(1, v)
}

func (x *FileRecord) AnotherHashIterator() *FileRecordAnotherHashIterator {
	return &FileRecordAnotherHashIterator{iterator: x.message.GetBytesArrayIterator(2)}
}

type FileRecordAnotherHashIterator struct {
	iterator *membuffers.Iterator
}

func (i *FileRecordAnotherHashIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *FileRecordAnotherHashIterator) NextAnotherHash() []byte {
	return i.iterator.NextBytes()
}

func (x *FileRecord) RawAnotherHashArray() []byte {
	return x.message.RawBufferForField(2, 0)
}

// builder

type FileRecordBuilder struct {
	builder membuffers.Builder
	Data []byte
	Hash []byte
	AnotherHash [][]byte
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
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Data)
	w.builder.WriteBytes(buf, w.Hash)
	w.builder.WriteBytesArray(buf, w.AnotherHash)
	return nil
}

func (w *FileRecordBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *FileRecordBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *FileRecordBuilder) Build() *FileRecord {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return FileRecordReader(buf)
}

