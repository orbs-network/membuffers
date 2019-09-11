#!/bin/bash -e

echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 ../go/membufc/*.go | grep -v _test.go) -m `find ./types -name "*.proto"`

echo "  * Running tests"
echo ""

go test -count=1 .
