APP_NAME=CICDRef
BINARY=bin/api

.PHONY: tidy test build run lint clean

tidy:
	go mod tidy

test:
	go test -v ./...

build:
	mkdir -p bin
	go build -o $(BINARY) ./cmd/api

run:
	go run ./cmd/api

lint:
	go vet ./...

clean:
	rm -rf bin