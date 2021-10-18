export GO111MODULE=on

.PHONY: setup
setup:
	GO111MODULE=off go get \
		golang.org/x/lint/golint

.PHONY: lint
lint:
	go vet ./...
	golint ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./...

.PHONY: test-bench
test-bench:
	go test -bench . -benchmem

.PHONY: all
all: lint fmt test

.PHONY: build
build: lint fmt test
	go build -o issues main.go
