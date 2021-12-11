#!/usr/bin/env bash

status=$(ps ax|grep "${APP}"|grep -v zeus|grep -v grep)
if [[ -n ${status} ]];then
  exit 0
else
  exit "${APP_STATUS_ERR}"
fi