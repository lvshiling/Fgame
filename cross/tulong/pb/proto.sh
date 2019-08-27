#!/bin/sh
 
protoc3 --go_out=plugins=grpc:. *.proto