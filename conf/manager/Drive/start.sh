#!/usr/bin/env bash

exist=$(docker ps -a|grep "${APP}")
if [ -z "${exist}" ];then
  docker run -d --name="${APP}" -v /renj.io/drive/data:/opt/alist/data -p 127.0.0.1:5244:5244 -e PUID=0 -e PGID=0 -e UMASK=022 xhofe/alist:latest
else
  docker start "${APP}"
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi