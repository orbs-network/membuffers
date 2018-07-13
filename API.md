# MemBuffers API

## Entities available from generated code

* **MemBuffer** (eg. `tx *Transaction`)
  
  Serialized object. Its underlying data is sequential on a single byte array. The MemBuffer is actually a thin accessor layer above the byte array that allows to access fields directly on it using a convenient API. 

* **Raw** bytes (eg. `tx.Raw()`)

  The underlying byte array of a MemBuffer. When using `Raw` for a specific field, a relevant slice over this byte array is returned. 

* **Reader** (eg. `TransactionReader`)

  A function that takes a byte array and returns a MemBuffer. This action is very cheap because the data is not really decoded, only the accessor layer is created.

* **Builder** (eg. `TransactionBuilder`)

  An object that assists developers with creating a single MemBuffer from a collection of fields and pointers. This action is not cheap because they all must be copied and serialized together as a single byte array. 

* Plain struct (non serializable, "POJO")

  Some messages in the generated code are marked as non-serializable. This means they're not MemBuffers, but will create a plain old struct instead (that doesn't have any methods).

##### When should you use each of them?

Most of the time, an object will travel within the system as a **MemBuffer**. This MemBuffer can be transmitted over the wire with its **Raw** bytes and can be reconstructed from bytes on the receiving end using a **Reader**. Whenever the MemBuffer is first created in the system, it is created though a **Builder**. 

## Known limitations

* You can't use an existing **MemBuffer** inside a **Builder**.

  This is on purpose. If you get to a situation where you need to place an already serialized MemBuffer inside another new MemBuffer through a Builder, stop and think. This usually indicates some architecture problem because the entire purpose of MemBuffers is to avoid excessive copying. Placing an already serialized MemBuffer inside another MemBuffer must involve copying since the existing MemBuffer already has an underlying byte array that can't be increased in size.   