#!/bin/sh
protoc --proto_path=../../../../../../../../ --proto_path=./ --js_out=import_style=commonjs,binary:./ messagetype.proto