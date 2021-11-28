/*
Project: dirichlet engine.go
Created: 2021/11/26 by Landers
*/

package engine

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

// Engine 一个包含gin和plnack的引擎
type Engine struct {
	Config      *EngineConfig
	ginEngine   *gin.Engine
	InnerEngine *InnerEngine

	MiddleWare      []gin.HandlerFunc
	EnablePlnack   bool
	HeaderMap map[string]string
}

type EngineConfig struct {
	Host string
	Port int
}

func NewEngine(cf *EngineConfig) *Engine {
	return &Engine{
		Config:    cf,
		ginEngine: newGin(),
	}
}

func newGin() *gin.Engine {
	return gin.New()
}

// Init 初始化全部配置
// 此函数应该单独执行
func (e *Engine) Init() {
	e.ginEngine.Use(e.MiddleWare...)
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
	return e.ginEngine.Run(path)
}

// Group 生成路由分组
func (e *Engine) Group(r string, ware ...gin.HandlerFunc) *gin.RouterGroup {
	return e.ginEngine.Group(r, ware...)
}

// Handle 路由控制
func (e *Engine) Handle(method, r string, handler ...gin.HandlerFunc) {
	e.ginEngine.Handle(method, r, handler...)
}

func (e *Engine) GET(r string, handler ...gin.HandlerFunc) {
	e.ginEngine.GET(r, handler...)
}

func (e *Engine) POST(r string, handler ...gin.HandlerFunc) {
	e.ginEngine.POST(r, handler...)
}

func (e *Engine) DELETE(r string, handler ...gin.HandlerFunc) {
	e.ginEngine.DELETE(r, handler...)
}

func (e *Engine) PUT(r string, handler ...gin.HandlerFunc) {
	e.ginEngine.PUT(r, handler...)
}