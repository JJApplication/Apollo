#!/usr/bin/env bash

status=$(ps ax|grep -w "${APP}"|grep -v zeus|grep -v grep|grep -v .sh)
if [[ -n ${status} ]];then
  exit 0
else
  exit "${APP_STATUS_ERR}"
fi