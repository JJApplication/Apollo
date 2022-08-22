/*
Project: Apollo open_mongo.go
Created: 2021/11/30 by Landers
*/

package database

import (
	"fmt"
	"time"

	"github.com/JJApplication/Apollo/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	Name = "ApolloMongo"
)

var DBName = config.ApolloConf.DB.Mongo.Name

const (
	MongoPrefix = "mongodb://"
)

// MongoPing 当mongo无法连接时不应该阻塞后续数据库操作，应该直接跳过
var MongoPing bool

// InitDBMongo 连接mongo
func InitDBMongo() error {
	if DBName == "" {
		DBName = Name
	}
	MongoPing = true
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 2 * time.Second}, DBName, options.Client().ApplyURI(ParseDSN(config.ApolloConf.DB.Mongo)))
	if err != nil {
		return err
	}
	WritePing()
	return nil
}

// ParseDSN 解析配置拼接mongo URI
func ParseDSN(conf config.Mongo) string {
	if conf.User == "" || conf.PassWd == "" {
		return fmt.Sprintf("%s%s", MongoPrefix, conf.URL)
	}
	return fmt.Sprintf("%s%s:%s@%s", MongoPrefix, conf.User, conf.PassWd, conf.URL)
}
