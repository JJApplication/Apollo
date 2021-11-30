/*
Project: dirichlet model_base.go
Created: 2021/12/1 by Landers
*/

package database

import (
	"github.com/kamva/mgm/v3"
)

// Base 基础表结构
type Base struct {
	mgm.DefaultModel
	Table string `json:"table" bson:"table"`
}
