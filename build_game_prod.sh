#!/bin/sh

SERVER_SVN="/e/project1/branch/main/server_exe"
GAME_SVN=$SERVER_SVN"/game/"

go install fgame/fgame/game
cp $GOPATH"/bin/game" $GAME_SVN
cp -r ./sql $SERVER_SVN"/sql"
