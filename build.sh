#!/bin/bash
set -o errexit
set -o nounset

if [ $# != 1 ]; then
    echo "Please add argument, like: ./build.sh kms-server"
    exit -1
fi

target=$1
readonly OUTPUT_DIR="output"
export GOPATH=$(pwd):$(pwd)/vendor

rm -fr ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}/{bin,etc,run,log}

build_kms_server() {
    cd src && go build -o ../${OUTPUT_DIR}/bin/kms-server ./kms-server && cd - >/dev/null 2>&1 || exit -1
    cd src && go build -o ../${OUTPUT_DIR}/bin/kms-tool ./kms-tool && cd - >/dev/null 2>&1 || exit -1
    cd src && go build -o ../${OUTPUT_DIR}/bin/kms-ipwhite ./kms-timer/ipwhite && cd - >/dev/null 2>&1 || exit -1
    cd gwc && make > /dev/null && cp -af gwc ../${OUTPUT_DIR}/bin/ && cd - > /dev/null 2>&1 || exit -1
    cp -af scripts/control_kms_server.sh ${OUTPUT_DIR}/bin/control
    chmod +x ${OUTPUT_DIR}/bin/control
}

build_kms_cap() {
    cd src && go build -o ../${OUTPUT_DIR}/bin/kms-cap ./kms-cap && cd - >/dev/null 2>&1 || exit -1
    cd gwc && make > /dev/null && cp -af gwc ../${OUTPUT_DIR}/bin/ && cd - > /dev/null 2>&1 || exit -1
    cp -af scripts/control_kms_cap.sh ${OUTPUT_DIR}/bin/control
    chmod +x ${OUTPUT_DIR}/bin/control
}

case $target in
    kms-server)
        build_kms_server
        ;;
    kms-cap)
        build_kms_cap
        ;;
    *)
        exit -1
        ;;
esac
