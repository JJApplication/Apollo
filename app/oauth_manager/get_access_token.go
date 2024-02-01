/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package oauth_manager

import (
	"encoding/json"
	"errors"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/utils"
	"strings"
)

// 获取access_token

type accessTokenBody struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	RedirectUrl  string `json:"redirect_uri"`
}

type accessTokenRes struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

// GetAccessToken 从配置文件中读取client信息
func GetAccessToken(code string) (string, error) {
	clientId := config.ApolloConf.Server.OAuth.ClientID
	clientSecret := config.ApolloConf.Server.OAuth.ClientSecret

	body, err := json.Marshal(accessTokenBody{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		Code:         code,
		//RedirectUrl:  ApolloRedirectUrl,
	})

	logger.LoggerSugar.Infof("%s try to get github access token with [%s]", OAuthManager, clientId)

	if err != nil {
		logger.LoggerSugar.Errorf("%s get github access token error: %s", OAuthManager, err.Error())
		return "", err
	}
	data, err := utils.HttpPost(GithubAccessToken, body)
	if err != nil {
		logger.LoggerSugar.Errorf("%s get github access token error: %s", OAuthManager, err.Error())
		return "", err
	}
	if strings.Contains(string(data), "error") {
		logger.LoggerSugar.Errorf("%s get github access token error: %s", OAuthManager, data)
		return "", errors.New(string(data))
	}

	var tokenRes accessTokenRes
	_ = json.Unmarshal(data, &tokenRes)

	return tokenRes.AccessToken, nil
}
