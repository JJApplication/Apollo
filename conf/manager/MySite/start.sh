#!/usr/bin/env bash

nohup gunicorn -c "${APP_ROOT}/${APP}/gun_mysite.py" app_mysite:app > "${APP_LOG}/${APP}/${APP}.log" 2>&1 &
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi