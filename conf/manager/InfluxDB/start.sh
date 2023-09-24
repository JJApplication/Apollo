#!/usr/bin/env bash

exist=$(docker ps -a|grep "${APP}")
if [ -z "${exist}" ];then
  exit 0
else
  exit 0
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi