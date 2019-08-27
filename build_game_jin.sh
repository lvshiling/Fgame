#!/bin/sh

SERVER_SVN="E:/project1/work1/server_exe"
GAME_SVN=$SERVER_SVN"/game/"

go install fgame/fgame/game
cp $GOPATH"/bin/game" $GAME_SVN
cp -r ./sql $SERVER_SVN"/sql"
