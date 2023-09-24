#!/usr/bin/env bash

exist=$(docker ps -a|grep "${APP}")
if [ -z "${exist}" ];then
  docker run -d --name="${APP}" -p 127.0.0.1:27017:27017 mongo
else
  docker start "${APP}"
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi