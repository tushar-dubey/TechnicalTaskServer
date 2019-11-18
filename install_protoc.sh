#!/bin/bash

PROTOBUF_VERSION=3.3.0
PROTOC_FILENAME=protoc-${PROTOBUF_VERSION}-linux-x86_64.zip
  - go get -v -d google.golang.org/grpc
  - go get -v -d -t github.com/golang/protobuf/...
  - curl -L https://github.com/google/protobuf/releases/download/v3.6.1/protoc-3.6.1-linux-x86_64.zip -o /tmp/protoc.zip
  - unzip /tmp/protoc.zip -d "$HOME"/protoc
  - mkdir -p "$HOME"/src && ln -s "$HOME"/protoc "$HOME"/src/protobuf