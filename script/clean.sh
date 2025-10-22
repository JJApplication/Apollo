#!/usr/bin/env bash

#################################
# 清理日志路径下的tar.gz无用日志
#################################

WORK_DIR=/renj.io/log
SCRIPT_NAME=clean.sh

if [ ! -d $WORK_DIR ];then
  exit 1
fi

cd $WORK_DIR || exit 1

dirs=$(find $WORK_DIR -type d)

for dir in $dirs;do
  if [ -z "$dir" ];then
    continue
  fi
  logs=$(find "$dir" -type f -name "*.gz")
  for log in $logs;do
    echo "clear $log"
    rm -f "$log"
  done
done
