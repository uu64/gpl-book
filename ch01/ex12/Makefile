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

.PHONY: all
all: lint fmt test