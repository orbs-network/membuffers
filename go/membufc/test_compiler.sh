go test .

go run *.go -m `find . -name "*.proto"`
go test ./e2e

rm  `find . -name "*.mb.go"`

packr install
`go env GOPATH`/bin/membufc --go --mock `find . -name "*.proto"`
go test ./e2e