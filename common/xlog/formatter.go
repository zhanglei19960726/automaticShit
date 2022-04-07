package xlog

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

const timeFormat = "2006-01-02 15:04:05"

type mineFormatter struct{}

func (m *mineFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	file, line := m.caller(entry)
	msg := fmt.Sprintf("[%s]\t[%s]\t%s:%d\t%s\n", time.Now().Local().Format(timeFormat), strings.ToUpper(entry.Level.String()),
		file, line, entry.Message)
	return []byte(msg), nil
}

func (m *mineFormatter) caller(entry *logrus.Entry) (file string, line int) {
	if entry.HasCaller() {
		file = entry.Caller.File
		line = entry.Caller.Line
	} else {
		ok := false
		_, file, line, ok = runtime.Caller(2)
		if !ok {
			return
		}
	}
	slash := strings.LastIndex(file, "/")
	if slash >= 0 {
		file = file[slash+1:]
	}
	return
}

// Write 重置logrus输出流，所以实现个空方法
func (m *mineFormatter) Write(p []byte) (n int, err error) {
	return
}
