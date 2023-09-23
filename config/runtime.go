/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package config

type RuntimeConf struct {
	ServiceRoot string `json:"serviceRoot"`
	APPRoot     string `json:"appRoot"`
	APPManager  string `json:"managerDir"`
	APPCacheDir string `json:"cacheDir"`
	APPLogDir   string `json:"logDir"`
	APPTmpDir   string `json:"tmpDir"`
	APPBackUp   string `json:"backupDir"`

	EnableStack    bool   `json:"enableStack"`
	EnableFunction bool   `json:"enableFunc"`
	EnableCaller   bool   `json:"enableCall"`
	LogFile        string `json:"logFile"`

	UICache     bool     `json:"enableCache"`
	UICacheTime int      `json:"cacheTime"`
	UIRouter    []string `json:"uiRouter"`
	AuthExpire  int      `json:"authExpire"`

	PID              int    `json:"pid"`
	Port             int    `json:"port"`
	UDS              string `json:"uds"`
	Mongo            string `json:"mongo"`
	DockerApi        string `json:"dockerApi"`
	DockerApiVersion string `json:"dockerApiVersion"`
	Goroutines       int    `json:"goroutines"`
	MaxProcs         int    `json:"maxProcs"`
}
