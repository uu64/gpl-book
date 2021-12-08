export GO111MODULE=on

.PHONY: setup
setup:
	go install honnef.co/go/tools/cmd/staticcheck@latest

.PHONY: lint
lint:
	go vet ./...
	staticcheck ./...

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
	go build -o clock clock.go
	go build -o clockwall clockwall.go
