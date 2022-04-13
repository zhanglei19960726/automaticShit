package notic

import (
	"automaticshit/automaticshit"
	"automaticshit/common/context"
	"fmt"
)

type INotic interface {
	NoticShit(ctx context.IContext, autoMatic automaticshit.IAutoMaticShit) error
}

type Notic struct {
	i int
}

func NewNotic() *Notic {
	return &Notic{}
}

func (n *Notic) NoticShit(ctx context.IContext, autoMatic automaticshit.IAutoMaticShit) error {
	people := autoMatic.GetCurShit()
	// ctx.Debug("NoticShit", people)
	fmt.Println(people, n.i)
	n.i++
	return nil
}
