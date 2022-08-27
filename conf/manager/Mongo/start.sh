#!/usr/bin/env bash

exist=$(docker ps -a|grep mongo)
if [ -z "${exist}" ];then
  docker run -d --name=mongo -p 127.0.0.1:27017:27017 mongo
else
  docker start mongo
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi