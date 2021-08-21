export GO111MODULE=on

.PHONY: setup
setup:
	GO111MODULE=off go get \
		golang.org/x/lint/golint

.PHONY: lint
lint:
	go vet ./...
	golint -set_exit_status ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: test
test:
	go test ./... --count=1

.PHONY: all
all: lint fmt test