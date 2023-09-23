/*
Project: Apollo dao_sqlite_test.go
Created: 2021/12/2 by Landers
*/

package database

import (
	"context"
	"testing"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	mgm.SetDefaultConfig(nil, DBName, options.Client().ApplyURI("mongodb://192.168.10.128:27017"))
}

type tmpTest struct {
	mgm.DefaultModel
	Name string `bson:"name"`
}

func TestDao(t *testing.T) {
	var e error
	e = mgm.Coll(&tmpTest{}).Create(&tmpTest{Name: "abc"})

	t.Log("create data", e)

	var res tmpTest

	e = mgm.Coll(&tmpTest{}).FindOne(context.Background(), bson.M{"name": "abc"}).Decode(&res)

	t.Log("find data", e)
	t.Logf("%+v", res)
}
