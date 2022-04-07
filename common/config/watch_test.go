package config

import (
	"automaticshit/common/context"
	"automaticshit/common/xlog"
	"testing"
)

func TestCfgWatch(t *testing.T) {
	logger := xlog.InitLog("../xlog/", "debug", 8)
	ctx := context.NewContext(logger)
	done := make(chan bool)
	CfgWatch(ctx, "../../etc/config.json")
	<-done
}
