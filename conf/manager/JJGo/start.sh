#!/usr/bin/env bash
cd "${APP_ROOT}/${APP}" || exit "${APP_START_ERR}"
if [[ ! -d "${APP_LOG}/${APP}" ]];then
  mkdir -p "${APP_LOG}/${APP}"
fi

conf="${APP_ROOT}/${APP}/conf/jjgo.ini"
conf_temp="${APP_ROOT}/${APP}/conf/jjgo.ini.env"
envsubst < "$conf_temp" > "$conf"

nohup "${APP_ROOT}/${APP}/app_jjgo" > "${APP_LOG}/${APP}/${APP}.log" 2>&1 &
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi