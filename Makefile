
.PHONY: clean docker run

.DEFAULT_GOAL := all

build-import:
	go build -o dist/import cmd/import/main.go

build-server:
	go build -o dist/server cmd/server/server.go cmd/server/main.go

all: build-import build-server

test:
	go test -cover ./...

docker:
	docker build -t nsr-parser .

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative trainroute/protos/trainroute.proto

clean:
	rm -rf dist
