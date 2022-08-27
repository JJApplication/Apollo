#!/usr/bin/env bash

docker stop mongo
result=$?
if [[ $result != 0 ]];then
    exit "${APP_STOP_ERR}"
else
    exit 0
fi