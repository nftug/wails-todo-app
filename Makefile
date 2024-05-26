.PHONEY: all build run test prepare wire deps wire_deps

all: test build

build: prepare
	wails build

run: prepare
	wails dev

test: prepare
	go test -v -cover `go list ./... | grep -v presentation | grep -v wails-todo-app$$` -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

prepare: wire

wire: wire_deps
	wire gen ./...\

deps: wire_deps

wire_deps:
ifeq ($(shell command -v wire 2> /dev/null),)
	go install github.com/google/wire/cmd/wire@latest
endif
