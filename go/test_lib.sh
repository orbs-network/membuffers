#!/bin/bash -e

echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 ./membufc/*.go | grep -v _test.go) -m `find ./e2e/types -name "*.proto"`

echo "  * Running tests"
echo ""

go test -count=1 .
go test -count=1 ./e2e 

echo "  /"
echo "\/"
echo "*** Tests for the Go library passed!"