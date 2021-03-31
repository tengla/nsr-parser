
.PHONY: clean

build:
	go build -o dist/parser .

clean:
	rm -rf dist
