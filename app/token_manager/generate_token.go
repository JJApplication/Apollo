/*
   Create: 2023/9/17
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package token_manager

import "github.com/JJApplication/Apollo/config"

// Token生成 基于时间戳和对称算法
//
// Token结构为JSON -> crypto -> byte

func GenerateToken(loginIp string, loginTime int64) Token {
	return generateToken(loginIp, loginTime,
		int64(config.ApolloConf.Server.AuthExpire),
		config.ApolloConf.Server.Account+config.ApolloConf.Server.PassWd)
}

func generateToken(loginIp string, loginTime int64, expireDuration int64, value string) Token {
	return Token{
		Value:   EncryptDefault(value),
		Random:  GetRandom(loginIp, loginTime),
		Create:  loginTime,
		Expire:  loginTime + expireDuration,
		LoginIP: loginIp,
	}
}
