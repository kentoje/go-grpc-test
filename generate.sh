#! /bin/bash

# unary
protoc unary/greet/greetpb/greet.proto --go_out=plugins=grpc:./unary
protoc unary/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:./unary

# streaming-server
protoc streaming-server/greet/greetpb/greet.proto --go_out=plugins=grpc:./streaming-server