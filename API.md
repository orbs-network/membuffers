# MemBuffers API

## Entities available from generated code

* **MemBuffer** - eg. `tx *Transaction`
  
  Serialized object. Its underlying data is a single sequential byte array. The MemBuffers interface is actually a thin accessor layer above the byte array that allows to access fields directly on it using a convenient API. 

* **Raw** bytes - eg. `tx.Raw()`

  Direct access to the underlying byte array of a MemBuffer. When using `Raw()` for a specific field, a relevant slice over this byte array is returned. 

* **Reader** - eg. `TransactionReader`

  A function that takes a byte array and returns a MemBuffer. This action is very cheap because the data is not really decoded, only the thin accessor layer is created.

* **Builder** - eg. `TransactionBuilder`

  An object that assists developers with creating a single MemBuffer from a collection of fragmented fields and pointers (non sequential in memory). This action is not cheap because all fields must be copied and serialized together to a single byte array. 

* Plain struct - non serializable, "POJO"

  Some messages in the generated code are marked as non-serializable. This means they're not MemBuffers, but will create a plain old struct instead that doesn't have any interesting methods.

### When should you use each of them?

Most of the time, an object will travel within the system as a **MemBuffer**. This MemBuffer can be transmitted over the wire with its **Raw** bytes and can be reconstructed from raw bytes on the receiving end using a **Reader**. Whenever the MemBuffer is first created in the system, it is usually created though a **Builder**. 

## Examples of working with each entity

### Let's start with the schema

Consider this rather complicated example of a .proto file describing a transaction:

```proto
message Transaction {
    TransactionData data = 1;
    bytes signature = 2;
    NetworkType type = 3;
}

message TransactionData {
    uint32 protocol_version = 1;
    uint64 virtual_chain = 2;
    repeated TransactionSender sender = 3;
    uint64 time_stamp = 4;
}

message TransactionSender {
    string name = 1;
    repeated string friend = 2;
}

enum NetworkType {
    NETWORK_TYPE_MAIN_NET = 0;
    NETWORK_TYPE_TEST_NET = 1;
    NETWORK_TYPE_RESERVED = 2;
}
```

Working with unions (`oneof`) is also a bit tricky, so here's a different .proto file describing a method with one:

```proto
message Method {
    string name = 1;
    repeated MethodCallArgument arg = 2;
}

message MethodCallArgument {
    oneof type {
        uint32 num = 1;
        string str = 2;
        bytes data = 3;
    }
}
```

In the next sections we're going to attempt to encode and decode these two .protos separately - `Transaction` and `Method`.

### Working with a Builder

The code generator will create `TransactionBuilder`, this is the **Builder** to create a MemBuffer for the schema. Here's the idiomatic way to work with the builder:

```go
builder := &types.TransactionBuilder{
  Data: &types.TransactionDataBuilder{
    ProtocolVersion: 0x01,
    VirtualChain: 0x11223344,
    Sender: []*types.TransactionSenderBuilder{
      { Name: "johnny", Friend: []string{"billy","jeff","alex"} },
      { Name: "rachel", Friend: []string{"jessica","sara"} },
    },
    TimeStamp: 0x445566778899,
  },
  Signature: []byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,},
  Type: types.NETWORK_TYPE_RESERVED,
}
tx := buider.Build()
```

First example complete. A builder that is working with unions (`oneof`) is a little tricky. So here's a separate example of building the `Method` message from before:

```go
builder := &types.MethodBuilder{
  Name: "MyMethod",
  Arg: []*types.MethodCallArgumentBuilder{
    { Type: types.METHOD_CALL_ARGUMENT_TYPE_NUM, Num:  0x17 },
    { Type: types.METHOD_CALL_ARGUMENT_TYPE_STR, Str:  "flower" },
    { Type: types.METHOD_CALL_ARGUMENT_TYPE_DATA, Data: []byte{0x01,0x02,0x03} },
  },
}
method := builder.Build()
```

The example above shows a method that has 3 arguments, one of each type.

### Working with a MemBuffer

The `tx` object that we received from `buider.Build()` is a **MemBuffer** of type `*Transaction`. It is very easy to access its various fields:

```go
// read some fields
fmt.Printf("%d", tx.Data().ProtocolVersion()) // prints "1"
fmt.Printf("%x", tx.Data().VirtualChain()) // prints "11223344"
fmt.Printf("%x", tx.Data().TimeStamp()) // prints "445566778899"
if bytes.Equal(tx.Signature(), []byte{0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,0x22,}) {
  fmt.Println("This works")
}
if tx.Type() == types.NETWORK_TYPE_RESERVED {
  fmt.Println("This also works")
}

// reading an array uses the iterator pattern
for i := tx.Data().SenderIterator(); i.HasNext(); {
  sender := i.NextSender()
  fmt.Printf("%s", sender.Name()) // prints "johnny" on first iteration, "rachel" on the second
}

// we can also mutate and change the fields
tx.Data().MutateProtocolVersion(0x02)
fmt.Printf("%d", tx.Data().ProtocolVersion()) // this will now print "2"

// we can print the whole object to string (for logging), or just a single field
fmt.Printf("%v", tx) // prints all fields as a nice string
fmt.Printf("%s", tx.String()) // this does the same thing
fmt.Printf("%s", tx.Data().StringVirtualChain()) // prints "11223344"
```

Let's see how we would access the `method` MemBuffer object we received from building the `Method` message:

```go
fmt.Printf("%s", method.Name()) // prints "MyMethod"
i := method.ArgIterator()

// read the first argument
arg0 := i.NextArg()
if arg0.IsTypeNum() {
  fmt.Printf("%x", arg0.Num()) // prints "17"
}
if arg0.Type() == types.METHOD_CALL_ARGUMENT_TYPE_NUM {
  fmt.Println("This also works")
}

// read the second argument
arg1 := i.NextArg()
if arg1.IsTypeStr() {
  fmt.Printf("%s", arg1.Str()) // prints "flower"
}
if arg1.Type() == types.METHOD_CALL_ARGUMENT_TYPE_STR {
  fmt.Println("This also works")
}
``` 

### Working with raw bytes of a MemBuffer

Just use the `Raw()` method on the **MemBuffer**:

```go
var txBytes []byte = tx.Raw() // get the underlying byte array of the whole transaction
var virtualChainBytes []byte = tx.Data().RawVirtualChain() // get a slice for a single field
```

And an example with our `method` MemBuffer:

```go
var methodBytes []byte = method.Raw() // get the underlying byte array of the whole method
arg0 := method.ArgIterator().NextArg() // access the first argument (like before)
var arg0Bytes []byte = arg0.RawType() // returns []byte{0x17,0x00,0x00,0x00} 
```

### Working with a Reader

Let's assume we've sent the transaction bytes `txBytes` over the wire, this is how we parse them back:

```go
var txBytes []byte = ReadFromWire()
tx := TransactionReader(txBytes) // tx is now a regular MemBuffer
```

And the same exact thing with `methodBytes`:

```go
var methodBytes []byte = ReadFromWire()
method := MethodReader(methodBytes) // method is now a regular MemBuffer
```

## Creating a Builder from a MemBuffer / raw bytes

Assume we have a field that is already encoded as a MemBuffer. Now, we want to place it inside another wrapper MemBuffer. How do we create the new MemBuffer? What's the API for doing that?

Normally, to create a Builder, you must provide the values for all fields separately. Consider if you have a message inside a message, meaning a parent Builder that wraps a child Builder. If you have the data for the child Builder only in MemBuffer form, it can get annoying to pull every field out of the MemBuffer. 

This is on purpose. If you get to a situation where you need to place an already serialized MemBuffer inside another new MemBuffer through a Builder, stop and think. This usually indicates some design problem because the entire purpose of MemBuffers is to avoid excessive copying. Placing an already serialized MemBuffer inside another MemBuffer must involve copying since the existing MemBuffer already has an underlying byte array that can't be increased in size.
  
So what design changes can you make in this situation? There are usually 3 alternatives, choose the most appropriate:

  1. The wrapper MemBuffer should not be a MemBuffer at all. It should maybe be a plain old struct. If this is the case, mark it with `option serialize_message = false` in its .proto file.
  
  2. Use a Builder and for every field access the field from the first MemBuffer. This basically means you're rebuilding it from scratch. This makes sense if the wrapper MemBuffer may be changed to a different serialization format and just by accident it's also a MemBuffer.
  
  3. Mark the field in the wrapper MemBuffer as an opaque byte array. To do that, in the .proto file of the wrapper change the type of the internal MemBuffer to `bytes`. This means it's now opaque. When building the wrapper, provide the field bytes by calling `Raw()` on the first MemBuffer. If you ever want to access its fields from the wrapper MemBuffer, just use a reader. This approach usually makes sense when there's separation of concerns and some parts of the code should not be aware of the internal format of the field. 

If none of these work out and your mind is set on copying, this is indeed possible - although not recommended:

```go
tx1 := givenAsMemBuffer()

builder := &types.TransactionBuilder{
  Data:      types.TransactionDataBuilderFromRaw(tx1.Data().Raw()), // here we create a builder using a raw buffer (copying the data!)
  Signature: []byte{0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22},
  Type:      types.NETWORK_TYPE_RESERVED,
}
tx2 := builder.Build()
```