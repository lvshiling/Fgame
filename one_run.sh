#!/bin/sh

# 工作目录
WORKPATH=$(cd `dirname $0`; pwd)
echo "当前目录$WORKPATH"

# 获取服务器名字
servername=`echo "$WORKPATH" | awk -F'/' '{print $NF}'`



game(){
        INDEX=$2
        cd $WORKPATH/$INDEX/game

        case "$1" in
        pull)
                sh run.sh game pull
                ;;
        stop)
                sh run.sh game stop all
                ;;
        start)
                sh run.sh game start all
                ;;
        update)
                sh run.sh game update all
                ;;
        upgrade)
                sh run.sh game upgrade all
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}

group(){
        INDEX=$2
        cd $WORKPATH/$INDEX/group

        case "$1" in
        pull)
                sh run.sh group pull
                ;;
        stop)
                sh run.sh group stop all
                ;;
        start)
                sh run.sh group start all
                ;;
        update)
                sh run.sh group update all
                ;;
        upgrade)
                sh run.sh group upgrade all
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}

region(){
        INDEX=$2
        cd $WORKPATH/$INDEX/region

        case "$1" in
        pull)
                sh run.sh region pull
                ;;
        stop)
                sh run.sh region stop all
                ;;
        start)
                sh run.sh region start all
                ;;
        update)
                sh run.sh region update all
                ;;
        upgrade)
                sh run.sh region upgrade all
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}

platform(){
        INDEX=$2
        cd $WORKPATH/$INDEX/platform

        case "$1" in
        pull)
                sh run.sh platform pull
                ;;
        stop)
                sh run.sh platform stop all
                ;;
        start)
                sh run.sh platform start all
                ;;
        update)
                sh run.sh platform update all
                ;;
        upgrade)
                sh run.sh platform upgrade all
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}


# 开始,更新,状态,停止
case "$1" in
        pull)
                # 获取剩余参数
                serverList=()
                if [ "$2" = "all" ]
                then
                        for fgameDir in `ls $WORKPATH`
                        do
                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir = "fgame" ]]
                                then
                                        cd $WORKPATH/$fgameDir
                                        sh pull.sh
                                fi
                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir =~ fgame_[a-z]+ ]]
                                then
                                        cd $WORKPATH/$fgameDir
                                        sh pull.sh
                                fi
                        done
                fi
                ;;   
        game)
                case "$2" in 
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for fgameDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir = "fgame" ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir =~ fgame_[a-z]+ ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                        done
                                
                                else
                                        # 单个或多个启动
                                        serverList=($@)
                                        serverList=(${serverList[@]:2})
                                fi
                                # 启动游戏服
                                for serverIndex in "${serverList[@]}"
                                do
                                        game "$2" $serverIndex
                                done
                                ;;
                esac
                ;;
        group)
                case "$2" in 
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for fgameDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir = "fgame" ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir =~ fgame_[a-z]+ ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                        done
                                
                                else
                                        # 单个或多个启动
                                        serverList=($@)
                                        serverList=(${serverList[@]:2})
                                fi
                                # 启动游戏服
                                for serverIndex in "${serverList[@]}"
                                do
                                        group "$2" $serverIndex
                                done
                                ;;
                esac
                ;;
        region)
                case "$2" in 
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for fgameDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir = "fgame" ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir =~ fgame_[a-z]+ ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                        done
                                
                                else
                                        # 单个或多个启动
                                        serverList=($@)
                                        serverList=(${serverList[@]:2})
                                fi
                                # 启动游戏服
                                for serverIndex in "${serverList[@]}"
                                do
                                        region "$2" $serverIndex
                                done
                                ;;
                esac
                ;;
        platform)
                case "$2" in 
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for fgameDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir = "fgame" ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                                if [[ -d "$WORKPATH/$fgameDir" && $fgameDir =~ fgame_[a-z]+ ]]
                                                then
                                                        serverList+=("$fgameDir")
                                                fi
                                        done
                                
                                else
                                        # 单个或多个启动
                                        serverList=($@)
                                        serverList=(${serverList[@]:2})
                                fi
                                # 启动游戏服
                                for serverIndex in "${serverList[@]}"
                                do
                                        platform "$2" $serverIndex
                                done
                                ;;
                esac
                ;;
        *)
                echo "请输入account,game参数"
                exit 1
                ;;
esac