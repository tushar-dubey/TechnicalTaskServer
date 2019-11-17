#!/bin/bash

protoc calculator/proto/calculator.proto --go_out=plugins=grpc:.