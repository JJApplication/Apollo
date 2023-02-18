#!/usr/bin/env bash
cd "${APP_ROOT}/${APP}" || exit "${APP_START_ERR}"
if [[ ! -d "${APP_LOG}/${APP}" ]];then
  mkdir -p "${APP_LOG}/${APP}"
fi

export V2RAY_VMESS_AEAD_FORCED=false
nohup "${APP_ROOT}/${APP}/v2ray" run -config /etc/v2ray/config.json >> "${APP_LOG}/${APP}/${APP}.log" 2>&1 &
result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi