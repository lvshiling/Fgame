#!/bin/sh

# 工作目录
WORKPATH=$(pwd)
echo "当前目录$WORKPATH"


# 发布路径
distPath=""
ip=""
port=""
destPath=""

case "$1" in
    test)
        distPath="deploy/wanshi_test"
        ip="47.98.43.104"
        port="20000"
        destPath="~/fgame/deploy/"
    ;;
    check)
        distPath="deploy/wanshi_check"
        ip="47.99.190.36"
        port="20000"
        destPath="~/fgame_shenhe/deploy/"
    ;;
    *)
        echo "不支持的参数,请输入prod,test,inner_test"
        exit 1
esac

echo "发布路径:$WORKPATH/$distPath"

rsync -azP --delete -e "ssh -p $port" "$WORKPATH/$distPath/" root@$ip:$destPath