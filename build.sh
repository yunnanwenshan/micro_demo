#!/usr/bin/env bash

set -e

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
OLDGOBIN="$GOBIN"
export GOPATH="$CURDIR"
export GOBIN="$CURDIR/bin/"
echo 'GOPATH:' $GOPATH
echo 'GOBIN:' $GOBIN
go build -race -o micro -gcflags "-N -l"  src/github.com/micro/micro/main.go
go build -race -o srv -gcflags "-N -l" src/github.com/micro/examples/greeter/srv/main.go 
go build -race -o service -gcflags "-N -l" src/test/service.go 
go build -race -o gateway -gcflags "-N -l" src/test/gateway.go 

if [ ! -d ./bin ]; then
    mkdir bin
fi

if [ -e ./micro ]; then
   mv micro ./bin/
fi

if [ -e ./srv ]; then
   mv srv ./bin/
fi

export GOPATH="$OLDGOPATH"
export GOBIN="$OLDGOBIN"

echo 'build finished'
