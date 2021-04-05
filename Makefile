
.PHONY: clean docker run

.DEFAULT_GOAL := all

build-import:
	go build -o dist/import cmd/import/main.go

build-find:
	go build -o dist/find cmd/find/main.go

build-server:
	go build -o dist/server cmd/server/main.go

all: build-import build-find build-server

docker:
	docker build -t nsr-parser .

run:
	docker run -it --rm --env ARANGODB_HOST=http://arangodb:8529 \
		--env ARANGODB_USER=root --env ARANGODB_PASSWORD=abc123 \
		nsr-parser

clean:
	rm -rf dist
