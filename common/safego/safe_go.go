package safego

import (
	"automaticshit/common/context"
	"runtime/debug"
)

func SafeFunc(ctx context.IContext, f func()) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Fatal(err)
			ctx.Fatal(string(debug.Stack()))
		}
	}()
	f()
}

func SafeGo(ctx context.IContext, f func()) {
	go SafeFunc(ctx, f)
}
