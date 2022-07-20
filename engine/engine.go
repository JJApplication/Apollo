/*
Project: Apollo engine.go
Created: 2021/11/26 by Landers
*/

package engine

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/JJApplication/Apollo/logger"
	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

const (
	PLNACK_PROTO = "+plnack"
	HTML_GLOB    = "web/**/*"
	StaticPath   = "./web"
)

var engine Engine

// Engine 一个包含gin和plnack的引擎
type Engine struct {
	Config    *EngineConfig
	ginEngine *gin.Engine

	MiddleWare   []gin.HandlerFunc
	EnablePlnack bool
	HeaderMap    map[string]string
}

type EngineConfig struct {
	Host string
	Port int
}

func NewEngine(cf *EngineConfig) *Engine {
	engine = Engine{
		Config:       cf,
		ginEngine:    newGin(),
		EnablePlnack: true,
	}

	return &engine
}

func newGin() *gin.Engine {
	g := gin.New()

	g.LoadHTMLGlob(HTML_GLOB)
	g.StaticFS("/static", http.Dir(StaticPath))
	g.NoRoute(MiddlewareNoRoute())
	g.NoMethod(MiddlewareNoMethod())
	loadMiddleWare(g)
	loadMiddlePlugins(g)
	return g
}

// Init 初始化全部配置
// 此函数应该单独执行
func (e *Engine) Init() {
	e.ginEngine.Use(e.MiddleWare...)
}

// GetEngine 获取内部的engine 注册路由
func (e *Engine) GetEngine() *gin.Engine {
	return engine.ginEngine
}

// LoadMiddleWare 加载中间件
func (e *Engine) LoadMiddleWare(m []gin.HandlerFunc) {
	e.MiddleWare = m
}

// SetHeadersMap 设置请求头
func (e *Engine) SetHeadersMap(m map[string]string) {
	e.HeaderMap = m
}

func (e *Engine) Run() error {
	path := fmt.Sprintf("%s:%d", e.Config.Host, e.Config.Port)
	logger.Logger.Info(fmt.Sprintf("listening on %s", path))
	return e.ginEngine.Run(path)
}

func (e *Engine) RunServer() error {
	path := fmt.Sprintf("%s:%d", e.Config.Host, e.Config.Port)
	logger.Logger.Info(fmt.Sprintf("listening on %s", path))

	server := new(http.Server)
	server.Handler = e.ginEngine
	server.Addr = path
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.IdleTimeout = 5 * time.Second
	server.MaxHeaderBytes = 5 << 20

	sig := make(chan os.Signal, 1)
	go RegisterSignals(server, sig)
	return server.ListenAndServe()
}

func (e *Engine) RunServerTLS(cert, key string) error {
	path := fmt.Sprintf("%s:%d", e.Config.Host, e.Config.Port)
	logger.Logger.Info(fmt.Sprintf("listening on %s", path))

	server := new(http.Server)
	server.Handler = e.ginEngine
	server.Addr = path
	server.ReadTimeout = 10 * time.Second
	server.WriteTimeout = 10 * time.Second
	server.IdleTimeout = 5 * time.Second
	server.MaxHeaderBytes = 5 << 20

	return server.ListenAndServeTLS(cert, key)
}

// Group 生成路由分组
func (e *Engine) Group(r string, ware ...gin.HandlerFunc) *gin.RouterGroup {
	return e.ginEngine.Group(r, ware...)
}

// Handle 路由控制
func (e *Engine) Handle(method, r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.Handle(method, r+PLNACK_PROTO, handler...).Use(MiddlewarePlnack())
	}
	e.ginEngine.Handle(method, r, handler...)
}

func (e *Engine) GET(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.GET(r+PLNACK_PROTO, handler...).Use(MiddlewarePlnack())
	}
	e.ginEngine.GET(r, handler...)
}

func (e *Engine) POST(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.POST(r+PLNACK_PROTO, handler...).Use(MiddlewarePlnack())
	}
	e.ginEngine.POST(r, handler...)
}

func (e *Engine) DELETE(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.DELETE(r+PLNACK_PROTO, handler...).Use(MiddlewarePlnack())
	}
	e.ginEngine.DELETE(r, handler...)
}

func (e *Engine) PUT(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.PUT(r+PLNACK_PROTO, handler...).Use(MiddlewarePlnack())
	}
	e.ginEngine.PUT(r, handler...)
}
