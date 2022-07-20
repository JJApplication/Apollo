// Package engine
//
// Engine for HTTP/Plnack
// When using engine type=plnack, Plnack data will return
//
// innerEngine which is for service connect by tcp
// using RPC by plnack-proto
// MiddleWare中间件 支持从配置文件加载，支持通过so的plugins机制加载
// MiddleWare的so库需要符合统一的入口函数，并且放置于指定的目录下lib
package engine

