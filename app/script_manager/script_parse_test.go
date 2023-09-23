/*
   Create: 2023/9/23
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package script_manager

import (
	"github.com/JJApplication/Apollo/utils"
	"testing"
)

func TestParseScriptModel(t *testing.T) {
	var res []scriptModel
	err := utils.ParseJsonFile("script.json", &res)
	t.Log(err)
	t.Log(res)
}
