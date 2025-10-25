package router_panel

import (
	"encoding/json"
	"fmt"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/router"
	"github.com/JJApplication/Apollo/utils"
	"github.com/gin-gonic/gin"
)

// 网络流量面板

const (
	ServiceNetwork = "Taco"
	APIStat        = "/api/stat"
	APIDomainStat  = "/api/domain"
	APIGeo         = "/api/geo"
)

func NetworkStat(c *gin.Context) {
	// 透传客户端IP
	clientIp := c.GetHeader("X-Real-Ip")
	header := map[string]string{
		"X-Real-Ip": clientIp,
	}
	Addr := config.ApolloConf.HttpLocal.GetAddr(ServiceNetwork)
	timeout := config.ApolloConf.HttpLocal.Timeout
	data, err := utils.HttpGet(fmt.Sprintf("http://127.0.0.1:%d%s", Addr, APIStat), header, timeout)
	if err != nil {
		router.Response(c, nil, false)
		return
	}
	// 转为JSON数据
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	fmt.Println(result)
	router.Response(c, result, true)
}

func NetworkDomain(c *gin.Context) {
	// 透传客户端IP
	clientIp := c.GetHeader("X-Real-Ip")
	header := map[string]string{
		"X-Real-Ip": clientIp,
	}
	Addr := config.ApolloConf.HttpLocal.GetAddr(ServiceNetwork)
	timeout := config.ApolloConf.HttpLocal.Timeout
	data, err := utils.HttpGet(fmt.Sprintf("http://127.0.0.1:%d%s", Addr, APIDomainStat), header, timeout)
	if err != nil {
		router.Response(c, nil, false)
		return
	}
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	router.Response(c, result, true)
}

func NetworkGeoIP(c *gin.Context) {
	// 透传客户端IP
	clientIp := c.GetHeader("X-Real-Ip")
	header := map[string]string{
		"X-Real-Ip": clientIp,
	}
	Addr := config.ApolloConf.HttpLocal.GetAddr(ServiceNetwork)
	timeout := config.ApolloConf.HttpLocal.Timeout
	data, err := utils.HttpGet(fmt.Sprintf("http://127.0.0.1:%d%s", Addr, APIGeo), header, timeout)
	if err != nil {
		router.Response(c, nil, false)
		return
	}
	var result map[string]interface{}
	json.Unmarshal(data, &result)
	router.Response(c, result, true)
}
