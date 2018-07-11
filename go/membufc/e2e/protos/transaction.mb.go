// AUTO GENERATED FILE (by membufc proto compiler v0.0.14)
package types

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	// Data TransactionData
	// Signature []byte
	// Type NetworkType

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _Transaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,membuffers.TypeUint16,}
var _Transaction_Unions = [][]membuffers.FieldType{}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Transaction_Scheme, _Transaction_Unions)
	return x
}

func (x *Transaction) IsValid() bool {
	return x._message.IsValid()
}

func (x *Transaction) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Transaction) Data() *TransactionData {
	b, s := x._message.GetMessage(0)
	return TransactionDataReader(b[:s])
}

func (x *Transaction) RawData() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Transaction) Signature() []byte {
	return x._message.GetBytes(1)
}

func (x *Transaction) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateSignature(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *Transaction) Type() NetworkType {
	return NetworkType(x._message.GetUint16(2))
}

func (x *Transaction) RawType() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Transaction) MutateType(v NetworkType) error {
	return x._message.SetUint16(2, uint16(v))
}

// builder

type TransactionBuilder struct {
	Data *TransactionDataBuilder
	Signature []byte
	Type NetworkType

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.Data)
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, w.Signature)
	w._builder.WriteUint16(buf, uint16(w.Type))
	return nil
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionBuilder) Build() *Transaction {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionData

// reader

type TransactionData struct {
	// ProtocolVersion uint32
	// VirtualChain uint64
	// Sender []TransactionSender
	// TimeStamp uint64

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionData_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessageArray,membuffers.TypeUint64,}
var _TransactionData_Unions = [][]membuffers.FieldType{}

func TransactionDataReader(buf []byte) *TransactionData {
	x := &TransactionData{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionData_Scheme, _TransactionData_Unions)
	return x
}

func (x *TransactionData) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionData) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionData) ProtocolVersion() uint32 {
	return x._message.GetUint32(0)
}

func (x *TransactionData) RawProtocolVersion() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionData) MutateProtocolVersion(v uint32) error {
	return x._message.SetUint32(0, v)
}

func (x *TransactionData) VirtualChain() uint64 {
	return x._message.GetUint64(1)
}

func (x *TransactionData) RawVirtualChain() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *TransactionData) MutateVirtualChain(v uint64) error {
	return x._message.SetUint64(1, v)
}

func (x *TransactionData) SenderIterator() *TransactionDataSenderIterator {
	return &TransactionDataSenderIterator{iterator: x._message.GetMessageArrayIterator(2)}
}

type TransactionDataSenderIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionDataSenderIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionDataSenderIterator) NextSender() *TransactionSender {
	b, s := i.iterator.NextMessage()
	return TransactionSenderReader(b[:s])
}

func (x *TransactionData) RawSenderArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *TransactionData) TimeStamp() uint64 {
	return x._message.GetUint64(3)
}

func (x *TransactionData) RawTimeStamp() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *TransactionData) MutateTimeStamp(v uint64) error {
	return x._message.SetUint64(3, v)
}

// builder

type TransactionDataBuilder struct {
	ProtocolVersion uint32
	VirtualChain uint64
	Sender []*TransactionSenderBuilder
	TimeStamp uint64

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionDataBuilder) arrayOfSender() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Sender))
	for i, v := range w.Sender {
		res[i] = v
	}
	return res
}

func (w *TransactionDataBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteUint32(buf, w.ProtocolVersion)
	w._builder.WriteUint64(buf, w.VirtualChain)
	err = w._builder.WriteMessageArray(buf, w.arrayOfSender())
	if err != nil {
		return
	}
	w._builder.WriteUint64(buf, w.TimeStamp)
	return nil
}

func (w *TransactionDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionDataBuilder) Build() *TransactionData {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionDataReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSender

// reader

type TransactionSender struct {
	// Name string
	// Friend []string

	// internal
	membuffers.Message // interface
	_message membuffers.InternalMessage
}

var _TransactionSender_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeStringArray,}
var _TransactionSender_Unions = [][]membuffers.FieldType{}

func TransactionSenderReader(buf []byte) *TransactionSender {
	x := &TransactionSender{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _TransactionSender_Scheme, _TransactionSender_Unions)
	return x
}

func (x *TransactionSender) IsValid() bool {
	return x._message.IsValid()
}

func (x *TransactionSender) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *TransactionSender) Name() string {
	return x._message.GetString(0)
}

func (x *TransactionSender) RawName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *TransactionSender) MutateName(v string) error {
	return x._message.SetString(0, v)
}

func (x *TransactionSender) FriendIterator() *TransactionSenderFriendIterator {
	return &TransactionSenderFriendIterator{iterator: x._message.GetStringArrayIterator(1)}
}

type TransactionSenderFriendIterator struct {
	iterator *membuffers.Iterator
}

func (i *TransactionSenderFriendIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *TransactionSenderFriendIterator) NextFriend() string {
	return i.iterator.NextString()
}

func (x *TransactionSender) RawFriendArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

// builder

type TransactionSenderBuilder struct {
	Name string
	Friend []string

	// internal
	membuffers.Builder // interface
	_builder membuffers.InternalBuilder
}

func (w *TransactionSenderBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.WriteString(buf, w.Name)
	w._builder.WriteStringArray(buf, w.Friend)
	return nil
}

func (w *TransactionSenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *TransactionSenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *TransactionSenderBuilder) Build() *TransactionSender {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return TransactionSenderReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

type NetworkType uint16

const (
	NETWORK_TYPE_MAIN_NET NetworkType = 0
	NETWORK_TYPE_TEST_NET NetworkType = 1
	NETWORK_TYPE_RESERVED NetworkType = 2
)

