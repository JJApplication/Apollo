// Package router_env
package router_env

import (
	"github.com/JJApplication/Apollo/app/env_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

type reqSetEnv struct {
	Service string `json:"service"`
	Envs    []struct {
		Key     string `json:"key"`
		Value   string `json:"value"`
		Encrypt bool   `json:"encrypt"`
	} `json:"envs"`
}

func ListServices(c *gin.Context) {
	envManager := env_manager.GetEnvManager()
	router.Response(c, envManager.ListAllServices(), true)
}

func GetEnvs(c *gin.Context) {
	service := c.Query("service")
	if service == "" {
		router.Response(c, nil, false)
		return
	}
	envManager := env_manager.GetEnvManager()
	data, err := envManager.GetServiceConfigs(service)
	if err != nil {
		router.Response(c, err.Error(), false)
		return
	}
	router.Response(c, data, true)
}

func GetEnv(c *gin.Context) {
	service := c.Query("service")
	key := c.Query("key")
	if service == "" || key == "" {
		router.Response(c, nil, false)
		return
	}
	envManager := env_manager.GetEnvManager()
	data, err := envManager.GetConfig(service, key)
	if err != nil {
		router.Response(c, err.Error(), false)
		return
	}
	router.Response(c, data, true)
}

func GetEnvWithAES(c *gin.Context) {
	service := c.Query("service")
	key := c.Query("key")
	if service == "" || key == "" {
		router.Response(c, nil, false)
		return
	}
	envManager := env_manager.GetEnvManager()
	data, err := envManager.GetConfig(service, key)
	if err != nil {
		router.Response(c, err.Error(), false)
		return
	}

	if !data.Encrypt {
		router.Response(c, data.GetValue(), true)
		return
	}
	decrypt, err := utils.DecryptAES256(data.Value, config.ApolloConf.AES.Key)
	if err != nil {
		router.Response(c, err.Error(), false)
		return
	}
	router.Response(c, decrypt, true)
}

func SetEnv(c *gin.Context) {
	var req reqSetEnv
	if err := c.ShouldBindJSON(&req); err != nil {
		router.Response(c, err.Error(), false)
		return
	}
	envManager := env_manager.GetEnvManager()
	for _, env := range req.Envs {
		if err := envManager.SetConfig(req.Service, env.Key, env.Value, env.Encrypt); err != nil {
			logger.LoggerSugar.Errorf("SetEnv [%s]-[%s] err: %s", req.Service, env.Key, err.Error())
			continue
		}
	}

	router.Response(c, nil, true)
}

func DeleteEnv(c *gin.Context) {
	service := c.Query("service")
	key := c.Query("key")
	if service == "" {
		router.Response(c, nil, false)
		return
	}
	// 不存在key为全部删除
	envManager := env_manager.GetEnvManager()
	if key == "" {
		_, err := envManager.DeleteServiceConfigs(service)
		if err != nil {
			router.Response(c, err.Error(), false)
			return
		}
		router.Response(c, nil, true)
		return
	}
	if err := envManager.DeleteConfig(service, key); err != nil {
		router.Response(c, err.Error(), false)
	}
	router.Response(c, nil, true)
}
