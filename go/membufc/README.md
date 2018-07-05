# `membufc` Compiler

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

## Extensions to Protobuf schema with options

The `membufc` compiler supports several useful extensions to the standard [Protobuf v3 schema](https://developers.google.com/protocol-buffers/docs/reference/proto3-spec) by utilizing `option` fields.

#### Inline types (aliases)

Inline types are new names that behave as aliases to standard system types. You can view them as `messages` with a single field which are inlined whenever appearing in a different `message`.

#### Service listener pattern

Circular dependencies between services are often resolved with a listener pattern where one of the services extracts its callback methods into a separate service and the other service exposes a registration method for the listener.