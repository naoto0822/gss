TARGET := ./atom ./session ./rss1

## exec all task
all: dep vet lint test

## install dependencies
dep:
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/golang/lint/golint
	go get github.com/Songmu/make2help/cmd/make2help

## run vet 
vet:
	go vet $(TARGET)

## check lint
lint:
	golint -set_exit_status $(TARGET)

## run test
test:
	go test -v $(TARGET)

help:
	@make2help $(MAKEFILE_LIST)

.PHONY: dep vet lint test help
