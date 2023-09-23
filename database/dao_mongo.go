/*
Project: Apollo dao_sqlite.go
Created: 2021/12/2 by Landers
*/

package database

// 基于xorm的sqlite备份恢复数据库
// 备份数据库 为主动方式恢复 不会与mongo冲突
//
// 提供最基础的关系性数据库操作能力 | 查询表 列出表 查看表结构 删除表 删除表内容
//
// 其他数据库操作基于exec直接输入sql操作
