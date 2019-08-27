#!/bin/sh

protoc --proto_path=./ --go_out=../pb ./*.proto 
