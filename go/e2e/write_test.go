package e2e

import (
	"testing"
	"bytes"
)

func createNewTransaction() *TransactionWriter {
	newTransaction := &TransactionWriter{
		Data: &TransactionDataWriter{
			ProtocolVersion: 0x01,
			VirtualChain: 0x11223344,
			Sender: []*TransactionSenderWriter{
				&TransactionSenderWriter{
					Name: "johnny",
					Friend: []string{"billy","jeff","alex"},
				},
				&TransactionSenderWriter{
					Name: "rachel",
					Friend: []string{"jessica","sara"},
				},
			},
			TimeStamp: 0x445566778899,
		},
		Signature: []byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,},
	}
	return newTransaction
}

func TestWriteTransaction(t *testing.T) {
	// write
	newTransaction := createNewTransaction()
	newTransaction.Write(nil)
	buf := make([]byte, newTransaction.GetSize())
	newTransaction.Write(buf)
	// read
	transaction := ReadTransaction(buf)
	if transaction.Data().ProtocolVersion() != 0x01 {
		t.Fatalf("ProtocolVersion: instead of expected got %v", transaction.Data().ProtocolVersion())
	}
	if transaction.Data().VirtualChain() != 0x11223344 {
		t.Fatalf("VirtualChain: instead of expected got %v", transaction.Data().VirtualChain())
	}
	if transaction.Data().TimeStamp() != 0x445566778899 {
		t.Fatalf("TimeStamp: instead of expected got %v", transaction.Data().TimeStamp())
	}
	if !bytes.Equal(transaction.Signature(), []byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,}) {
		t.Fatalf("Signature: instead of expected got %v", transaction.Signature())
	}
	senderCount := 0
	for i := transaction.Data().SenderIterator(); i.HasNext(); {
		sender := i.NextSender()
		if senderCount == 0 {
			if sender.Name() != "johnny" {
				t.Fatalf("Sender0.Name: instead of expected got %v", sender.Name())
			}
			j := sender.FriendIterator()
			if j.NextFriend() != "billy" {
				t.Fatalf("Sender0.Friend0: did not get as expected")
			}
			if j.NextFriend() != "jeff" {
				t.Fatalf("Sender0.Friend1: did not get as expected")
			}
			if j.NextFriend() != "alex" {
				t.Fatalf("Sender0.Friend2: did not get as expected")
			}
		}
		if senderCount == 1 {
			if sender.Name() != "rachel" {
				t.Fatalf("Sender1.Name: instead of expected got %v", sender.Name())
			}
			j := sender.FriendIterator()
			if j.NextFriend() != "jessica" {
				t.Fatalf("Sender1.Friend0: did not get as expected")
			}
			if j.NextFriend() != "sara" {
				t.Fatalf("Sender1.Friend1: did not get as expected")
			}
		}
		senderCount++
	}
}

func createNewMethod() *MethodWriter {
	newMethod := &MethodWriter{
		Name: "MyMethod",
		Arg: []*MethodCallArgumentWriter{
			&MethodCallArgumentWriter{
				Type: MethodCallArgument_Type_Num,
				Num: 0x17,
			},
			&MethodCallArgumentWriter{
				Type: MethodCallArgument_Type_Str,
				Str: "flower",
			},
			&MethodCallArgumentWriter{
				Type: MethodCallArgument_Type_Data,
				Data: []byte{0x01,0x02,0x03},
			},
		},
	}
	return newMethod
}

func TestWriteMethod(t *testing.T) {
	// write
	newMethod := createNewMethod()
	newMethod.Write(nil)
	buf := make([]byte, newMethod.GetSize())
	newMethod.Write(buf)
	// read
	method := ReadMethod(buf)
	if method.Name() != "MyMethod" {
		t.Fatalf("Name: instead of expected got %v", method.Name())
	}
	i := method.ArgIterator()
	arg0 := i.NextArg()
	if !arg0.IsType_Num() {
		t.Fatalf("Arg0: is type is not num")
	}
	if arg0.IsType_Str() {
		t.Fatalf("Arg0: is type is str")
	}
	if arg0.Type() != MethodCallArgument_Type_Num {
		t.Fatalf("Arg0: type is not num")
	}
	if arg0.Type_Num() != 0x17 {
		t.Fatalf("Arg0.Num: instead of expected got %v", arg0.Type_Num())
	}
	arg1 := i.NextArg()
	if !arg1.IsType_Str() {
		t.Fatalf("Arg1: is type is not str")
	}
	if arg1.IsType_Num() {
		t.Fatalf("Arg1: is type is num")
	}
	if arg1.Type() != MethodCallArgument_Type_Str {
		t.Fatalf("Arg1: type is not str")
	}
	if arg1.Type_Str() != "flower" {
		t.Fatalf("Arg1.Str: instead of expected got %v", arg1.Type_Str())
	}
}