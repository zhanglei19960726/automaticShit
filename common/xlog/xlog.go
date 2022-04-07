package xlog

import (
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

const (
	infoLog  = "info"
	warnLog  = "warn"
	errorLog = "error"
	debugLog = "debug"
	fatalLog = "fatal"
)

type ILog interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

var (
	_log = logrus.New()
)

func InitLog(logPath string, logLevel string, save uint) ILog {
	_log.SetLevel(GetLogLevel(logLevel))
	_log.SetReportCaller(true)
	_log.SetOutput(&mineFormatter{})
	NewSimpleLogger(_log, logPath, save)
	return _log
}

func GetLogLevel(logLevel string) logrus.Level {
	switch strings.ToLower(logLevel) {
	case infoLog:
		return logrus.InfoLevel
	case errorLog:
		return logrus.ErrorLevel
	case warnLog:
		return logrus.WarnLevel
	case fatalLog:
		return logrus.FatalLevel
	case debugLog:
		return logrus.DebugLevel
	default:
		return logrus.DebugLevel
	}
}

func NewSimpleLogger(log *logrus.Logger, logPath string, save uint) {
	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: writer(logPath, save),
		logrus.InfoLevel:  writer(logPath, save),
		logrus.WarnLevel:  writer(logPath, save),
		logrus.ErrorLevel: writer(logPath, save),
		logrus.FatalLevel: writer(logPath, save),
	}, &mineFormatter{})
	log.AddHook(lfHook)

}

func writer(logPath string, save uint) *rotatelogs.RotateLogs {
	logPath += time.Now().Format("20060102")
	if !isExist(logPath) {
		if err := os.Mkdir(logPath, os.ModePerm); err != nil {
			panic(err)
		}
	}
	logier, err := rotatelogs.New(
		logPath+"/error_%H.log",
		rotatelogs.WithRotationTime(time.Minute*time.Duration(save)),
		rotatelogs.WithMaxAge(time.Duration(720)*time.Hour),
	)
	if err != nil {
		panic(err)
	}

	return logier
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
