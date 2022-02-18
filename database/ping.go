/*
Project: dirichlet ping.go
Created: 2021/12/25 by Landers
*/

package database

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

// ping测试

type PingCtx struct {
	mgm.Model
}

// WritePing 尝试写ping数据
func WritePing() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, _ = mgm.Coll(&PingCtx{}).InsertOne(ctx, PingCtx{})
}

// Ping 异步的连接返回 使用轮询尝试
func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := mgm.Coll(&PingCtx{}, nil).FindOne(ctx, bson.M{}).Err(); err != nil {
		// context上下文超时则返回异常
		select {
		case <-ctx.Done():
			if _, ok := ctx.Deadline(); ok {
				return errors.New("error ping mongo")
			}
		}
		if strings.Contains(err.Error(), "connection refused") {
			return errors.New("error ping mongo")
		}
	}
	return nil
}
