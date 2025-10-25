/*
   Create: 2024/1/28
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package oauth_manager

import (
	"encoding/json"
	"fmt"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
)

type GithubUserInfo struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	HomeUrl   string `json:"html_url"`
}

type ApolloOAuthUser struct {
	Login       string `json:"login"`
	AvatarUrl   string `json:"avatarUrl"`
	HomeUrl     string `json:"homeUrl"`
	AccessToken string `json:"accessToken"`
}

func GetGithubUser(token string) (GithubUserInfo, error) {
	res, err := utils.HttpGet(GithubUserApi, map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", token),
	}, 0)
	if err != nil {
		return GithubUserInfo{}, err
	}
	var user GithubUserInfo
	_ = json.Unmarshal(res, &user)
	return user, err
}

func GetGithubOAuthUrl() string {
	return fmt.Sprintf(GithubOAuthApi, config.ApolloConf.Server.OAuth.ClientID)
}
