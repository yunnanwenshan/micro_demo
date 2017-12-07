#!/usr/bin/env bash

set -e

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
OLDGOBIN="$GOBIN"
export GOPATH="$CURDIR"
export GOBIN="$CURDIR/bin/"
echo 'GOPATH:' $GOPATH
echo 'GOBIN:' $GOBIN

#go get github.com/micro/go-micro 
#go get github.com/micro/protobuf/protoc-gen-go 
#go get github.com/micro/protobuf/proto
#go get github.com/micro/micro
#go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
#go get -u github.com/micro/protobuf/protoc-gen-go
go get github.com/micro/go-grpc
export GOPATH="$OLDGOPATH"
export GOBIN="$OLDGOBIN"

echo 'build finished'
