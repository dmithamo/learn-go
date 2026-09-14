.DEFAULT_GOAL := build

fmt:
	go fmt ./...
.PHONY:fmt

lint: fmt
	golangci-lint run
.PHONY:lint

vet: fmt lint
	go vet ./...
.PHONY:vet

build: vet
	go build hello.go
.PHONY:build

run: build
	./hello

clean:
	go clean
