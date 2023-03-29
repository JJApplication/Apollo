#!/usr/bin/env bash

exist=$(docker ps -a|grep "${APP}")
if [ -z "${exist}" ];then
  docker run -d --name="${APP}" -v "${APP_ROOT}/${APP}":/var/opt/memos -p 127.0.0.1:5230:5230 neosmemo/memos:latest
else
  docker start "${APP}"
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi