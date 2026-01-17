
all: build

clean:
	rm -f *.so

lint:
	golangci-lint run

build:
	go mod tidy
	go build -buildmode=plugin -o fhir.so *.go

test: lint
	make -C fhir test

.PHONY: all build
