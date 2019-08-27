#!/bin/sh
 
protoc3 --go_out=plugins=grpc:./pb/ ./*.proto
