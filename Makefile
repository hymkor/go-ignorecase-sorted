ifeq ($(OS),Windows_NT)
    SHELL=CMD.EXE
    SET=set
else
    SET=export
endif

build:
	go fmt ./...

test:
	go test -v

.PHONY: build test readme
