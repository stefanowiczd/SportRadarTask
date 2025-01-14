.PHONY: deps test run build lint test-race-cond

deps:
	go mod tidy

test:
	go test -tags unit -v ./...

test-race-cond:
	go test -race -tags race_cond -v ./...

run:
	go run .

build:
	go build -o main

lint:
	golangci-lint run -v

