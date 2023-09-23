/*
   Create: 2023/9/18
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

import (
	"encoding/json"
	"github.com/JJApplication/Apollo/config"
	"sync"
	"time"
)

// 同一时间只有一个token会生效
// 键为登录的IP地址
// 键作为历史登录IP使用， 当前登录生效后，将当前登录IP记录为active 其他置为passive

type TokenItem struct {
	LoginIP   string
	LoginUser string
	LoginTime int64
	Expire    int64 // 本地记录的失效值 需要和token中解密出来的一致再校验
	IsActive  bool
}

var tokenPool sync.Map

func GetTokenMap() []TokenItem {
	var res []TokenItem
	tokenPool.Range(func(key, value any) bool {
		res = append(res, value.(TokenItem))
		return true
	})

	return res
}

// GetToken 获取当前用户IP所属的token判断是否生效
// isActive为false时未生效， 此时需要重新登录
func GetToken(ip string) TokenItem {
	if t, ok := tokenPool.Load(ip); ok {
		return t.(TokenItem)
	}
	return TokenItem{}
}

func RemoveToken(ip string) {
	tokenPool.Delete(ip)
}

func GetActiveToken() TokenItem {
	var t TokenItem
	tokenPool.Range(func(key, value any) bool {
		if value.(TokenItem).IsActive {
			t = value.(TokenItem)
		}
		return true
	})

	return t
}

// SetToken 更新token
func SetToken(token Token) {
	// 其他token失效
	DisActiveAllToken()
	tokenPool.Store(token.LoginIP, TokenItem{
		LoginIP:   token.LoginIP,
		LoginUser: config.ApolloConf.Server.Account, // 当前不支持多用户
		LoginTime: token.Create,
		Expire:    token.Expire,
		IsActive:  true,
	})
}

// DisActiveAllToken 用户登出
func DisActiveAllToken() {
	var keys []string
	tokenPool.Range(func(key, value any) bool {
		keys = append(keys, key.(string))
		return true
	})

	for _, key := range keys {
		if t, ok := tokenPool.Load(key); ok {
			tt := t.(TokenItem)
			tt.IsActive = false
			tokenPool.Store(key, tt)
		}
	}
}

func ValidateToken(ip string, token string) bool {
	if t, ok := tokenPool.Load(ip); ok {
		cacheToken := t.(TokenItem)
		if !cacheToken.IsActive {
			return false
		}
		if ip != cacheToken.LoginIP {
			return false
		}
		var tt Token
		err := json.Unmarshal([]byte(DecryptDefault(token)), &tt)
		if err != nil {
			return false
		}
		// 判断是否过期
		if time.Now().Unix() >= tt.Expire {
			return false
		}
		// 判断IP是否一致
		if tt.LoginIP != ip {
			return false
		}
		// 判断随机值和value加密值是否匹配
		return validateValue(tt,
			config.ApolloConf.Server.Account+config.ApolloConf.Server.PassWd,
			cacheToken.LoginIP, cacheToken.LoginTime)
	}
	return false
}

func validateValue(t Token, value, loginIp string, loginTime int64) bool {
	return t.Value == EncryptDefault(value) && t.Random == GetRandom(loginIp, loginTime)
}
