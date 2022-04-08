package automaticshit

import (
	"automaticshit/common/config"
	"automaticshit/common/context"
	"automaticshit/common/xlog"
	"testing"
)

func TestAutomaticShit(t *testing.T) {
	people := []string{"test1", "test2", "test3", "test4", "test5"}
	shit := AutomaticShit(people, 2, 2, 31)
	t.Log(shit)
	t.Log(len(shit))
}

func TestAutomticShitMgr(t *testing.T) {
	cfgPath := "../etc/config.json"
	if err := config.LoadConfig(cfgPath); err != nil {
		t.Fatal(err.Error())
	}
	cfg := config.GetConfig()
	logger := xlog.InitLog(cfg.Log.LogPath, cfg.Log.LogLevel, cfg.Log.LogSave)
	ctx := context.NewContext(logger)
	config.CfgWatch(ctx, cfgPath)
	d := make(chan bool)
	a, err := NewAutomaticShitMgr(ctx, cfg.DataSavePath)
	if err != nil {
		t.Fatal(err.Error())
	}
	for i := 0; i < 30; i++ {
		shits := a.GetCurShit()
		t.Log(len(a.shits.Shit))
		t.Log(shits)
	}
	<-d
}
