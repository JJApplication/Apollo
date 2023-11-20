/*
   Create: 2023/10/25
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package cmd

import "flag"

// 为相同命令注册多个flag参数

type FlagType int

type Flag struct {
	Args    []string
	Help    string
	Default any
	Func    func()
}

const (
	INT = iota
	INT64
	FLOAT
	STRING
	BOOL
)

// var should be an addressable value
func registerFlags(t FlagType, val any, flags Flag) func() {
	for _, arg := range flags.Args {
		switch t {
		case INT:
			flag.IntVar(val.(*int), arg, flags.Default.(int), flags.Help)
		case INT64:
			flag.Int64Var(val.(*int64), arg, flags.Default.(int64), flags.Help)
		case FLOAT:
			flag.Float64Var(val.(*float64), arg, flags.Default.(float64), flags.Help)
		case STRING:
			flag.StringVar(val.(*string), arg, flags.Default.(string), flags.Help)
		case BOOL:
			flag.BoolVar(val.(*bool), arg, flags.Default.(bool), flags.Help)
		default:
			continue
		}
	}

	return flags.Func
}

func registerDone() {
	flag.Parse()
}
