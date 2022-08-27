#!/usr/bin/env bash

exist=$(docker ps -a|grep redis)
if [ -z "${exist}" ];then
  docker run -d --name=redis -v "${APP_ROOT}/${APP}"/6379.conf:/etc/redis/redis.conf -p 127.0.0.1:6379:6379 redis
else
  docker start redis
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi