BUILD ?= CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-s -w'

GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

.PHONY: build
build: proto

	$(BUILD) -o user-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-srv:latest
