#!/bin/bash


export GOPATH=$(shell pwd):$(shell pwd)/vendor

default: kms-server kms-cap kms-tool kms-ipwhite

kms-server kms-cap kms-tool:
	cd src && go build -gcflags "-N -l" -o ../$@ ./$@

kms-ipwhite:
	cd src && go build -gcflags "-N -l" -o ../$@ ./kms-timer/ipwhite

-include .deps

dep:
	echo -n "kms-server kms-cap kms-tool kms-ipwhite:" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps

clean:
	rm -fr kms-server kms-cap kms-tool kms-ipwhite

