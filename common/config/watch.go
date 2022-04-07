package config

import (
	"automaticshit/common/context"
	"automaticshit/common/safego"

	"github.com/fsnotify/fsnotify"
)

func watch(ctx context.IContext, watcher *fsnotify.Watcher) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			ctx.Debug("cfg watch event", event)
			if event.Op != fsnotify.Write {
				break
			}
			if err := readConfig(event.Name); err != nil {
				ctx.Error("readConfig error", err.Error())
				break
			}
			ctx.Debug("cfg is ", GetConfig())
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			ctx.Error("cfg watch error", err.Error())
		}
	}
}

// cfgWatch 开始配置文件监听
func CfgWatch(ctx context.IContext, path string) error {
	if err := readConfig(path); err != nil {
		return err
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	watcher.Add(path)
	safego.SafeGo(ctx, func() {
		watch(ctx, watcher)
	})
	return nil
}
