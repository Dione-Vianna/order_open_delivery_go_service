.PHONY: run build

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

start: build
	./bin/server
