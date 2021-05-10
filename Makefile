
.PHONY: clean docker run

.DEFAULT_GOAL := all

build-import:
	go build -o dist/import cmd/import/main.go

build-find:
	go build -o dist/find cmd/find/main.go

build-server:
	go build -o dist/server cmd/server/main.go

build-grpc-server:
	go build -o dist/grpc-server cmd/grpc/server/main.go

build-grpc-client:
	go build -o dist/grpc-client cmd/grpc/client/main.go

all: build-import build-find build-server \
	build-grpc-server build-grpc-client

test:
	go test -cover ./...

docker:
	docker build -t nsr-parser .

run:
	docker run -it --rm --env ARANGODB_HOST=http://arangodb:8529 \
		--env ARANGODB_USER=root --env ARANGODB_PASSWORD=abc123 \
		nsr-parser

protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative trainroute/protos/trainroute.proto

clean:
	rm -rf dist
