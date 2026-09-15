.DEFAULT_GOAL := build

.PHONY: fmt vet build
fmt:
	go fmt ./...

vet: fmt 
	go vet ./...

build: vet
	go build 

run: build
	./ch2

clean:
	go clean
