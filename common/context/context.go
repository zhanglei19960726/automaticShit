package context

import (
	"automaticshit/common/xlog"
)

type IContext interface {
	xlog.ILog
}

type context struct {
	xlog.ILog
}

func NewContext(logger xlog.ILog) IContext {
	return &context{
		ILog: logger,
	}
}
