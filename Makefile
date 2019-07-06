PROTO_SRC=$(shell find . -name '*.proto')
PROTO_GEN=$(PROTO_SRC:.proto=.pb.go)

BINDIR?=$(PWD)/bin

GOPATH?=$(go env GOPATH)

export GOBIN=$(BINDIR)
export PATH:=$(GOBIN):$(PATH)
export GO111MODULE=on

clean: proto-clean
	rm -rf $(BINDIR)

proto-clean:
	rm $(PROTO_GEN) || true

build: clean proto vet go-build install

vet:
	go vet ./...

go-build:
	go build ./...

proto:
	protoc -I . $(PROTO_SRC) --go_out=plugins=grpc,paths=source_relative:.

install:
	go install ./cmd/...

