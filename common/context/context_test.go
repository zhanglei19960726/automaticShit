package context

import (
	"automaticshit/common/xlog"
	"testing"
)

func TestMyContext(t *testing.T) {
	logger := xlog.InitLog("../xlog/", "debug", 8)
	ctx := NewContext(logger)
	ctx.Debug("zhanglei111")
}
