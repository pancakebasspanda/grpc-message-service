#!/bin/bash

protoc -I . --go_out=plugins=grpc:protos message_service.proto
cp -r protos/message_service/protos/message_service.pb.go protos
rm -rf protos/message_service/
