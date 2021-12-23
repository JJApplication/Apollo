#!/usr/bin/env bash

# 因为是定时脚本由cron接管所以，只需要判断cron中是否存在规则
result=$(crontab -l|grep jjbot.sh)
if [[ -z "$result" ]];then
  exit "${APP_START_ERR}"
fi
exit 0
