#!/bin/bash
# protoc handOne/greet/greetpb/greet.proto --go_out=plugins=grpc:.
protoc handOne/calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.