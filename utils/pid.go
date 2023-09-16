/*
   Create: 2023/9/15
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import "os"

func GetRuntimePID() int {
	return os.Getpid()
}
