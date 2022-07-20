// 数据库操作层 当前所有数据基于mongo保存和备份
// 备份数据格式为mongo支持的bson数据

package database

// 为保证随时可以插入新的数据和新表
// 使用mongoDB实现存储 超大json实现转储
