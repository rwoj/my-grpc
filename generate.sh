#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
# option go_package="greet/greetpb"

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.