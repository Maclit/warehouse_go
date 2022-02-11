.PHONY: all build run lint

all: run

build: go build .

run: go run .

lint: golangci-lint run