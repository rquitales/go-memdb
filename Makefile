# Copyright (c) 2020 Ramon Quitales
# 
# This software is released under the MIT License.
# https://opensource.org/licenses/MIT

all: vet test build

fmt:	
	goimports -d -e -l -w .
	gofmt -d -e -l -s -w .
build:	
	go build -v .
run:	
	go run .
test:	
	go test -v . -coverprofile /tmp/go-memdb.txt
vet:
	go vet ./...
deps:	
	go mod tidy 
build-linux: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v .
build-mac: 
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -v .