#!/usr/bin/env bash

nohup "${APP_ROOT}/${APP}/frps" -c "${APP_ROOT}/${APP}/frps.ini" > "${APP_LOG}/${APP}/${APP}.log" 2>&1 &
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi