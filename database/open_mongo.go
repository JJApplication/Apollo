/*
Project: dirichlet open_mongo.go
Created: 2021/11/30 by Landers
*/

package database

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/landers1037/dirichlet/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName = "DirichletMongo"
)

// InitDBMongo 连接mongo
func InitDBMongo() error {
	return mgm.SetDefaultConfig(&mgm.Config{CtxTimeout: 10 * time.Second}, DBName, options.Client().ApplyURI(config.DirichletConf.DB.Mongo.URL))
}
