# 基础变量
GOCMD=go
PACKCMD=tar
GOENV=$(GOCMD) env
GOBUILD=$(GOCMD) build
GOPATH=$(shell go env GOPATH)

# 输出文件
BINARY_NAME=tcompiler

# make设置
.PHONY: launch build pb
.DEFAULT_GOAL := launch

launch:
	$(GOBUILD) -o build/$(BINARY_NAME) cmd/main.go && ./build/${BINARY_NAME}
build:
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o build/$(BINARY_NAME) cmd/main.go
