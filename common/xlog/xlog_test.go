package xlog

import (
	"fmt"
	"testing"
	"time"
)

func TestInitLog(t *testing.T) {
	logger := InitLog("./", "debug", uint(8))
	for i := 0; i < 100000; i++ {
		testLog(logger)
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func testLog(logger ILog) {
	logger.Debug("zhanglei")
	logger.Warn("zhanglei")
	logger.Info("zhanglei")
	logger.Fatal("zhanglei")
}
