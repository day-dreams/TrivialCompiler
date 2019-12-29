# 基础变量
GOCMD=go
PACKCMD=tar
GOENV=$(GOCMD) env
GOBUILD=$(GOCMD) build
GOPATH=$(shell go env GOPATH)

# 输出文件
BINARY_NAME=tcompiler

# make设置
.PHONY: launch build gocc test
.DEFAULT_GOAL := build

gocc:
# 	gocc -p github.com/day-dreams/TrivialCompiler bnf/tcompiler.bnf
	gocc -p github.com/day-dreams/TrivialCompiler bnf/cmd.bnf
launch:
	$(GOBUILD) -o build/$(BINARY_NAME) cmd/main.go && ./build/${BINARY_NAME}
build:
	$(GOBUILD) -o build/$(BINARY_NAME) cmd/main.go
test:
	go test -v test/*.go
install: build
	cp build/${BINARY_NAME} /usr/local/bin/${BINARY_NAME}
uninstall:
	rm /usr/local/bin/${BINARY_NAME}

