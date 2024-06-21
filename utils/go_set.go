/*
   Create: 2024/6/22
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package utils

import mapset "github.com/deckarep/golang-set/v2"

// 数据结构set
// NewSet 默认创建string类型的set

func NewSet() mapset.Set[string] {
	return mapset.NewSet[string]()
}

func NewSetInt() mapset.Set[int] {
	return mapset.NewSet[int]()
}

func NewSetAny() mapset.Set[any] {
	return mapset.NewSet[any]()
}
