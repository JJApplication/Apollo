#!/usr/bin/env bash

for i in {1..10}
do
  pid=$(ps ax|grep "${APP}"|grep -v grep|grep -v zeus|grep -v ".sh"|sed -n '1P'|awk '{print $1}')
  if [[ -z "$pid" ]];then
    break
  else
    kill -9 "$pid" > /dev/null
    result=$?
    if [[ $result != 0 ]];then
      exit "${APP_STOP_ERR}"
    else
      exit 0
    fi
  fi
done