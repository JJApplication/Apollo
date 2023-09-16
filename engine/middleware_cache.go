/*
Create: 2022/8/29
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package engine
package engine

import (
	"github.com/JJApplication/Apollo/config"
	"time"

	"github.com/gin-contrib/cache"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

// 缓存静态文件

const (
	DefaultExpire = time.Hour * 24 * 7
	PageExpire    = time.Hour * 24
)

var store = persistence.NewInMemoryStore(DefaultExpire)

func GetStore() *persistence.InMemoryStore {
	return store
}

func MiddleCache(handle gin.HandlerFunc) gin.HandlerFunc {
	if config.ApolloConf.Server.UICache {
		if config.ApolloConf.Server.UICacheTime > 0 {
			return cache.CachePage(store, time.Duration(config.ApolloConf.Server.UICacheTime), handle)
		}
		return cache.CachePage(store, PageExpire, handle)
	}
	return handle
}
