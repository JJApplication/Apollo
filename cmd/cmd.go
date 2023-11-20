/*
   Create: 2023/10/24
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package cmd

import (
	"fmt"
	"github.com/JJApplication/Apollo/vars"
	"os"
)

// values from vars
var (
	Version   = vars.BuildDate
	GitCommit = vars.GitCommit
	GOOS      = vars.GOOS
	GOARCH    = vars.GOARCH
	GOVersion = vars.GOVersion
)

// flags
var (
	FlagVersion bool
)

func ApolloCmd() {
	versionFunc := registerFlags(BOOL, &FlagVersion, Flag{
		Args:    []string{"v", "V", "version"},
		Help:    "version",
		Default: false,
		Func: func() {
			fmt.Printf("version: %s\narch: %s\ncommit: %s\n", Version, GOARCH, GitCommit)
		},
	})

	registerDone()

	if FlagVersion {
		versionFunc()
		os.Exit(0)
	}
}
