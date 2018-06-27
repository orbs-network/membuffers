go run *.go ./e2e/protos/transaction.proto
go run *.go ./e2e/protos/method.proto
go run *.go ./e2e/protos/dep1/dependency1.proto
go run *.go ./e2e/protos/dep1/dep11/dependency11.proto
go run *.go ./e2e/protos/dep2/dependent.proto
go test ./e2e

rm  `find . -name "*.mb.go"`

packr install
membufc --go ./e2e/protos/*.proto
membufc --go ./e2e/protos/dep1/*.proto
membufc --go ./e2e/protos/dep1/dep11/*.proto
membufc --go ./e2e/protos/dep2/*.proto
go test ./e2e