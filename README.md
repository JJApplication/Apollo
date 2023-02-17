# Apollo
<a href="https://goreportcard.com/report/github.com/JJApplication/Apollo"><img src="./copyright/goreport.svg" /></a>
<a><img src="./copyright/version.svg"/></a>
<a href="http://service.renj.io"><img src="./copyright/renj.io.svg"/></a>
<a href="https://github.com/JJApplication"><img src="./copyright/copyright-JJService.svg"/></a>

`apollo` 是一个微服务注册，部署，管理服务
适用于**JJApplication**的一整套完整微服务

在线地址 [service.renj.io](http://service.renj.io)

## 配置文件

```json
{
  "service_root": "/renj.io",
  "app_root": "/renj.io/app",
  "app_manager": "/renj.io/manager",
  "app_cache_dir": "/renj.io/cache",
  "app_log_dir": "/renj.io/log",
  "app_tmp_dir": "/renj.io/tmp",
  "app_back_up": "/renj.io/backup",
  "app_pids": "/renj.io/pids",
  "log": {
    "enable_log": "",
    "enable_stack": "no",
    "enable_function": "no",
    "enable_caller": "no",
    "log_file": "",
    "encoding": ""
  },
  "db": {
    "sqlite": {},
    "mongo": {
      "name": "ApolloMongo",
      "url": "192.168.100.10:27017",
      "user": "",
      "pass_wd": ""
    },
    "redis": {}
  },
  "server": {
    "host": "0.0.0.0",
    "port": 9090,
    "uds": "/tmp/Apollo.sock"
  },
  "ci": {
    "docker_host": "tcp://192.168.100.10:2375",
    "docker_timeout": 5,
    "docker_api_version": "1.41"
  }
}
```

**Apollo**使用以JSON Schema为基础的`.pig`文件作为配置文件

配置分为5部分

`:root` 根路径下的键负责配置apollo运行时读取的路径

`log` 日志路径和格式的配置

`db` **Apollo**采用mongo数据库，主要负责mongo连接的配置

`server` 运行时http服务器的配置

`ci` 后端连接的远程docker接口配置

## 服务架构

```bash
- NoEngine
	- **Apollo**
	- service1
	- service2
```

在**JJAPP**服务中**Apollo**处于顶层负载均衡服务下层，和其他微服务同级

Apollo的运行时级别高于普通微服务，并且接管了所有微服务的注册和进程操作，不通过Apollo启动的服务是无法被统一管理的

## 服务模型

**JJAPP**微服务的模型定义于[OctopusMeta](https://github.com/JJApplication/octopus_meta)， 负责记录微服务元数据信息和运行时数据

形如：

```json
{
  "name": "demo",
  "id": "app_demo",
  "type": "app",
  "release_status": "unreleased",
  "eng_des": "",
  "chs_des": "",
  "link": "",
  "manage_cmd": {
    "start": "start.sh",
    "stop": "stop.sh",
    "restart": "restart.sh",
    "force_kill": "kill.sh",
    "check": "check.sh"
  },
  "meta": {
    "author": "",
    "domain": "",
    "language": null,
    "create_date": "",
    "version": "0.0.1",
    "dynamic_conf": false,
    "conf_type": "",
    "conf_path": ""
  },
  "run_data": {
    "envs": null,
    "ports": null,
    "random_port": false,
    "host": "",
    "run_dep": null,
    "stop_chain": null
  },
  "runtime": {
    "pid": "",
    "ports": null,
    "stop_operation": false
  },
  "resource_limit": {
    "min_cpu": 0,
    "max_cpu": 0,
    "min_mem": 0,
    "max_mem": 0,
    "ave_cpu_peak": 0,
    "ave_mem_peak": 0,
    "max_read": 0,
    "max_write": 0,
    "max_request": 0,
    "max_client": 0
  }
}
```

### 加载时机

在微服务启动时会从默认配置的模型目录下加载微服务模型到内存中

> 何时刷新？
>
> 模型数据会持久化到数据库中，在每次服务重启时会刷新数据库数据
>
> 同时会有后台定时任务负责对模型目录进行重载，保证模型是最新的

### 运行时模型

**运行时数据**指微服务运行中写入到Apollo的临时数据，在模型中仅作定义并初始化未空值
仅会存储到数据库中而不会因为用户的改动重新持久化到模型文件中

### 运行时调用链

**Apollo**负责管理所有的微服务，一些微服务并不支持`独立启动`即需要依赖其他微服务
`run_data.run_dep`和`run_data.stop_chain`就是为此而生

`run_dep` 定义了微服务启动时依赖的其他微服务，并在启动时同时保证依赖的服务被启动

`stop_chain` 定义了微服务停止链，在微服务停止时同时应该停止的其他微服务

## 服务管理

**Apollo**对服务操作进行了抽象，并对每一个服务进行单独的配置管理

服务的脚本存储在服务同名的`$APP_NAME`目录下

```bash
├── check.sh
├── start.sh
└── stop.sh

0 directories, 3 files
```

默认支持`启动`，`停止`，`状态检查`操作， `octopus`模型额外定义了`restart`和`kill`来负责自定义重启和强制停止

## 模块

### 服务管理模块

服务管理模块为`app_manager`负责对微服务的状态进行操作

### 容器管理模块

容器管理模块为`docker_manager`负责对远程容器服务进行操作

### 文件管理模块

文件管理模块为`file_manager` 负责对有权限的用户提供`$APP_ROOT`下文件目录的查看和操作权限

### 任务管理模块

任务管理模块为`task_manager`负责对**Apollo**运行时中注册的定时任务，轮询任务，单次任务进行统一管理

### 告警模块

告警模块为`alarm_manager`负责对错误异常进行告警通知，在**JJAPP**中**Hermes**服务负责全局的告警邮件发送，**Apollo**会调用接口通知**Hermes**进行邮件发送以及错误上报

### 鉴权模块

鉴权模块为`token_manager`负责对登录用户进行鉴权

### 流水线模块(开发中)

流水线模块为`pipeline_manager`负责调用服务器上的`CI/CD`程序或`github workflow`进行自动编译发布任务

## 动态模块

在设计时为保证**Apollo**后续功能的可扩展性，对于不依赖**Apollo**运行时数据的数据处理模块进行了拆分

以`plugin`的形式动态注册到**Apollo**的运行时上

在相应的路由模块变更时，**Apollo**会读取`modules`目录下的模块链接库将模块中定义的路由动态注册到运行时中

```bash
├── module-sys.so
├── module-influx.so
├── module-db-migrate.so
└── module-tmp.so
```



## CLI(apolliCLI)

**apolloCLI**是专为**Apollo**设计的服务器命令行终端

```bash
ApolloCLI » help

 █████╗ ██████╗  ██████╗ ██╗     ██╗      ██████╗ 
██╔══██╗██╔══██╗██╔═══██╗██║     ██║     ██╔═══██╗
███████║██████╔╝██║   ██║██║     ██║     ██║   ██║
██╔══██║██╔═══╝ ██║   ██║██║     ██║     ██║   ██║
██║  ██║██║     ╚██████╔╝███████╗███████╗╚██████╔╝
╚═╝  ╚═╝╚═╝      ╚═════╝ ╚══════╝╚══════╝ ╚═════╝


Apollo交互式终端

Commands:
=========
  address    查看连接的unix地址
  app        显示注册的微服务列表
  backup     全局同步备份
  check      检查Apollo服务状态
  clear      clear the screen
  exit       exit the shell
  help       use 'help [command]' for command help
  reconnect  重连指定的unix地址
  reload     重载微服务模型文件
  restart    重启指定微服务
  start      启动指定微服务
  status     查看指定微服务
  stop       停止指定微服务
  sync       同步指定微服务
  version    查看ApolloCLI版本

ApolloCLI » 
```

除了通过web api的形式还可以通过命令行程序对微服务进行管理

## CJA(create-jjapp)

**CJA**是基于`octopus`模型编写的微服务创建命令行脚手架

```bash
➜  ~ cja 
👻 Thanks for using CJA to create application of projectJJ
🐼 Copyright: CreateJJAPP
🔗 Go to site: http://renj.io
📧 Email: liaorenj@gmail.com

? your project name:  1
? Choose application type: empty
? Choose metadata type: pig
? What protocols do you prefer: http
? Do you like to create the project right now? Yes
[CJA] start to create project
[CJA] start to create project tag
[CJA] start to create project meta
[CJA] 1.pig created
[CJA] start to create project scripts
[CJA] create script start.sh
[CJA] create script stop.sh
[CJA] create script check.sh
[CJA] create script restart.sh
[CJA] create script kill.sh
[CJA] 🥤 Enjoy yourself!!!
```

**CJA**可以创建一个基于`octopus`模型定义的基础微服务，包含了服务配置信息，服务脚本和服务模板

## 尝试

想要尝试**Apollo**只需要下载微服务的打包集合[Apollo Release](https://github.com/JJApplication/Apollo/releases)

每当Apollo发布新的版本时，你都可以通过release页面快速获取包

```bash
➜  Apollo git:(master) ✗ ./Apollo
[Apollo] load config in 1ms
[Apollo] 🚀🚀🚀
⌛ StartTime: 2023-02-17 15:42:38.838735184 +0000 UTC m=+0.162816595
🔥 Listening on http://0.0.0.0:9090
📁 ServiceRoot: /renj.io
📁 AppRoot: /renj.io/app
📁 AppManger: /renj.io/manager
📁 AppLog: /renj.io/log
📁 BackUpDir: /renj.io/backup
📁 CacheDir: /renj.io/cache

{"Level":"INFO","Time":"2023-2-17 15:42:38","Message":"logger init success","Name":"[Apollo]"}
```

**Apollo**默认依赖`mongo`和`docker`服务来存储数据和管理容器

没有这些依赖服务时，运行时会记录错误日志但是你仍然可以访问其他的功能

## 开发

如何开发？

fork本项目后进入根目录

```bash
go mod tidy
bash build/build.sh
```



## LICENSE

基于MIT License
