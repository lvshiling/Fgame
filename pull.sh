#!/bin/sh
ip="192.168.0.128"
port="20000"
rsync -azP --delete -e "ssh -p $port"  root@$ip:~/fgame/deploy/ ./deploy/