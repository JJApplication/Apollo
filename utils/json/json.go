package json

import (
	jsoniter "github.com/json-iterator/go"
	"os"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseJson 通过byte读取
func ParseJson(data []byte, m interface{}) error {
	return json.Unmarshal(data, m)
}

// ParseJsonFile 解析文件
func ParseJsonFile(fileName string, m interface{}) error {
	data, err := os.ReadFile(fileName)
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

	return os.WriteFile(fileName, data, 0644)
}

func PrettyJson(v interface{}) string {
	s, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}

	return string(s)
}

func JsonString(v interface{}) string {
	s, err := json.Marshal(v)
	if err != nil {
		return ""
	}

	return string(s)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func MustMarshal(v interface{}) []byte {
	s, err := Marshal(v)
	if err != nil {
		return nil
	}
	return s
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
