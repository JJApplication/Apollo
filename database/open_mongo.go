/*
Project: dirichlet open_mongo.go
Created: 2021/11/30 by Landers
*/

package database

import (
	"github.com/kamva/mgm/v3"
	"github.com/landers1037/dirichlet/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DBName = "DirichletMongo"
)

// InitDBMongo 连接mongo
func InitDBMongo() error {
	return mgm.SetDefaultConfig(nil, DBName, options.Client().ApplyURI(config.DirichletConf.DB.Mongo.URL))
}
