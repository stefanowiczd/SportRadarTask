.PHONY: deps test run build lint

deps:
	go mod tidy

test:
	go test -tags unit -v ./...

run:
	go run .

build:
	go build -o main

lint:
	golangci-lint run -v

