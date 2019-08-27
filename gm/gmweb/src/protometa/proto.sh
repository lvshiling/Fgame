#!/bin/sh
protoc.exe --js_out=import_style=commonjs,binary:. basic.proto
protoc.exe --js_out=import_style=commonjs,binary:. common.proto
protoc.exe --js_out=import_style=commonjs,binary:. login.proto
protoc.exe --js_out=import_style=commonjs,binary:. messagetype.proto
protoc.exe --js_out=import_style=commonjs,binary:. chat.proto
protoc.exe --js_out=import_style=commonjs,binary:. chatmessagetype.proto