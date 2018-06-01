#!/bin/bash

export CGO_CXXFLAGS_ALLOW=".*" 
export CGO_LDFLAGS_ALLOW=".*" 
export CGO_CFLAGS_ALLOW=".*" 

rm -rf leveldb-test kvdb

protoc -I server/pb server/pb/service.proto --go_out=plugins=grpc:server/pb
# or
# without -I , it means
# protoc server/pb/service.proto  --go_out=plugins=grpc:.

go build  -v 
