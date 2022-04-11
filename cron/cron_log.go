package cron

import "automaticshit/common/context"

type cronLog struct {
	ctx context.IContext
}

func newCronLog(ctx context.IContext) *cronLog {
	return &cronLog{ctx: ctx}
}

func (c *cronLog) Info(msg string, keysAndValues ...interface{}) {
	c.ctx.Info(msg, keysAndValues)
}

func (c *cronLog) Error(err error, msg string, keysAndValues ...interface{}) {
	c.ctx.Error(err.Error(), msg, keysAndValues)
}
