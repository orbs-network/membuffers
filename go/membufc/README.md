# membufc Compiler

There instructions are relevant only to people trying to build the `membufc` compiler from source.

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
    