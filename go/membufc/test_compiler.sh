echo "  * Running unit tests"
echo ""

go test .
echo ""

echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 *.go | grep -v _test.go) --go --mock `find . -name "*.proto"`
go run $(ls -1 *.go | grep -v _test.go) --go --mock --go-ctx `find . -name "*_with_ctx.proto"`

echo "  * Running end to end tests"
echo ""

go test ./e2e
echo ""

rm `find . -name "*.mb.go"`

echo "  * Building protos for tests (with building compiler)"
echo ""

packr install
`go env GOPATH`/bin/membufc --go --mock `find . -name "*.proto"`
`go env GOPATH`/bin/membufc --go --mock --go-ctx `find . -name "*_with_ctx.proto"`

echo "  * Running end to end tests"
echo ""

go test ./e2e
echo ""