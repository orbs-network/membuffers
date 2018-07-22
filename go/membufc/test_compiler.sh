echo "  * Running unit tests"
echo ""

go test .
echo ""

echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 *.go | grep -v _test.go) -m `find . -name "*.proto"`

echo "  * Running end to end tests"
echo ""

go test ./e2e
echo ""

rm `find . -name "*.mb.go"`

echo "  * Building protos for tests (with building compiler)"
echo ""

packr install
`go env GOPATH`/bin/membufc --go --mock `find . -name "*.proto"`

echo "  * Running end to end tests"
echo ""

go test ./e2e
echo ""