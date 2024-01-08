.PHONY: dc run test lint

dc:
	echo 'hello'

run:
	go run cmd/server/main.go

test:
	go test -race ./...

lint:
	golangci-lint run

all:
	echo 'hello2'