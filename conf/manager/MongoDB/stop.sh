#!/usr/bin/env bash

APP=mongo
docker stop "${APP}"
result=$?
if [[ $result != 0 ]];then
    exit "${APP_STOP_ERR}"
else
    exit 0
fi