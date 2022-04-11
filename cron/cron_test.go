package cron

import (
	"automaticshit/common/config"
	"automaticshit/common/context"
	"automaticshit/common/xlog"
	"fmt"
	"testing"
)

func TestCron(t *testing.T) {
	cfgPath := "../etc/config.json"
	if err := config.LoadConfig(cfgPath); err != nil {
		t.Fatal(err.Error())
	}
	cfg := config.GetConfig()
	logger := xlog.InitLog(cfg.Log.LogPath, cfg.Log.LogLevel, cfg.Log.LogSave)
	ctx := context.NewContext(logger)
	c := make(chan int)
	cr, err := NewCron(ctx, "0 28 15 * * *", func() {
		fmt.Println("zhanglei")
	})
	if err != nil {
		t.Fatal(err)
	}
	_ = cr
	<-c
}
