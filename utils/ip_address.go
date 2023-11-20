/*
   Create: 2023/9/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import (
	"net/http"
	"strings"
)

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

func GetRemoteIP(r *http.Request) string {
	remoteAddr := r.RemoteAddr
	realIP := r.Header.Get("X-Real-IP")

	if realIP != "" {
		return RemoveAddrPort(realIP)
	}

	return RemoveAddrPort(remoteAddr)
}
