#!/bin/sh

# 工作目录
WORKPATH=$(pwd)
echo "当前目录$WORKPATH"

# 发布路径
distPath=""
ip=""
port=""
destPath=""



echo "发布路径:$WORKPATH/$distPath"

rsync -azP --delete -e "ssh -p $port" "$WORKPATH/$distPath/" root@$ip:$destPath