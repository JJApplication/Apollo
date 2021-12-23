#!/usr/bin/env bash

cd "${APP_ROOT}/${APP}" || exit "${APP_START_ERR}"
if [[ ! -d "${APP_LOG}/${APP}" ]];then
  mkdir -p "${APP_LOG}/${APP}"
fi

# 因为是定时脚本由cron接管所以，只需要判断cron中是否存在规则
result=$(crontab -l|grep jjbot.sh)
if [[ -z "$result" ]];then
  echo "0 8 * * * /renj.io/apps/jjbot/jjbot.sh > /dev/null" >> /var/spool/cron/crontabs/root
  systemctl restart cron
fi

result=$?
if [[ $result != 0 ]];then
  exit "${APP_START_ERR}"
else
  exit 0
fi