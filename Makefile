
all: build

clean:
	rm -f *.so

lint:
	golangci-lint run

build:
	go mod tidy
	go build -buildmode=plugin -o fhir.so *.go
	make -C cmd

test: lint
	make -C fhir test

.PHONY: all build
