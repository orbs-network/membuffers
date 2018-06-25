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

var m_Transaction_Scheme = []membuffers.FieldType{ membuffers.TypeMessage,membuffers.TypeBytes, }
var m_Transaction_Unions = [][]membuffers.FieldType{ {} }

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

/////////////////////////////////////////////////////////////////////////////
// message TransactionData

// reader

type TransactionData struct {
	message membuffers.Message
}

var m_TransactionData_Scheme = []membuffers.FieldType{ membuffers.TypeUint32,membuffers.TypeUint64,membuffers.TypeMessageArray,membuffers.TypeUint64, }
var m_TransactionData_Unions = [][]membuffers.FieldType{ {} }

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

func (x *TransactionData) TimeStamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *TransactionData) RawTimeStamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *TransactionData) MutateTimeStamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

/////////////////////////////////////////////////////////////////////////////
// message TransactionSender

// reader

type TransactionSender struct {
	message membuffers.Message
}

var m_TransactionSender_Scheme = []membuffers.FieldType{ membuffers.TypeString,membuffers.TypeStringArray, }
var m_TransactionSender_Unions = [][]membuffers.FieldType{ {} }

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

