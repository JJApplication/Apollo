/*
Project: dirichlet bson_parser.go
Created: 2021/12/4 by Landers
*/

package utils

import (
	"io/ioutil"

	"go.mongodb.org/mongo-driver/bson"
)

// bson格式支持

func SaveBson(data interface{}, file string) error {
	_, b, err := bson.MarshalValue(data)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(file, b, 0644)
}

func ParseBson(file string, r interface{}) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return bson.Unmarshal(b, r)
}
