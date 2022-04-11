package main

import (
	"automaticshit/common/config"
	"automaticshit/common/context"
	"automaticshit/common/xlog"
	"automaticshit/cron"
	"flag"
	"fmt"
)

func main() {
	cfgs := flag.String("c", "./etc/config.json", "配置文件")
	flag.Parse()
	if err := config.LoadConfig(*cfgs); err != nil {
		panic(err)
	}
	cfg := config.GetConfig()
	logger := xlog.InitLog(cfg.Log.LogPath, cfg.Log.LogLevel, cfg.Log.LogSave)
	ctx := context.NewContext(logger)
	config.CfgWatch(ctx, *cfgs)
	ch := make(chan int)
	cro, err := cron.NewCron(ctx, cfg.CronConfig.Space, func() {
		fmt.Println("zhanglei")
	})
	if err != nil {
		panic(err)
	}
	config.RegisterReloadCfgFunc(cro.ReloadConfig)
	<-ch
}
