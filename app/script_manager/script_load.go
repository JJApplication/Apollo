/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"github.com/JJApplication/Apollo/config"
	"github.com/JJApplication/Apollo/utils"
	"path/filepath"
)

const ScriptManagerPrefix = "[Script Manager]"
const ScriptConfig = "script.json"

func LoadScripts() []scriptModel {
	var res []scriptModel
	if err := utils.ParseJsonFile(filepath.Join(config.GlobalConfigRoot, ScriptConfig), &res); err != nil {
		return []scriptModel{}
	}

	return res
}
