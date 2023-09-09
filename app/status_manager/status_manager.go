/*
   Create: 2023/7/31
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

// Package status_manager
//
// 服务状态树的生成
// 性能树以json文件的形式提供 定时刷新
package status_manager

func UpdateStatusTree() error {
	tree, err := generateStatusTree()
	if err != nil {
		return err
	}

	return writeStatusOptFile(tree)
}

func generateStatusTree() (StatusTree, error) {
	var tree StatusTree

	return tree, nil
}
