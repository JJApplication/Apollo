#!/usr/bin/env bash

cd "${APP_ROOT}/${APP}" || exit "${APP_START_ERR}"
if [[ ! -d "${APP_LOG}/${APP}" ]];then
  mkdir -p "${APP_LOG}/${APP}"
fi

exist=$(docker ps -a|grep "${APP}")
if [ -z "${exist}" ];then
  bash "${APP_ROOT}/${APP}/run.sh"
else
  docker start "${APP}"
fi
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi