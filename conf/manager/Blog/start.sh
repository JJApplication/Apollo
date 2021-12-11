#!/usr/bin/env bash

nohup "${APP_ROOT}/${APP}/app_blog" web cluster -p 10002 > "${APP_LOG}/${APP}/${APP}.log" 2>&1 &
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi