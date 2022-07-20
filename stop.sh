#!/usr/bin/env bash
run_path=$(pwd)
app=apollo
executable="${run_path}/${app}"

check()
{
  result=$(ps ax|grep ${executable}|grep -v grep|awk '{print $1}')
  if [ -z $result ];then
    exit 0
  else
    kill -9 $result
    exit 0
  fi
}

check