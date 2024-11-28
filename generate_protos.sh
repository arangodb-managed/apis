#!/bin/bash

# Check if the package path is provided as an argument
if [ -z "$1" ]; then
  echo "ERROR: Usage: $0 <proto-file>"
  exit 1
fi

PROTO_FILE=$1

## Run protoc with the necessary arguments
protoc \
  -I .:../../:../../vendor/:../../vendor/googleapis/:../../vendor/github.com/gogo/protobuf/protobuf/ \
  --gofast_out=Mgithub.com/golang/protobuf/ptypes/duration/duration.proto=github.com/gogo/protobuf/types,Mgithub.com/golang/protobuf/ptypes/timestamp/timestamp.proto=github.com/gogo/protobuf/types,plugins=grpc,paths=source_relative:. \
  --grpc-gateway_out=paths=source_relative,logtostderr=true,allow_delete_body=true:. \
  ./$PROTO_FILE
