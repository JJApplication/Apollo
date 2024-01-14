{
  "service_root": "/renj.io",
  "app_root": "/renj.io/app",
  "app_manager": "/renj.io/manager",
  "app_cache_dir": "/renj.io/cache",
  "app_log_dir": "/renj.io/log",
  "app_tmp_dir": "/renj.io/tmp",
  "app_back_up": "/renj.io/backup",
  "app_pids": "/renj.io/pids",
  "app_bridge": "rainbow_bridge",
  "log": {
    "enable_log": true,
    "enable_stack": false,
    "enable_function": false,
    "enable_caller": false,
    "log_file": "",
    "encoding": ""
  },
  "db": {
    "sqlite": {},
    "mongo": {
      "name": "ApolloMongo",
      "url": "192.168.100.10:27017",
      "user": "",
      "passwd": ""
    },
    "redis": {},
    "tidb": {
      "db":  "",
      "host": "",
      "port": 0,
      "user": "",
      "passwd": ""
    }
  },
  "server": {
    "host": "0.0.0.0",
    "port": 9090,
    "uds": "/tmp/Apollo.sock",
    "ui_cache": false,
    "ui_cache_time": 84600,
    "ui_router": [
      "/panel",
      "/panel/*route",
      "/next",
      "/next/*route"
    ],
    "auth_code": "",
    "auth_expire": 600,
    "account": "landers",
    "passwd": "123456"
  },
  "ci": {
    "docker_host": "tcp://192.168.100.10:2375",
    "docker_timeout": 5,
    "docker_api_version": "1.41"
  },
  "module": {
    "enable": false
  }
}