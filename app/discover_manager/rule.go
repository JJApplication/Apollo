/*
   Create: 2024/1/17
   Project: Apollo
   Github: https://github.com/landers1037
   Copyright Renj
*/

package discover_manager

import (
	"github.com/JJApplication/Apollo/logger"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// 基于自动发现规则进行的自动发现任务触发

// DiscoverRule 按照文件修改时间与缓存对比是否一致来确定是否重加载
type DiscoverRule struct {
	WatchRoot   string   // 监控的根目录
	Include     []string // 包含的文件 为空时监控根目录下的全部文件
	Exclude     []string // 排除的文件
	CompareFlag int      // 默认比较文件修改时间 存在多个比较量时必须同时改变才生效

	loading bool                // 正在加载的标识符
	cache   map[string]fileInfo // 以文件名称为键
}

type fileInfo struct {
	modify int64
	size   int64
}

const (
	FlagAll = iota
	FlagSize
	FlagModify
)

func NewDiscoverRule(rule DiscoverRule) *DiscoverRule {
	logger.LoggerSugar.Infof("%s load Discover Root: %s", DiscoverManagerPrefix, rule.WatchRoot)
	return &DiscoverRule{
		WatchRoot:   rule.WatchRoot,
		Include:     rule.Include,
		Exclude:     rule.Exclude,
		CompareFlag: rule.CompareFlag,
		loading:     false,
		cache:       make(map[string]fileInfo),
	}
}

func (d *DiscoverRule) NeedDiscover() bool {
	if d.loading {
		return false
	}
	fileStat := d.loadAll()
	// nocache
	if len(d.cache) <= 0 {
		return true
	}
	// file list changes
	if len(d.cache) > 0 && len(fileStat) != len(d.cache) {
		return true
	}
	// compare and refresh cache
	for file, info := range fileStat {
		// 从缓存取值对比
		if _, ok := d.cache[file]; !ok {
			return true
		}
		switch d.CompareFlag {
		case FlagAll:
			if d.cache[file].modify != info.modify && d.cache[file].size != info.size {
				return true
			}
		case FlagModify:
			if d.cache[file].modify != info.modify {
				return true
			}
		case FlagSize:
			if d.cache[file].size != info.size {
				return true
			}
		default:
			continue
		}
	}

	return false
}

// CheckFile 检查文件更新
func (d *DiscoverRule) CheckFile(f string) bool {
	if _, ok := d.cache[f]; !ok {
		return true
	}

	info, err := os.Stat(f)
	if err != nil {
		return false
	}

	if info.ModTime().Unix() != d.cache[f].modify || info.Size() != d.cache[f].size {
		return true
	}

	return false
}

// 加载监控目录下的全部文件信息到缓存中
func (d *DiscoverRule) loadAll() map[string]fileInfo {
	d.loading = true

	if strings.TrimSpace(d.WatchRoot) == "" {
		return nil
	}
	tmp := make(map[string]fileInfo)
	err := filepath.Walk(d.WatchRoot, func(path string, info fs.FileInfo, err error) error {
		// 优先判断排除文件
		if d.containExclude(info.Name()) {
			return nil
		}
		// 判断包含文件
		if !d.containInclude(info.Name()) {
			return nil
		}
		// 文件合法
		tmp[path] = fileInfo{
			modify: info.ModTime().Unix(),
			size:   info.Size(),
		}

		return err
	})

	if err != nil {
		return nil
	}

	d.loading = false
	return tmp
}

func (d *DiscoverRule) containInclude(f string) bool {
	if len(d.Include) > 0 {
		for _, s := range d.Include {
			if s == f {
				return true
			}
		}
	}

	return true
}

// 排除列表为空时 始终返回true代表未排除
func (d *DiscoverRule) containExclude(f string) bool {
	if len(d.Exclude) > 0 {
		for _, s := range d.Exclude {
			if s == f {
				return true
			}
		}
	}

	return false
}
