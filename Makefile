#!/bin/bash


export GOPATH=$(shell pwd):$(shell pwd)/vendors/src

default: harbour

harbour:
	cd src && go build -gcflags "-N -l" -o ../$@ .

-include .deps

dep:
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps

clean:
	rm -fr harbour

