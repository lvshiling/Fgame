#!/bin/sh

# 工作目录
WORKPATH=$(cd `dirname $0`; pwd)
echo "当前目录$WORKPATH"

# 获取服务器名字
servername=`echo "$WORKPATH" | awk -F'/' '{print $NF}'`
suffix=""

account(){
        exeFile="./account"
        pidFile="account.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,登陆服已经启动或奔溃"
                        else
                                echo "正在启动登陆服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "登陆服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "登陆服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,登陆服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止登陆服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候登陆服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "登陆服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ~/fgame/deploy/account/account ./
                        chmod 777 account
                ;;
                status)
                ;;
                *)  
                        echo "账户服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}

center(){
        exeFile="./center"
        pidFile="center.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,中心服已经启动或奔溃"
                        else
                                echo "正在启动中心服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "中心服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "中心服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,中心服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止中心服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候中心服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "中心服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ~/fgame/deploy/center/center ./
                        chmod 777 center
                ;;
                status)
                ;;
                *)  
                        echo "账户服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}


charge_server(){
        exeFile="./charge_server"
        pidFile="charge_server.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,充值服已经启动或奔溃"
                        else
                                echo "正在启动充值服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "充值服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "充值服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,充值服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止充值服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候充值服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "充值服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ~/fgame/deploy/charge_server/charge_server ./
                        chmod 777 charge_server
                ;;
                status)
                ;;
                *)  
                        echo "账户服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}


coupon_server(){
        exeFile="./coupon_server"
        pidFile="coupon_server.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,coupon服已经启动或奔溃"
                        else
                                echo "正在启动coupon服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "coupon服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "coupon服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,coupon服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止coupon服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候coupon服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "coupon服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ~/fgame/deploy/coupon_server/coupon_server ./
                        chmod 777 coupon_server
                ;;
                status)
                ;;
                *)  
                        echo "coupon服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}


trade_server(){
        exeFile="./trade_server"
        pidFile="trade_server.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,trade服已经启动或奔溃"
                        else
                                echo "正在启动trade服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "trade服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "trade服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,trade服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止trade服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候trade服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "trade服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ~/fgame/deploy/trade_server/trade_server ./
                        chmod 777 trade_server
                ;;
                status)
                ;;
                *)  
                        echo "trade服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}



cross(){
        exeFile="./cross"
        pidFile="cross.pid"
        confFile="./config/config.json"
        case "$1" in
                start)
                        if [ -f $pidFile ] 
                        then
                                echo "$pidFile 已经存在,全平台服已经启动或奔溃"
                        else
                                echo "正在启动登陆服"
                                nohup $exeFile --config $confFile >>/dev/null 2>&1 &
                                PID=$!
                        if [ -x /proc/${PID} ]
                        then
                                echo "全平台跨服启动成功"
                                echo $PID>$pidFile
                        else
                                echo "全平台跨服启动失败"
                        fi 
                        fi
                ;;
                stop)
                        if [ ! -f $pidFile ]
                        then
                                echo "$pidFile 不存在,全平台跨服已经关闭"
                        else
                                PID=$(cat $pidFile)
                                echo "正在停止全平台跨服"
                                kill $PID
                                while [ -x /proc/${PID} ]
                                do
                                echo "等候全平台跨服关闭"
                                sleep 1
                                done
                                rm $pidFile
                                echo "全平台跨服关闭中"
                        fi
                ;;
                restart)
                ;;
                update)
                ;;
                pull)
                        rsync -azP --delete ../deploy/cross/cross ./
                        rsync -azP --delete ../deploy/resources/* ./resources/  
                        chmod 777 cross
                ;;
                status)
                        if [ -f $pidFile ]       
                        then
                                PID=$(cat $pidFile)
                                if [ -x /proc/${PID} ]
                                then
                                        echo "全平台跨服正在运行"
                                else 
                                        echo "全平台跨服已经停止"
                                fi
                        else
                                echo "全平台跨服已经停止"
                        fi
                       
                ;;
                *)  
                        echo "全平台跨服务器请输入start,stop,restart,status,update参数"
                        exit 1
                ;;
        esac
}

groupStart(){
        INDEX=$1
        EXEC=./group_$suffix$INDEX
        PIDFILE=./group_$INDEX.pid
        CONF=./config/group_$INDEX.json
        cd $WORKPATH/group_$INDEX

        if [ -f $PIDFILE ]
        then
                echo "$PIDFILE 已经存在,组服($INDEX)已经启动或奔溃"
        else
                
                echo "正在启动组服($INDEX)"
                nohup $EXEC --config $CONF >>/dev/null 2>&1 &
                PID=$!
                if [ -x /proc/${PID} ]
                then
                        echo "组服($INDEX)启动成功"
                        echo $PID>$PIDFILE
                else
                        echo "组服($INDEX)启动失败"
                fi       
        fi           
}

groupStop(){
        INDEX=$1
        EXEC=./group_$suffix$INDEX
        PIDFILE=./group_$INDEX.pid
        CONF=./config/group_$INDEX.json
        cd $WORKPATH/group_$INDEX
        if [ ! -f $PIDFILE ]
        then
                echo "$PIDFILE 不存在,组服($INDEX)已经关闭"
        else
                PID=$(cat $PIDFILE)
                echo "正在停止组服($INDEX)"
                kill $PID
                while [ -x /proc/${PID} ]
                do
                echo "等候组服($INDEX)关闭"
                sleep 1
                done
                rm $PIDFILE
                echo "组服($INDEX)关闭中"
        fi         
}

group(){
        INDEX=$2
        EXEC=./group_$suffix$INDEX
        PIDFILE=./group_$INDEX.pid
        CONF=./config/group_$INDEX.json
        cd $WORKPATH/group_$INDEX

        case "$1" in
        start)
                groupStart $INDEX
                ;;
        stop)
                groupStop $INDEX
                ;;
        status)
                if [ -f $PIDFILE ]       
                then
                        PID=$(cat $PIDFILE)
                        if [ -x /proc/${PID} ]
                        then
                                echo "组服($INDEX)正在运行"
                        else 
                                echo "组服($INDEX)已经停止"
                        fi
                else
                        echo "组服($INDEX)已经停止"
                fi
                ;;
        update)
                cp ../group ./group_$suffix$INDEX
                echo "组服($INDEX)更新成功"
                ;;
        upgrade)
                groupStop $INDEX
                cp ../group ./group_$suffix$INDEX
                echo "组服($INDEX)更新成功"
                groupStart $INDEX
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}

gameStart(){
        INDEX=$1
        EXEC=./game_$suffix$INDEX
        PIDFILE=./game_$INDEX.pid
        CONF=./config/game_$INDEX.json
        cd $WORKPATH/game_$INDEX
        if [ -f $PIDFILE ]
        then
                echo "$PIDFILE 已经存在,游戏服($INDEX)已经启动或奔溃"
        else
                export GOTRACEBACK=crash
                echo "正在启动游戏服($INDEX)"
                nohup $EXEC --config $CONF >>/dev/null 2>&1 &
                PID=$!
                if [ -x /proc/${PID} ]
                then
                        echo "游戏服($INDEX)启动成功"
                        echo $PID>$PIDFILE
                else
                        echo "游戏服($INDEX)启动失败"
                fi       
        fi             
}

gameStop(){
        INDEX=$1
        EXEC=./game_$suffix$INDEX
        PIDFILE=./game_$INDEX.pid
        CONF=./config/game_$INDEX.json
        cd $WORKPATH/game_$INDEX
        if [ ! -f $PIDFILE ]
        then
                echo "$PIDFILE 不存在,游戏服($INDEX)已经关闭"
        else
                PID=$(cat $PIDFILE)
                echo "正在停止游戏服($INDEX)"
                kill $PID
                while [ -x /proc/${PID} ]
                do
                echo "等候游戏服($INDEX)关闭"
                sleep 1
                done
                rm $PIDFILE
                echo "游戏服($INDEX)关闭中"
        fi           
}

game(){
        INDEX=$2
        EXEC=./game_$suffix$INDEX
        PIDFILE=./game_$INDEX.pid
        CONF=./config/game_$INDEX.json
        cd $WORKPATH/game_$INDEX

        case "$1" in
        restart)
                gameStop $INDEX
                gameStart $INDEX
                ;;
        start)
                gameStart $INDEX
                ;;
        stop)
                gameStop $INDEX
                ;;
        status)
                if [ -f $PIDFILE ]       
                then
                        PID=$(cat $PIDFILE)
                        if [ -x /proc/${PID} ]
                        then
                                echo "游戏服($INDEX)正在运行"
                        else 
                                echo "游戏服($INDEX)已经停止"
                        fi
                else
                        echo "游戏服($INDEX)已经停止"
                fi
                ;;
        update)
                cp ../game ./game_$suffix$INDEX
                echo "游戏服($INDEX)更新成功"
                ;;
        upgrade)
                gameStop $INDEX
                cp ../game ./game_$suffix$INDEX
                echo "游戏服($INDEX)更新成功"
                gameStart $INDEX
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}

platformStart(){
        INDEX=$1
        EXEC=./platform_$suffix$INDEX
        PIDFILE=./platform_$INDEX.pid
        CONF=./config/platform_$INDEX.json
        cd $WORKPATH/platform_$INDEX

        if [ -f $PIDFILE ]
        then
                echo "$PIDFILE 已经存在,平台服($INDEX)已经启动或奔溃"
        else
                
                echo "正在启动平台服($INDEX)"
                nohup $EXEC --config $CONF >>/dev/null 2>&1 &
                PID=$!
                if [ -x /proc/${PID} ]
                then
                        echo "平台服($INDEX)启动成功"
                        echo $PID>$PIDFILE
                else
                        echo "平台服($INDEX)启动失败"
                fi       
        fi           
}

platformStop(){
        INDEX=$1
        EXEC=./platform_$suffix$INDEX
        PIDFILE=./platform_$INDEX.pid
        CONF=./config/platform_$INDEX.json
        cd $WORKPATH/platform_$INDEX
        if [ ! -f $PIDFILE ]
        then
                echo "$PIDFILE 不存在,平台服($INDEX)已经关闭"
        else
                PID=$(cat $PIDFILE)
                echo "正在停止平台服($INDEX)"
                kill $PID
                while [ -x /proc/${PID} ]
                do
                echo "等候平台服($INDEX)关闭"
                sleep 1
                done
                rm $PIDFILE
                echo "平台服($INDEX)关闭中"
        fi         
}

platform(){
        INDEX=$2
        EXEC=./platform_$suffix$INDEX
        PIDFILE=./platform_$INDEX.pid
        CONF=./config/platform_$INDEX.json
        cd $WORKPATH/platform_$INDEX

        case "$1" in
        start)
                platformStart $INDEX
                ;;
        stop)
                platformStop $INDEX
                ;;
        status)
                if [ -f $PIDFILE ]       
                then
                        PID=$(cat $PIDFILE)
                        if [ -x /proc/${PID} ]
                        then
                                echo "平台服($INDEX)正在运行"
                        else 
                                echo "平台服($INDEX)已经停止"
                        fi
                else
                        echo "平台服($INDEX)已经停止"
                fi
                ;;
        update)
                cp ../platform ./platform_$suffix$INDEX
                echo "平台服($INDEX)更新成功"
                ;;
        upgrade)
                platformStop $INDEX
                cp ../platform ./platform_$suffix$INDEX
                echo "平台服($INDEX)更新成功"
                platformStart $INDEX
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}




regionStart(){
        INDEX=$1
        EXEC=./region_$suffix$INDEX
        PIDFILE=./region_$INDEX.pid
        CONF=./config/region_$INDEX.json
        cd $WORKPATH/region_$INDEX

        if [ -f $PIDFILE ]
        then
                echo "$PIDFILE 已经存在,战区服($INDEX)已经启动或奔溃"
        else
                
                echo "正在启动战区服($INDEX)"
                nohup $EXEC --config $CONF >>/dev/null 2>&1 &
                PID=$!
                if [ -x /proc/${PID} ]
                then
                        echo "战区服($INDEX)启动成功"
                        echo $PID>$PIDFILE
                else
                        echo "战区服($INDEX)启动失败"
                fi       
        fi           
}

regionStop(){
        INDEX=$1
        EXEC=./region_$suffix$INDEX
        PIDFILE=./region_$INDEX.pid
        CONF=./config/region_$INDEX.json
        cd $WORKPATH/region_$INDEX
        if [ ! -f $PIDFILE ]
        then
                echo "$PIDFILE 不存在,战区服($INDEX)已经关闭"
        else
                PID=$(cat $PIDFILE)
                echo "正在停止战区服($INDEX)"
                kill $PID
                while [ -x /proc/${PID} ]
                do
                echo "等候战区服($INDEX)关闭"
                sleep 1
                done
                rm $PIDFILE
                echo "战区服($INDEX)关闭中"
        fi         
}

region(){
        INDEX=$2
        EXEC=./region_$suffix$INDEX
        PIDFILE=./region_$INDEX.pid
        CONF=./config/region_$INDEX.json
        cd $WORKPATH/region_$INDEX

        case "$1" in
        start)
                regionStart $INDEX
                ;;
        stop)
                regionStop $INDEX
                ;;
        status)
                if [ -f $PIDFILE ]       
                then
                        PID=$(cat $PIDFILE)
                        if [ -x /proc/${PID} ]
                        then
                                echo "战区服($INDEX)正在运行"
                        else 
                                echo "战区服($INDEX)已经停止"
                        fi
                else
                        echo "战区服($INDEX)已经停止"
                fi
                ;;
        update)
                cp ../region ./region_$suffix$INDEX
                echo "战区服($INDEX)更新成功"
                ;;
        upgrade)
                regionStop $INDEX
                cp ../region ./region_$suffix$INDEX
                echo "战区服($INDEX)更新成功"
                regionStart $INDEX
                ;;
        *)
                echo "Please use start or stop as first argument"
                ;;
        esac
}



# 开始,更新,状态,停止
case "$1" in
        trade_server)
                trade_server "$2"
                ;;
        coupon_server)
                coupon_server "$2"
                ;;
        charge_server)
                charge_server "$2"
                ;;
        center)
                center "$2"
                ;;
        account)
                account "$2"
                ;;
        cross)
                cross "$2"
                ;;
        game)
                case "$2" in 
                        pull)
                                echo "游戏服拉取"
                                rsync -azP --delete ../deploy/game/game ./
                                rsync -azP --delete ../deploy/resources/* ./resources/
                                chmod 777 game
                                ;;
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
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
                        pull)
                                echo "组服拉取"
                                rsync -azP --delete ../deploy/cross/cross ./group
                                rsync -azP --delete ../deploy/resources/* ./resources/
                                chmod 777 group
                                ;;
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for groupDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$groupDir" && $groupDir =~ group_[0-9]+ ]]
                                                then
                                                index=$(echo $groupDir | cut -d "_" -f2)
                                                echo "组服目录$groupDir,索引$index"
                                                serverList+=("$index")
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
        platform)
                case "$2" in 
                        pull)
                                echo "平台服拉取"
                                rsync -azP --delete ../deploy/cross/cross ./platform
                                rsync -azP --delete ../deploy/resources/* ./resources/
                                chmod 777 platform
                                ;;
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for platformDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$platformDir" && $platformDir =~ platform_[0-9]+ ]]
                                                then
                                                index=$(echo $platformDir | cut -d "_" -f2)
                                                echo "平台服目录$platformDir,索引$index"
                                                serverList+=("$index")
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
         region)
                case "$2" in 
                        pull)
                                echo "战区服拉取"
                                rsync -azP --delete ../deploy/cross/cross ./region
                                rsync -azP --delete ../deploy/resources/* ./resources/
                                chmod 777 region
                                ;;
                        *)
                                
                                # 获取剩余参数
                                serverList=()
                                if [ "$3" = "all" ]
                                then
                                        # 查找游戏服
                                        for regionDir in `ls $WORKPATH`
                                        do
                                                if [[ -d "$WORKPATH/$regionDir" && $regionDir =~ region_[0-9]+ ]]
                                                then
                                                index=$(echo $regionDir | cut -d "_" -f2)
                                                echo "战区服目录$regionDir,索引$index"
                                                serverList+=("$index")
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
        *)
                echo "请输入account,game参数"
                exit 1
                ;;
esac