echo "  * Building protos for tests (without building compiler)"
echo ""

go run $(ls -1 ./membufc/*.go | grep -v _test.go) -m `find ./e2e/types -name "*.proto"`

echo "  * Running tests"
echo ""

go test
go test ./e2e