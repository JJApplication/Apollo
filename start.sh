#!/usr/bin/env bash
run_path=$(cd $(dirname $0);pwd)
app=apollo
executable="${run_path}/${app}"

check()
{
  result=$(ps ax|grep "${executable}"|grep -v grep|awk '{print $1}')
  if [ -z "$result" ];then
    nohup "${executable}" > /dev/null 2>&1 &
    if [ $? != 0 ];then
      exit 1
    fi
    exit 0
  fi
  exit 0
}

check