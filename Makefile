.PHONEY: all build run test prepare deps

all: test build

build: prepare
	wails build -ldflags="-s -w" -trimpath

run: prepare
	wails dev

test: prepare
	go test -v -cover `go list ./... | grep -v presentation | grep -v wails-todo-app$$` -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

prepare:
