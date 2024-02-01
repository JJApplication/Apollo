/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package oauth_manager

import (
	"github.com/JJApplication/Apollo/config"
	"sync"
	"time"
)

const (
	OAuthManager = "[OAuth Manager] "
	OAuthExpire  = 3600
)

// 不使用github的access_token有效期
// 一旦登录完成默认有效期为1h

// OAuthUser
// 存储github的用户名
// access token 登录时间 主页地址 头像地址
type OAuthUser struct {
	Username  string `json:"username"`
	Token     string `json:"token"`
	HomeUrl   string `json:"homeUrl"`
	Avatar    string `json:"avatar"`
	LoginTime int64  `json:"loginTime"`
}

var OAuthManagerMap = sync.Map{}

// AddOAuthUser 登录OAuth用户
func AddOAuthUser(user OAuthUser) {
	if user.Username == "" {
		return
	}
	OAuthManagerMap.Store(user.Username, user)
}

// HasOAuthUser 是否存在此用户
func HasOAuthUser(name string) bool {
	_, ok := OAuthManagerMap.Load(name)
	return ok
}

// ValidateOAuth 校验OAuth用户accessToken
// 超过内置的token有效期后退出
func ValidateOAuth(token string) OAuthUser {
	var ou OAuthUser
	OAuthManagerMap.Range(func(name, info any) bool {
		user := info.(OAuthUser)
		if user.Token == token {
			ou = user
			return false
		}
		return true
	})

	// 失效返回空的用户
	if ExpireOAuthUser(ou.Username) {
		return OAuthUser{}
	}

	return ou
}

// ValidateOAuthAdmin 校验OAuth用户是否有最高管理权限
func ValidateOAuthAdmin(name string) bool {
	if name == "" {
		return false
	}
	for _, allow := range config.ApolloConf.Server.OAuth.AuthorizeList {
		if allow == name {
			return true
		}
	}

	return false
}

// RemoveOAuthUser 删除OAuth用户
func RemoveOAuthUser(token string) {
	user := GetOAuthUser(token)
	OAuthManagerMap.Delete(user.Username)
}

// GetOAuthUser 获取OAuth用户
func GetOAuthUser(token string) OAuthUser {
	user := ValidateOAuth(token)
	if user.Username == "" {
		return OAuthUser{}
	}
	user.Token = ""
	return user
}

// ExpireOAuthUser 查询用户是否失效
func ExpireOAuthUser(name string) bool {
	user, ok := OAuthManagerMap.Load(name)
	if !ok {
		return true
	}

	u := user.(OAuthUser)
	if (time.Now().Unix() - u.LoginTime) > OAuthExpire {
		return true
	}
	return false
}

// SyncFromGithub 从github同步oauth状态
// 查询状态为是否过期
func SyncFromGithub(name string) bool {
	user, ok := OAuthManagerMap.Load(name)
	if !ok {
		return true
	}
	info, err := GetGithubUser(user.(OAuthUser).Token)
	if err != nil {
		return true
	}
	if info.Login == "" {
		return true
	}

	return false
}
