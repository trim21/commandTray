GO_FILES := $(shell find . -type f -name "*.go")

default: dist/app

dist/app: $(GO_FILES)
	go build -mod=readonly -o dist/app

install:
	go mod -mod=readonly download
