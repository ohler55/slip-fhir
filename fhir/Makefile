
all: test

lint:
	golangci-lint run

build:
	go build *.go

test: lint
	go test -coverprofile=cov.out

.PHONY: all build test
