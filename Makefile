GOVERSION = $(shell go version | awk '{print $$3;}')

export CGO_ENABLED := 0

clean:
	rm -rf ./dist
.PHONY: clean

# test:
# 	gotestsum -- -failfast -v -covermode count -timeout 5m ./...
# .PHONY: test

build:
	go build -o ./dist/go-pub-sub ./
.PHONY: build
