# membufc Compiler

Build instructions are relevant only if you're trying to build the `membufc` compiler from source.

If you don't want to build from source, install the compiler with `brew install orbs-network/membuffers/membufc`.

## Building

#### Prerequisites

1. Make sure [Go](https://golang.org/doc/install) is installed (version 1.10 or later).
  
    > Verify with `go version`

2. Make sure [Go workspace bin](https://stackoverflow.com/questions/42965673/cant-run-go-bin-in-terminal) is in your path.
  
    > Install with ``export PATH=$PATH:`go env GOPATH`/bin``
  
    > Verify with `echo $PATH`

3. Get [packr](https://github.com/gobuffalo/packr).

    > Install with `go get -u github.com/gobuffalo/packr/...`

    > Verify with `packr --help`

#### Get and build

1. Get the library into your Go workspace:
 
     ```sh
     go get github.com/orbs-network/membuffers/go
     cd `go env GOPATH`/src/github.com/orbs-network/membuffers/go
     ```

* Build and install the `membufc` compiler:

    ```sh
    cd ./membufc
    packr install
    ```
    > Verify with `membufc --version`

&nbsp;
## Releasing

The compiler executable is released on Homebrew. To release a new version follow these steps:

1. Push to git a new tag with the version number (eg. `1.0.1`)

    1. Usually the commit will have the version number changed under const `MEMBUFC_VERSION` in the compiler code in:
    
        ```
        https://github.com/orbs-network/membuffers/blob/master/go/membufc/main.go
        ```

2. Edit the file `https://github.com/orbs-network/homebrew-membuffers/blob/master/membufc.rb`

    1. Change the source URL (line 6) to the correct version number, eg:
  
        ```
        https://github.com/orbs-network/membuffers/archive/1.0.1.tar.gz
        ```
  
    2. Calc SHA256 over the source code gzip, eg:
    
        ```
        wget -nv -O- https://github.com/orbs-network/membuffers/archive/1.0.1.tar.gz | openssl sha256
        ```
        
    3. Change the SHA256 (line 7) to what you calculated
    
    4. Change the version in the test (line 32) to the correct one, eg:
    
        ```
        assert_match "membufc 1.0.1", shell_output("#{bin}/membufc --version 2>&1", 2)
        ```

3. Make sure the edited file is pushed to git   

&nbsp;
## Extensions to Protobuf schema with options

The `membufc` compiler supports several useful extensions to the standard [Protobuf v3 schema](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec) by utilizing `option` fields.

### Inline types (aliases)

Inline types are new names that behave as aliases to standard system types. You can view them as `messages` with a single field which are inlined whenever appearing in a different `message`.

Consider the following example which aliases the type `sha256` for type `bytes`:

##### crypto/aliases.proto:
```proto
syntax = "proto3";
package crypto;

// NOTE: inline files must be in packages having only inline files

option inline = true;

message sha256 {
    option inline_type = "bytes";
    bytes value = 1;
}
```

##### file_record.proto:
```proto
syntax = "proto3";
package files;

import "crypto/aliases.proto";

message FileRecord {
    bytes data = 1;
    crypto.sha256 hash = 2;
}
```

You can see a working example of this feature in the compiler [test suite](e2e/inline_test.go) and [test protos](e2e/protos/aliases_user.proto).

### Service listener pattern

Circular dependencies between services are often resolved with a listener pattern where one of the services extracts its callback methods into a separate service and the other service exposes a registration method for the listener.

Consider these two services:

##### notifier.proto:
```proto
syntax = "proto3";
package notifier;

service Notifier {
    rpc SendNotification (SNInput) returns (SNOutput);
}
```

##### consumer.proto:
```proto
syntax = "proto3";
package consumer;

service Consumer {
    rpc NotificationReceived (NRInput) returns (NROutput);
    rpc AnotherMethod (AMInput) returns (AMOutput);
}
```

The service Consumer uses service Notifier to send a notification by calling `Notifier.SendNotification`, but it may also receive a notification back. This works by service Notifier calling `Consumer.NotificationReceived`.

This is a circular dependency between the services, that we may want to break. One of the common solutions is extracting the callback into a new listener interface:

##### notifier.proto:
```proto
syntax = "proto3";
package notifier;

service Notifier {
    // RegisterListener (NotificationListener) returns ();
    rpc SendNotification (SNInput) returns (SNOutput);
}
```

##### notification_listener.proto:
```proto
syntax = "proto3";
package listener;

service NotificationListener {
    rpc NotificationReceived (NRInput) returns (NROutput);
}
```

##### consumer.proto:
```proto
syntax = "proto3";
package consumer;

service Consumer {
    // implements NotificationListener
    rpc AnotherMethod (AMInput) returns (AMOutput);
}
```

Now, service Consumer relies on service Notifier (we broke the other direction) and calls `Notifier.SendNotification`. It also registers itself as a callback listener by calling `Notifier.RegisterListener(self)`. It has to implement the new interface `NotificationListener` which decouples the services.

This pattern can be implemented using `option` schema extensions this way:

##### notifier.proto:
```proto
syntax = "proto3";
package notifier;

service Notifier {
   option register_handler = "listener.NotificationListener";
   rpc SendNotification (SNInput) returns (SNOutput);
}
```
 
##### notification_listener.proto:
```proto
syntax = "proto3";
package listener;

service NotificationListener {
   rpc NotificationReceived (NRInput) returns (NROutput);
}
```
 
##### consumer.proto:
```proto
syntax = "proto3";
package consumer;
 
service Consumer {
    option implement_handler = "listener.NotificationListener";
    rpc AnotherMethod (AMInput) returns (AMOutput);
}
```
 
You can see a working example of this feature in the compiler [test suite](e2e/handlers_test.go) and [test protos](e2e/protos/options/handlers.proto).

### Services with non serializable arguments

Wrapping an already encoded MemBuffers message with another MemBuffers message causes data copy. This is particularly taxing with argument wrappers for service methods which can be avoided by encoding them as plain structs instead of MemBuffers messages.

Consider this service:

```proto
service StateStorage {
    rpc WriteKey (WriteKeyInput) returns (WriteKeyOutput);
}

message WriteKeyInput {
    string key = 1;
    uint32 value = 2;
}

message WriteKeyOutput {
    uint32 result = 1;
}
```

The messages `WriteKeyInput` and `WriteKeyOutput` by default will also be MemBuffers messages which can be serialized quickly to a data stream.

Normally, this does not add much overhead but if you think about it, isn't really needed since they're just argument wrappers. It may add overhead if one of the messages contained another MemBuffers message which was already encoded. In this case, its data would need to be copied in order to fit inside the wrapper MemBuffers.

By adding the following option, all messages in the proto file will be encoded as regular structs instead of MemBuffers (will not be serializable):

```proto
option serialize_messages = false;

service StateStorage {
    rpc WriteKey (WriteKeyInput) returns (WriteKeyOutput);
}

message WriteKeyInput {
    string key = 1;
    uint32 value = 2;
}

message WriteKeyOutput {
    uint32 result = 1;
}
```

You can see a working example of this feature in the compiler [test suite](e2e/service_test.go) and [test protos](e2e/protos/service_no_serialization.proto).

### Non serializable messages

The schema supports specifying that any message should be compiled to a plain struct instead of a MemBuffer. This is useful for messages that have no need in being serialized.

Consider this example:

```proto
message ExampleContainer {
    option serialize_message = false;
    MessageInContainer message1 = 1;
    NestedContainer container1 = 2;
    repeated NestedContainer containers2 = 3;
}

message MessageInContainer {
    string field = 1;
}

message NestedContainer {
    option serialize_message = false;
    string name = 1;
}
```

The messages `ExampleContainer` and `NestedContainer` will be compiled to plain structs whereas `MessageInContainer` will be compiled to a regular MemBuffer which supports serialization.