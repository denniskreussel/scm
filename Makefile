VERSION=$(shell git describe --tags --dirty)
SCM_LDFLAGS += -X "main.version=$(VERSION)"
SCM_LDFLAGS += -X "main.date=$(shell date --iso-8601=s)"
SCM_LDFLAGS += -X "main.commit=$(shell git rev-parse HEAD)"
SCM_LDFLAGS += -X "main.builtBy=$(shell echo `whoami`@`hostname`)"
DEFAULT_CFG_PATH = /etc/scm/scm.config.yaml

GO        := CGO_ENABLED=0 go
GOBUILD   := $(GO) build $(BUILD_FLAG)


.PHONY: pre
pre:
	go mod tidy

.PHONY: build
build: pre
	$(GOBUILD) -ldflags '$(SCM_LDFLAGS)' -o bin/ ./...

.PHONY: debug
debug: pre
	$(GOBUILD) -ldflags '$(SCM_LDFLAGS)' -gcflags "all=-N -l" -o bin/ ./...

.PHONY: gotest
gotest: pre
	go test -v ./... -coverprofile=coverage.out -covermode count
	go tool cover -func coverage.out

.PHONY: lint
lint:
	golangci-lint run -D errcheck,govet,gosimple

.PHONY: run
run: pre
	go run cmd/scm/main.go
