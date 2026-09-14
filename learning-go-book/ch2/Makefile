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
	go build ch2.go
.PHONY:build

run: build
	./ch2

clean:
	go clean
