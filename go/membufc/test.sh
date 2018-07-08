go run *.go ./e2e/protos/transaction.proto
go run *.go ./e2e/protos/method.proto
go run *.go ./e2e/protos/dep1/dependency1.proto
go run *.go ./e2e/protos/dep1/dep11/dependency11.proto
go run *.go ./e2e/protos/dep2/dependent.proto
go run *.go -m ./e2e/protos/service.proto
go run *.go -m ./e2e/protos/service_no_serialization.proto
go run *.go ./e2e/protos/crypto/aliases.proto
go run *.go ./e2e/protos/aliases_user.proto
go run *.go ./e2e/protos/unions.proto
go run *.go -m ./e2e/protos/options/handlers/handlers.proto
go run *.go -m ./e2e/protos/options/handlers.proto
go test ./e2e

rm  `find . -name "*.mb.go"`

packr install
membufc --go --mock ./e2e/protos/*.proto
membufc --go ./e2e/protos/dep1/*.proto
membufc --go ./e2e/protos/dep1/dep11/*.proto
membufc --go ./e2e/protos/dep2/*.proto
membufc --go ./e2e/protos/crypto/*.proto
membufc --go --mock ./e2e/protos/options/handlers/*.proto
membufc --go --mock ./e2e/protos/options/*.proto
go test ./e2e