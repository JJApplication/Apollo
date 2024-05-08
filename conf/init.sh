#!/bin/bash
# 重启环境时，启动全部apollo服务

# todo按需启动

APP_ROOT=/renj.io/app
echo "JJApp INIT LOG - $(date)"

# 检查依赖服务
# 依赖mongodb和NoEngine服务的启动

##############################
# NoEngine
# 保证web服务入口启动
##############################
result=$(docker ps|grep NoEngine)
if [ -n "$result" ];then
  echo "NoEngine is alive"
  exit 0
fi

echo "try to start NoEngine"
docker start NoEngine

##############################
# MongoDB
# 保证mongo启动，后续的微服务状态更新依赖
##############################
result=$(docker ps|grep MongoDB)
if [ -n "$result" ];then
  echo "MongoDB is alive"
  exit 0
fi

echo "try to start MongoDB"
docker start MongoDB

##############################
# Apollo
# 启动Apollo管理进程
##############################
result=$(ps ax|grep Apollo|grep -v grep)
if [ -n "$result" ];then
  echo "Apollo is alive"
  exit 0
fi

echo "try to start Apollo"
cd $APP_ROOT/Apollo || exit 1
bash ./start.sh
res=$?

if [ $res != 0 ];then
  echo "Apollo start failed"
  exit 1
fi


