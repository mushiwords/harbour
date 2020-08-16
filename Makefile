export GO111MODULE=on
export GOPROXY=https://goproxy.io

OBJ = monitor

all: $(OBJ)

$(OBJ):
	cd src && go build -gcflags "-N -l" -o ../$@ .

-include .deps

dep:
	echo -n "$(OBJ):" > .deps
	find src -name '*.go' | awk '{print $$0 " \\"}' >> .deps
clean:
	rm -fr $(OBJ)

