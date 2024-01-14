/*
   Create: 2023/7/19
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package noengine_manager 负责nginx进程或容器的管理
//
// 日志读取 清除
// 配置文件读取 修改 删除
// 缓存清除
// MIME修改
//
// NoEngine Manager对NoEngine原有启动方式全部重构
// 新的web容器直接基于一个单独得openresty容器启动 http服务监听在80端口 http服务监听在443端口
// root路径统一为/app
package noengine_manager
