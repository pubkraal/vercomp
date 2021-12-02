.PHONY: all test

all: target/vercomp target/verfilter

test: lint
	go test -cover ./...

lint:
	golangci-lint run

target/vercomp: $(shell find cmd/vercomp -type f) $(shell find pkg -name \*.go)
	go build -o ./target/vercomp ./cmd/vercomp

target/verfilter: $(shell find cmd/verfilter -type f) $(shell find pkg -name \*.go)
	go build -o ./target/verfilter ./cmd/verfilter