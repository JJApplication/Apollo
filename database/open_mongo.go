/*
Project: Apollo open_mongo.go
Created: 2021/11/30 by Landers
*/

package database

import (
	"time"

	"github.com/JJApplication/Apollo/config"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName = "ApolloMongo"
)

// MongoPing 当mongo无法连接时不应该阻塞后续数据库操作，应该直接跳过
var MongoPing bool

// InitDBMongo 连接mongo
func InitDBMongo() error {
	MongoPing = true
	err := mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 2 * time.Second}, DBName, options.Client().ApplyURI(config.ApolloConf.DB.Mongo.URL))
	if err != nil {
		return err
	}
	WritePing()
	return nil
}
