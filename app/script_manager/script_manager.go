/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"sync"
)

var scripts []scriptModel

func init() {
	UpdateScript(LoadScripts())
	autoLoadScripts()
}

func UpdateScript(data []scriptModel) {
	lock := sync.Mutex{}
	lock.Lock()
	scripts = data
	defer lock.Unlock()
}

func GetScripts() []scriptModel {
	return scripts
}
