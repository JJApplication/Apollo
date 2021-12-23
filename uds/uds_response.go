/*
Project: dirichlet uds_response.go
Created: 2021/12/22 by Landers
*/

package uds

import (
	"github.com/landers1037/dirichlet/utils"
)

// 响应带格式的json

type UDSRes struct {
	Error string `json:"error"`
	Data  string `json:"data"`
}

// UDSResponse UDS响应
// 数据可信 传入的结构体一定可以格式化
func UDSResponse(u UDSRes) []byte {
	b, err := utils.Save2Json(u)
	if err != nil {
		return []byte(ErrNoResponse)
	}
	return b
}
