/*
Project: Apollo engine.go
Created: 2021/11/26 by Landers
*/

package engine

import (
	"fmt"
	"github.com/JJApplication/Apollo/middleware"
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
	HtmlTmpl     = "web/tmpl/**/*"
	StaticPath   = "web/static"
	TmplPath     = "web/tmpl"
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
	gin.SetMode(gin.ReleaseMode)
	g := gin.New()

	// 默认会开启gzip
	// 废弃方法 不再使用render渲染go模板
	g.LoadHTMLGlob(HtmlTmpl)
	g.StaticFS("/static", http.Dir(StaticPath))
	g.StaticFS("/tmpl", http.Dir(TmplPath))
	g.NoRoute(middleware.MiddlewareNoRoute())
	g.NoMethod(middleware.MiddlewareNoMethod())
	middleware.LoadMiddleWare(g)
	middleware.LoadMiddlePlugins(g)
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
	logger.LoggerSugar.Infof("listening on %s", path)
	return e.ginEngine.Run(path)
}

func (e *Engine) RunServer() error {
	path := fmt.Sprintf("%s:%d", e.Config.Host, e.Config.Port)
	logger.LoggerSugar.Infof("listening on %s", path)

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
	logger.LoggerSugar.Infof("listening on %s", path)

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
		e.ginEngine.Handle(method, r+PLNACK_PROTO, handler...).Use(middleware.MiddlewarePlnack())
	}
	e.ginEngine.Handle(method, r, handler...)
}

func (e *Engine) GET(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.GET(r+PLNACK_PROTO, handler...).Use(middleware.MiddlewarePlnack())
	}
	e.ginEngine.GET(r, handler...)
}

func (e *Engine) POST(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.POST(r+PLNACK_PROTO, handler...).Use(middleware.MiddlewarePlnack())
	}
	e.ginEngine.POST(r, handler...)
}

func (e *Engine) DELETE(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.DELETE(r+PLNACK_PROTO, handler...).Use(middleware.MiddlewarePlnack())
	}
	e.ginEngine.DELETE(r, handler...)
}

func (e *Engine) PUT(r string, handler ...gin.HandlerFunc) {
	if e.EnablePlnack {
		e.ginEngine.PUT(r+PLNACK_PROTO, handler...).Use(middleware.MiddlewarePlnack())
	}
	e.ginEngine.PUT(r, handler...)
}
