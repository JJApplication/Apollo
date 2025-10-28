package database_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils/json"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

type Database struct {
	db   *gorm.DB
	Node Node  `json:"node"`
	Info *Info `json:"info"`
}

type Node struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"` // 数据库类型
}

// Info 数据库信息
type Info struct {
	DBSize     int64 `json:"db_size"`
	CreateTime int64 `json:"create_time"`
	UpdateTime int64 `json:"update_time"`
	// 数据库详细信息
	Table []Table `json:"table"`
}

type Table struct {
	Name        string   `json:"name"`         // 表名
	Columns     []string `json:"columns"`      // 列名
	Rows        int64    `json:"rows"`         // 行数
	BusyTimeout int64    `json:"busy_timeout"` // busy超时
	Encoding    string   `json:"encoding"`     // 编码
	Synchronous int64    `json:"synchronous"`  // 同步模式
}

// LoadDatabaseNode 从节点文件中加载数据库配置
func LoadDatabaseNode() ([]Database, error) {
	data, err := os.ReadFile(filepath.Join(config.GlobalConfigRoot, DBNodeFile))
	if err != nil {
		return nil, err
	}
	var nodes []Node
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		return nil, err
	}

	dbs := make([]Database, len(nodes))
	for i, n := range nodes {
		dbs[i].Node = n
	}

	return dbs, nil
}
