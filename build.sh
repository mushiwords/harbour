#!/bin/bash
set -o errexit
set -o nounset

target=$1
readonly OUTPUT_DIR="output"
export GOPATH=$(pwd):$(pwd)/vendor

rm -fr ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}/{bin,etc,run,log}

build_harbour() {
    cd src && go build -o ../${OUTPUT_DIR}/bin/harbour . && cd - >/dev/null 2>&1 || exit -1
    cd gwc && make > /dev/null && cp -af gwc ../${OUTPUT_DIR}/bin/ && cd - > /dev/null 2>&1 || exit -1
}
build_harbour