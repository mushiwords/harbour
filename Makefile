#!/bin/bash


export GOPATH=$(shell pwd):$(shell pwd)/vendors

OBJ = harbour

all: $(OBJ)

$(OBJ):
	cd src && go build -gcflags "-N -l" -o ../$@ .

-include .deps

dep:
	echo -n "$(OBJ):" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps

clean:
	rm -fr $(OBJ)
