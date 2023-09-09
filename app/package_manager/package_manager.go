/*
   Create: 2023/6/16
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package package_manager
// 包管理器 针对当前集成到apollo的部署服务包
// 所有的服务包使用统一的二进制.jp [jjapp package]二进制格式打包为单文件
// 二进制的头256字节为校验码用于检查包完整性
package package_manager
