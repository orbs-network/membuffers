go run ./membufc/*.go -m `find ./e2e/types -name "*.proto"`

go test
go test ./e2e