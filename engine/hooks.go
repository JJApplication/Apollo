/*
Create: 2023/2/17
Project: Apollo
Github: https://github.com/landers1037
Copyright Renj
*/

package engine

import (
	"io/fs"
	"os"
	"path/filepath"
	"plugin"
	"strings"

	"github.com/JJApplication/Apollo/logger"
	"github.com/JJApplication/fushin/utils/module"
)

// 加载插件Altas

const (
	ModulePath   = "modules"
	logPrefix    = "[Hot Module]"
	ModuleLookup = "Module"
	So           = ".so"
)

var ApolloModules = module.CreateModules()

func LoadModules() []string {
	if _, e := os.Stat(ModulePath); os.IsExist(e) {
		return nil
	}

	var files []string
	_ = filepath.WalkDir(ModulePath, func(path string, d fs.DirEntry, err error) error {
		if err == nil && !d.IsDir() && d.Name() != "" && strings.HasSuffix(path, So) {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func Hooks(s *Engine) {
	modules := LoadModules()
	logger.LoggerSugar.Infof("%s load modules: %d", logPrefix, len(modules))

	for _, m := range modules {
		p, err := plugin.Open(m)
		if err != nil {
			logger.LoggerSugar.Errorf("%s load error: %s", logPrefix, err.Error())
		}
		symbol, err := p.Lookup(ModuleLookup)
		if err != nil {
			logger.LoggerSugar.Errorf("%s lookup error: %s", logPrefix, err.Error())
		}
		mod, ok := symbol.(module.M)
		if !ok {
			logger.LoggerSugar.Errorf("%s unexpected symbol error", logPrefix)
		}
		mod.Hooks(s.GetEngine())
		mod.Enable()
		module.AddModule(ApolloModules, module.Module{Name: mod.Name(), Status: true})
		logger.LoggerSugar.Infof("%s module: [%s] loaded", logPrefix, mod.Name())
	}
}
