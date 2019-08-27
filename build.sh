#!/bin/sh


#复制文件 
copyFile(){
    local src=$1
    local dir=$2
    local name=$3
    if [ ! -d $dir ]
    then
        mkdir -p $dir
    fi
    cp "$src" "$dir/$name"
}

#检查数组包含
containStr(){
    arr=$1
    for str in ${arr[@]}
    do  
        if [ "$str" = "$2" ]
        then
            return 1
        fi
    done
    return 0
}


# 编译
build(){
    local src=$1
    echo "build $src"
    cd $WORKPATH/$src && sh build.sh
    echo "build $src 完成"
}

# TODO 添加版本号 数据库
# 可以编译的列表,用来检验参数
buildList=("account" "center" "charge_server" "cross" "game" "logserver" "gm" "coupon_server" "trade_server")
# 工作目录
WORKPATH=$(pwd)
echo "当前目录$WORKPATH"

# 发布路径
distPath=""

case "$1" in
    master)
        distPath="deploy/wanshi_master"
    ;;
    test)
        distPath="deploy/wanshi_test"
    ;;
    inner_test)
        distPath="deploy/wanshi_inner_test"
    ;;
    prod)
        distPath="deploy/wanshi"
    ;;
    check)
        distPath="deploy/wanshi_check"
    ;;
    *)
        echo "不支持的参数,请输入prod,test,inner_test"
        exit 1
esac


# 需要编译的
needBuildList=()
if [ "$2" = "all" ];
then
    needBuildList=(${buildList[@]})
else
    needBuildList=($@)
    needBuildList=(${needBuildList[@]:1})
fi

# 检查参数
for needBuild in "${needBuildList[@]}"
do
    containStr "${buildList[*]}" $needBuild
    ret=$?
    if [ $ret -eq 0 ];
    then
        echo "输入参数[$needBuild]错误,不能构建"
        exit 1
    fi
done

# 编译
for needBuild in "${needBuildList[@]}"
do
    build $needBuild
done

echo "发布路径:$WORKPATH/$distPath"

# 复制二进制文件
for needBuild in "${needBuildList[@]}"
do
    echo "复制$needBuild服"
    copyFile "$WORKPATH/$needBuild/main" "$WORKPATH/$distPath/$needBuild" $needBuild
    echo "复制$needBuild服完成"
done

mapDir="$WORKPATH/$distPath/resources/map"
templateDir="$WORKPATH/$distPath/resources/template"
# 复制数据
echo "复制策划数据"
if [ ! -d $mapDir ]
then
    mkdir -p $mapDir
fi

# 复制地图数据
cd "$WORKPATH/resources/map"
for f in *.txt 
do
    cp -rf "$f" "$mapDir/"
done

# 复制模板数据
if [ ! -d $templateDir ]
then
    mkdir -p $templateDir
fi

cd "$WORKPATH/resources/template"
for f in *.json 
do
    cp -rf "$f" "$templateDir/"
done
echo "复制策划数据完成"

