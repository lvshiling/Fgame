#!/bin/sh

protoc --proto_path=./ui/ --go_out=../common/codec/pb/ui/ ./ui/*.proto 
protoc --proto_path=./scene/ --go_out=../common/codec/pb/scene/ ./scene/*.proto 
protoc --proto_path=./cross/ --go_out=../common/codec/pb/cross/ ./cross/*.proto 
