.PHONEY: all build run test prepare deps

all: test build

build: prepare
	wails build -ldflags="-s -w" -trimpath

# ネイティブバイナリを使用しないので不要
# winbuild: prepare
# 	wails generate module
#	env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc \
		wails build -ldflags="-s -w" -trimpath -skipbindings

run: prepare
	wails dev

test: prepare
	go test -v -cover `go list ./... | grep -v presentation | grep -v wails-todo-app$$` -coverprofile=cover.out
	go tool cover -html=cover.out -o cover.html

prepare:
