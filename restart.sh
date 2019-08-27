#!/bin/bash

WORKPATH=$(cd `dirname $0`; pwd)
echo "当前目录$WORKPATH"
# 获取服务器列表
serverList=()
suffix=""
# 查找游戏服
for gameDir in `ls $WORKPATH`
do
        if [[ -d "$WORKPATH/$gameDir" && $gameDir =~ game_[0-9]+ ]]
        then
            index=$(echo $gameDir | cut -d "_" -f2)
            echo "游戏目录$gameDir,索引$index"
            serverList+=("$index")
        fi
done

 
while true  
do 
    for serverIndex in "${serverList[@]}"
    do
        program="game_$suffix$serverIndex"
        result=`pidof $program`
        if [ -z "$result" ]
        then
            echo "$(date),服务器$program已经关闭,准备重启" 
            sh run.sh game restart $serverIndex  
        fi
    done
    sleep 5 
done

