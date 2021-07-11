#!/bin/bash
protoc handOne/greet/greetpb/greet.proto --go_out=plugins=grpc:.