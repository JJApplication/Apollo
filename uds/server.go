/*
Create: 2022/8/26
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

// Package uds
package uds

import (
	"strings"

	"github.com/JJApplication/Apollo/app/app_manager"
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/fushin/server/uds"
	"github.com/JJApplication/fushin/utils/files"
	"github.com/JJApplication/fushin/utils/json"
)

const (
	UdsName = "Apollo"
)

var udsServer *uds.UDSServer

func initServer() {
	udsServer = uds.Default(getSocket())
	udsServer.Option.AutoCheck = false
	udsServer.Option.MaxSize = 5 << 20
	udsServer.Logger = &udsLogger{logger: logger.Logger}
}

func run() {
	initServer()
	udsServer.AddFunc("ping", func(c *uds.UDSContext, req uds.Req) {
		c.Response(uds.Res{
			Error: "",
			Data:  "",
			From:  UdsName,
			To:    nil,
		})
	})
	udsServer.AddFunc("app", func(c *uds.UDSContext, req uds.Req) {
		// data即app名称
		app, err := app_manager.GetApp(req.Data)
		if err != nil {
			c.Response(uds.Res{
				Error: err.Error(),
				Data:  "",
				From:  "",
				To:    nil,
			})
			return
		}
		data, _ := json.JSON.MarshalToString(app)
		c.Response(uds.Res{
			Error: "",
			Data:  data,
			From:  "",
			To:    nil,
		})
	})

	udsServer.AddFunc("app-all", func(c *uds.UDSContext, req uds.Req) {
		apps, err := app_manager.GetAllAppName()
		if err != nil {
			c.Response(uds.Res{
				Error: err.Error(),
				Data:  "",
				From:  "",
				To:    nil,
			})
			return
		}
		c.Response(uds.Res{
			Error: "",
			Data:  strings.Join(apps, " "),
			From:  "",
			To:    nil,
		})
		return
	})

	udsServer.AddFunc("start", func(c *uds.UDSContext, req uds.Req) {
		ok, err := app_manager.Start(req.Data)
		if ok {
			c.Response(uds.Res{})
			return
		}
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	udsServer.AddFunc("start-all", func(c *uds.UDSContext, req uds.Req) {
		res, err := app_manager.StartAll()
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{Data: strings.Join(res, ",")})
	})

	udsServer.AddFunc("stop", func(c *uds.UDSContext, req uds.Req) {
		ok, err := app_manager.Stop(req.Data)
		if ok {
			c.Response(uds.Res{})
			return
		}
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	udsServer.AddFunc("stop-all", func(c *uds.UDSContext, req uds.Req) {
		res, err := app_manager.StopAll()
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{Data: strings.Join(res, ",")})
	})

	udsServer.AddFunc("restart", func(c *uds.UDSContext, req uds.Req) {
		ok, err := app_manager.ReStart(req.Data)
		if ok {
			c.Response(uds.Res{})
			return
		}
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	udsServer.AddFunc("restart-all", func(c *uds.UDSContext, req uds.Req) {
		c.Response(uds.Res{Error: "restart all not support"})
	})

	udsServer.AddFunc("status", func(c *uds.UDSContext, req uds.Req) {
		res, err := app_manager.Status(req.Data)
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{Data: res})
	})

	udsServer.AddFunc("status-all", func(c *uds.UDSContext, req uds.Req) {
		res, err := app_manager.StatusAll()
		if err != nil {
			c.Response(uds.Res{Error: err.Error(), Data: strings.Join(res, ",")})
			return
		}
		c.Response(uds.Res{Data: strings.Join(res, ",")})
	})

	udsServer.AddFunc("sync", func(c *uds.UDSContext, req uds.Req) {
		_, err := app_manager.SyncApp(req.Data)
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	udsServer.AddFunc("sync-all", func(c *uds.UDSContext, req uds.Req) {
		err := app_manager.SyncAll()
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	udsServer.AddFunc("reload", func(c *uds.UDSContext, req uds.Req) {
		err := app_manager.ReloadManagerMap()
		if err != nil {
			c.Response(uds.Res{Error: err.Error()})
			return
		}
		c.Response(uds.Res{})
	})

	// 异步接口
	udsServer.AddFunc("backup", func(c *uds.UDSContext, req uds.Req) {
		go func() {
			if err := files.RsyncAndTar(config.ApolloConf.APPRoot, config.ApolloConf.APPBackUp, true); err != nil {
				logger.LoggerSugar.Errorf("run backup from uds error: %s", err.Error())
			}
		}()

		c.Response(uds.Res{})
	})

	if err := udsServer.Listen(); err != nil {
		logger.LoggerSugar.Errorf("init uds server error: %s", err.Error())
	}
}
