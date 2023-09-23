/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import "strings"

func RemoveAddrPort(addr string) string {
	if strings.Contains(addr, ":") {
		addrs := strings.Split(addr, ":")
		if len(addrs) < 1 {
			return ""
		}
		return addrs[0]
	}

	return addr
}
