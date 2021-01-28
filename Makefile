build-api:
	go build cmd/api

build-domain:
	go build cmd/domain

test:
	go test -race ./...

lint:
	golangci-lint run

security:
	gosec -conf gosec.json

generate:
	 protoc --go_out=. --go-grpc_out=. proto/ports.proto
