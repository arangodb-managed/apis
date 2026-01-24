#!/bin/bash

# Check if the package path is provided as an argument
if [ -z "$1" ]; then
  echo "ERROR: Usage: $0 <proto-file>"
  exit 1
fi

PROTO_FILE=$1

echo "Generating Go code from $PROTO_FILE"

protoc \
  --proto_path . \
    -I ../../vendor-proto/ \
    -I ../../ \
  --go_out . \
    --go_opt=plugins=grpc,paths=source_relative, \
    --go_opt=Mgoogle/protobuf/timestamp.proto=google.golang.org/protobuf/types/known/timestamppb \
    --go_opt=Mgoogle/protobuf/duration.proto=google.golang.org/protobuf/types/known/durationpb \
    --go_opt=Mgoogle/protobuf/empty.proto=google.golang.org/protobuf/types/known/emptypb \
    --go_opt=Mgoogle/protobuf/descriptor.proto=google.golang.org/protobuf/types/known/descriptorpb \
  --grpc-gateway_out=. \
    --grpc-gateway_opt=logtostderr=true \
    --grpc-gateway_opt=paths=source_relative \
    --grpc-gateway_opt=allow_delete_body=true \
  ./$PROTO_FILE
