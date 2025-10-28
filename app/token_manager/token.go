/*
   Create: 2023/9/18
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

import (
	"github.com/JJApplication/Apollo/utils"
	"github.com/JJApplication/Apollo/utils/json"
)

// Token 基于IP控制 单点登录 每个IP的token有效期单独维护
type Token struct {
	Value   string `json:"value"`   // 密钥
	Random  string `json:"random"`  // 随机值(基于创建时间+IP生成绝对唯一)
	Create  int64  `json:"create"`  // 创建时间
	Expire  int64  `json:"expire"`  // 失效时间
	LoginIP string `json:"loginIP"` // 登入IP
}

// 返回的是处理后的json加密字符串
// 基于default加密算法
func (t *Token) String() string {
	return EncryptDefault(utils.JsonString(t))
}

func DecryptToken(text string) Token {
	code := DecryptDefault(text)
	var t Token
	_ = json.Unmarshal([]byte(code), &t)

	return t
}
