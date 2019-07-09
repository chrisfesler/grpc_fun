PROTO_SRC=$(shell find . -name '*.proto')
PROTO_GEN=$(PROTO_SRC:.proto=.pb.go)
PROTO_PLUGINS=github.com/golang/protobuf/protoc-gen-go	

BINDIR?=$(PWD)/bin

GOPATH?=$(go env GOPATH)
PROTO_BIN?=$(BINDIR)/proto-plugins

export GOBIN=$(BINDIR)
export PATH:=$(GOBIN):$(PATH)
export GO111MODULE=on

build: clean proto vet install

clean: proto-clean
	rm -rf $(BINDIR)

proto-clean:
	rm $(PROTO_GEN) || true

vet:
	go vet ./...

protoc-plugins:
	GOBIN=$(PROTO_BIN) go install $(PROTO_PLUGINS)

proto: protoc-plugins
	protoc -I . $(PROTO_SRC) --go_out=plugins=grpc,paths=source_relative:.

install:
	go install ./cmd/...

