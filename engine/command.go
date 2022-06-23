package engine

import (
	"fmt"
)

type printCommand struct {
	Arg string
}

func (print *printCommand) Execute(loop handler) {
	fmt.Println(print.Arg)
}

type catCommand struct {
	Arg1 string
	Arg2 string
}

func (pcc *catCommand) Execute(Handler handler) {
	res := pcc.Arg1 + pcc.Arg2
	Handler.Post(&printCommand{Arg: res})
}

type stopCommand struct{}

func (s stopCommand) Execute(h handler) {
  h.(*EventLoop).stop = true
}