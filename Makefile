.PHONY: all build test lint gendocs clean

all: test build

build:
	go build -v ./cmd/gitverse

test:
	go test -v -race ./...
	node --test "internal/webui/jstest/*.test.mjs"

lint:
	gofmt -l .
	go vet ./...

gendocs:
	go run ./cmd/gendocs

clean:
	rm -rf dist/
