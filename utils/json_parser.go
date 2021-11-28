/*
Project: dirichlet json_parser.go
Created: 2021/11/20 by Landers
*/

package utils

import (
	"io/ioutil"

	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseJson 通过byte读取
func ParseJson(data []byte, m interface{}) error {
	return json.Unmarshal(data, m)
}

// ParseJsonFile 解析文件
func ParseJsonFile(fileName string, m interface{}) error {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	return ParseJson(data, m)
}

// Save2Json 输出配置到byte
func Save2Json(m interface{}) ([]byte, error) {
	return json.Marshal(m)
}

// Save2JsonFile 保存为json文件
func Save2JsonFile(m interface{}, fileName string) error {
	data, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, data, 0644)
}
