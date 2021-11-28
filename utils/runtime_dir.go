/*
Project: dirichlet runtime_dir.go
Created: 2021/11/20 by Landers
*/

package utils

import (
	"os"
	"path"
)

// GetAppDir get application dir
func GetAppDir() string {
	p, e := os.Getwd()
	if e != nil {
		return ""
	}

	return p
}

// CalDir calculate dirs
func CalDir(dirs ...string) string {
	return path.Join(dirs...)
}
