#! /bin/sh

# 工作目录
WORKPATH=$(pwd)
echo "当前目录$WORKPATH"


# 发布路径
distPath=""
ip=""
port=""
destPath=""

case "$1" in
    master)
        distPath="deploy/wanshi_master"
        ip="192.168.1.123"
        port="22"
        destPath="~/fgame_master/deploy/"
    ;;
    test)
        distPath="deploy/wanshi_test"
        ip="47.99.185.231"
        port="20000"
        destPath="~/fgame/deploy/"
    ;;
    inner_test)
        distPath="deploy/wanshi_inner_test"
        ip="192.168.1.123"
        port="22"
        destPath="~/fgame/deploy/"
    ;;
    prod)
        distPath="deploy/wanshi"
        destPath="~/fgame/deploy/"
    ;;
    *)
        echo "不支持的参数,请输入prod,test,inner_test"
        exit 1
esac

echo "发布路径:$WORKPATH/$distPath"

rsync -azP --delete -e "ssh -p $port" "$WORKPATH/$distPath/" root@$ip:$destPath