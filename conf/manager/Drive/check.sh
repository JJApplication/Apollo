#!/usr/bin/env bash

exist=$(docker ps|grep "${APP}")
if [ -n "${exist}" ];then
  exit 0
else
  exit "${APP_STATUS_ERR}"
fi