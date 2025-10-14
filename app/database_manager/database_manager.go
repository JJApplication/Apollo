/*
   Create: 2023/6/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package database_manager
// 全局的数据库管理
// 对于当前服务器上的全部sqlite3数据库非认证数据库进行统一的数据管理和迁移
// 需要维护全局数据库文件路径名单：conf/database.node.json
package database_manager

import (
	"errors"
	"fmt"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

const (
	ManagerPrefix = "[DatabaseManager]"
	DBNodeFile    = "database.node.json"
)

var dbm *DatabaseManager

// DatabaseManager 数据库管理器的连接都是非持久化的 方便资源的控制
type DatabaseManager struct {
	lock      sync.RWMutex
	Databases map[string]*Database
}

func InitDatabaseManager() {
	sync.OnceFunc(func() {
		dbm = NewDatabaseManager()
	})()
}

func NewDatabaseManager() *DatabaseManager {
	dbs, err := LoadDatabaseNode()
	if err != nil {
		logger.LoggerSugar.Errorf("%s load database node file error: %v", ManagerPrefix, err)
		return nil
	}
	newDB := new(DatabaseManager)

	newDB.lock.Lock()
	defer newDB.lock.Unlock()
	newDB.Databases = make(map[string]*Database)
	for _, db := range dbs {
		newDB.Databases[db.Node.Name] = &db
	}

	return newDB
}

func Get() *DatabaseManager {
	if dbm == nil {
		InitDatabaseManager()
	}
	return dbm
}

// 打开数据库连接 如果已经连接则忽略
func (d *DatabaseManager) connect(name string) error {
	d.lock.Lock()
	defer d.lock.Unlock()
	if db, ok := d.Databases[name]; ok {
		if db.db != nil {
			return nil
		}
		g, err := gorm.Open(sqlite.Open(db.Node.Path), &gorm.Config{})
		if err != nil {
			return err
		}
		db.db = g
	}

	return errors.New(ErrorDBNotExist)
}

func (d *DatabaseManager) disconnect(name string) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	if db, ok := d.Databases[name]; ok {
		if db.db != nil {
			sqlDB, err := db.db.DB()
			if err != nil {
				return err
			}
			return sqlDB.Close()
		}
	}

	return errors.New(ErrorDBNotExist)
}

// List 获取节点信息
func (d *DatabaseManager) List() []Node {
	d.lock.RLock()
	defer d.lock.RUnlock()

	var nodes []Node
	for _, db := range d.Databases {
		nodes = append(nodes, db.Node)
	}

	return nodes
}

func (d *DatabaseManager) DB(name string) (*Database, error) {
	if db, ok := d.Databases[name]; ok {
		return db, nil
	}
	return nil, errors.New(ErrorDBNotExist)
}

func (d *DatabaseManager) Get(name string) *Info {
	db, err := d.DB(name)
	if err != nil {
		return nil
	}
	if db.Info != nil {
		return db.Info
	}

	if err = d.connect(name); err != nil {
		return nil
	}
	d.lock.RLock()
	defer d.lock.RUnlock()

	var info = new(Info)
	fileSize := utils.GetFileSize(db.Node.Path)
	info.DBSize = fileSize

	type tb struct {
		Name string `gorm:"column:name"`
	}
	var tables []tb
	db.db.Raw("SELECT name FROM sqlite_master WHERE type='table'").Scan(&tables)
	tablesData := make([]Table, len(tables))

	for i, table := range tables {
		var count int64
		var busytimeout int64
		var encoding string
		var synchronous int64
		var cols []string
		db.db.Raw(fmt.Sprintf("SELECT COUNT(1) FROM %s", table.Name)).Count(&count)
		db.db.Raw(fmt.Sprintf("SELECT name FROM PRAGMA_table_info('%s')", table.Name)).Find(&cols)
		db.db.Raw("PRAGMA busy_timeout").Scan(&busytimeout)
		db.db.Raw("PRAGMA encoding").Scan(&encoding)
		db.db.Raw("PRAGMA synchronous").Scan(&synchronous)
		tablesData[i] = Table{
			Name:        table.Name,
			Columns:     cols,
			Rows:        count,
			BusyTimeout: busytimeout,
			Encoding:    encoding,
			Synchronous: synchronous,
		}
	}
	info.Table = tablesData

	create, update := utils.GetFileTimes(db.Node.Path)
	info.CreateTime = create
	info.UpdateTime = update

	db.Info = info

	return info
}

// Refresh 刷新数据 刷新的是更新时间和统计信息
func (d *DatabaseManager) Refresh(name string) error {
	// 清空info即可
	d.lock.Lock()
	defer d.lock.Unlock()

	db, err := d.DB(name)
	if err != nil {
		return err
	}
	db.Info = nil

	return nil
}
