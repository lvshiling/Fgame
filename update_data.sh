#!/bin/sh

WORKPATH=$(pwd)
echo "当前目录$WORKPATH"

PROD_DATA_PATH="https://192.168.1.13:18080/svn/project1/branch/main/Export/server/data"
PROD_MAP_PATH="https://192.168.1.13:18080/svn/project1/branch/main/map"

MASTER_DATA_PATH="https://192.168.1.13:18080/svn/project1/trunk/Export/server/data"
MASTER_MAP_PATH="https://192.168.1.13:18080/svn/project1/trunk/map"

SHENHE_DATA_PATH="https://192.168.1.13:18080/svn/project1/branch/shenhe/Export/server/data"
SHENHE_MAP_PATH="https://192.168.1.13:18080/svn/project1/branch/shenhe/map"

dataPath=""
mapPath=""

tempDataPath="template"
tempMapPath="map"

case "$1" in
    prod)
        echo "正在检出线上数据"
        dataPath=$PROD_DATA_PATH
        mapPath=$PROD_MAP_PATH
    ;;
    master)
        echo "正在检出主干数据"
        dataPath=$MASTER_DATA_PATH
        mapPath=$MASTER_MAP_PATH
    ;;
    shenhe)
        echo "正在检出审核数据"
        dataPath=$SHENHE_DATA_PATH
        mapPath=$SHENHE_MAP_PATH
    ;;
    *)
        echo "不支持的参数,请输入prod,master"
        exit 1
esac

tempDir="$WORKPATH/.temp"
echo "创建临时目录"

if [ ! -d "$tempDir" ];then
    mkdir "$tempDir"
fi

echo "正在检出策划数据"
svn checkout "$dataPath" "$tempDir/$tempDataPath"

echo "正在检出地图数据"
svn checkout "$mapPath" "$tempDir/$tempMapPath"

echo "复制地图数据"
# 复制地图数据
cd "$tempDir/map"
for f in *.txt 
do
    cp -rf "$f" "$WORKPATH/resources/map/"
done

echo "复制策划数据"
cd "$tempDir/template"
for f in *.json 
do
    cp -rf "$f" "$WORKPATH/resources/template/"
done

echo "移除临时目录"
rm -rf "$tempDir"

