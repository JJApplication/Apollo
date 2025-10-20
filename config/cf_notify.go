package config

// 配置文件监听

import (
	"fmt"
	"github.com/JJApplication/Apollo/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/landers1037/configen"
)

func InitConfigNotify() {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("[File Notifier] create watcher error:", err.Error())
		return
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				fmt.Println("[File Notifier] event:", event)
				if event.Op == fsnotify.Write || event.Op == fsnotify.Rename || event.Op == fsnotify.Chmod || event.Op == fsnotify.Remove {
					fmt.Println("modified file:", event.Name)
					ApolloConf.lock.Lock()
					reloadErr := configen.ParseConfig(
						&ApolloConf,
						configen.Pig,
						utils.CalDir(
							utils.GetAppDir(),
							GlobalConfigRoot,
							GlobalConfigFile))
					if reloadErr != nil {
						fmt.Println("[File Notifier] reload config error: ", reloadErr.Error())
					}
					ApolloConf.lock.Unlock()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				if err != nil {
					fmt.Println("[File Notifier] notify config error:", err.Error())
				}
			}
		}
	}()

	// Add a path.
	err = watcher.Add(utils.CalDir(
		utils.GetAppDir(),
		GlobalConfigRoot,
		GlobalConfigFile))
	if err != nil {
		fmt.Println("[File Notifier] create notify error:", err.Error())
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}
