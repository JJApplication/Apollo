/*
   Create: 2023/9/18
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

import (
	"github.com/JJApplication/Apollo/utils"
	"github.com/google/uuid"
)

// 获取随机的UUID

func UUID() string {
	return uuid.NewString()
}

// GetRandom 获取本地登录的随机值
func GetRandom(loginIp string, loginTime int64) string {
	return hashMap[HMAC_SHA256](loginIp + utils.GetTimeString(loginTime))
}
