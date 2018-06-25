// AUTO GENERATED FILE (by membufc proto compiler)
package types

import (
	"github.com/orbs-network/membuffers/go"
)

/////////////////////////////////////////////////////////////////////////////
// message Transaction

// reader

type Transaction struct {
	message membuffers.Message
}

var m_Transaction_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeBytes,}
var m_Transaction_Unions = [][]membuffers.FieldType{{}}

func TransactionReader(buf []byte) *Transaction {
	x := &Transaction{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_Transaction_Scheme, m_Transaction_Unions)
	return x
}

func (x *Transaction) IsValid() bool {
	return x.message.IsValid()
}

func (x *Transaction) Raw() []byte {
	return x.message.RawBuffer()
}
func (x *Transaction) Data() *TransactionData {
	b, s := x.message.GetMessage(0)
	return TransactionDataReader(b[:s])
}

func (x *Transaction) RawData() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *Transaction) Signature() []byte {
	return x.message.GetBytes(1)
}

func (x *Transaction) RawSignature() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *Transaction) MutateSignature(v []byte) error {
	return x.message.SetBytes(1, v)
}

// builder

type TransactionBuilder struct {
	builder membuffers.Builder
	Data *TransactionDataBuilder
	Signature []byte
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
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Data)
	if err != nil {
		return
	}
	w.builder.WriteBytes(buf, w.Signature)
	return nil
}

func (w *TransactionBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionData

// reader

type TransactionData struct {
	message membuffers.Message
}

var m_TransactionData_Scheme = []membuffers.FieldType{membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessageArray,membuffers.TypeUint64,}
var m_TransactionData_Unions = [][]membuffers.FieldType{{}}

func TransactionDataReader(buf []byte) *TransactionData {
	x := &TransactionData{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionData_Scheme, m_TransactionData_Unions)
	return x
}

func (x *TransactionData) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionData) Raw() []byte {
	return x.message.RawBuffer()
}
func (x *TransactionData) ProtocolVersion() uint32 {
	return x.message.GetUint32(0)
}

func (x *TransactionData) RawProtocolVersion() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionData) MutateProtocolVersion(v uint32) error {
	return x.message.SetUint32(0, v)
}

func (x *TransactionData) VirtualChain() uint64 {
	return x.message.GetUint64(1)
}

func (x *TransactionData) RawVirtualChain() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *TransactionData) MutateVirtualChain(v uint64) error {
	return x.message.SetUint64(1, v)
}

func (x *TransactionData) SenderIterator() *TransactionDataSenderIterator {
	return &TransactionDataSenderIterator{iterator: x.message.GetMessageArrayIterator(2)}
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
	return x.message.RawBufferForField(2, 0)
}

func (x *TransactionData) TimeStamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *TransactionData) RawTimeStamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *TransactionData) MutateTimeStamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type TransactionDataBuilder struct {
	builder membuffers.Builder
	ProtocolVersion uint32
	VirtualChain uint64
	Sender []*TransactionSenderBuilder
	TimeStamp uint64
}

func (w *TransactionDataBuilder) arrayOfSender() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.Sender))
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
	w.builder.Reset()
	w.builder.WriteUint32(buf, w.ProtocolVersion)
	w.builder.WriteUint64(buf, w.VirtualChain)
	err = w.builder.WriteMessageArray(buf, w.arrayOfSender())
	if err != nil {
		return
	}
	w.builder.WriteUint64(buf, w.TimeStamp)
	return nil
}

func (w *TransactionDataBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionDataBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSender

// reader

type TransactionSender struct {
	message membuffers.Message
}

var m_TransactionSender_Scheme = []membuffers.FieldType{membuffers.TypeString,membuffers.TypeStringArray,}
var m_TransactionSender_Unions = [][]membuffers.FieldType{{}}

func TransactionSenderReader(buf []byte) *TransactionSender {
	x := &TransactionSender{}
	x.message.Init(buf, membuffers.Offset(len(buf)), m_TransactionSender_Scheme, m_TransactionSender_Unions)
	return x
}

func (x *TransactionSender) IsValid() bool {
	return x.message.IsValid()
}

func (x *TransactionSender) Raw() []byte {
	return x.message.RawBuffer()
}
func (x *TransactionSender) Name() string {
	return x.message.GetString(0)
}

func (x *TransactionSender) RawName() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *TransactionSender) MutateName(v string) error {
	return x.message.SetString(0, v)
}

func (x *TransactionSender) FriendIterator() *TransactionSenderFriendIterator {
	return &TransactionSenderFriendIterator{iterator: x.message.GetStringArrayIterator(1)}
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
	return x.message.RawBufferForField(1, 0)
}

// builder

type TransactionSenderBuilder struct {
	builder membuffers.Builder
	Name string
	Friend []string
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
	w.builder.Reset()
	w.builder.WriteString(buf, w.Name)
	w.builder.WriteStringArray(buf, w.Friend)
	return nil
}

func (w *TransactionSenderBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *TransactionSenderBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

