#!/bin/bash

protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.
protoc blog/blogpb/blog.proto --go_out=plugins=grpc:.


 
protoc greet/greetpb/greet.proto --go-grpc_out=plugins=grpc:./greet/
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:./calculator
protoc blog/blogpb/blog.proto --go_out=plugins=grpc:./blog

protoc greet/greetpb/greet.proto --go-grpc_out=.

protoc --go-grpc_out=. greet/greetpb/greet.proto  
protoc --go-grpc_out=. file_path/file_name*.proto 

protoc -I ./greet \
   --go_out ./greet --go_opt paths=source_relative \
   --go-grpc_out ./greet --go-grpc_opt paths=source_relative \
   ./greet/greetpb/greet.proto
