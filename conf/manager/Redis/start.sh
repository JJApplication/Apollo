#!/usr/bin/env bash

/usr/local/redis/bin/redis-server /usr/local/redis/6379.conf
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi