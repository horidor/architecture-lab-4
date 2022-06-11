package engine

import (
	"fmt"
)

type PrintCommand struct {
	Arg string
}

func (print *PrintCommand) Execute(loop Handler) {
	fmt.Println(print.Arg)
}

type CatCommand struct {
	Arg1 string
	Arg2 string
}

func (pcc *CatCommand) Execute(handler Handler) {
	res := pcc.Arg1 + pcc.Arg2
	handler.Post(&PrintCommand{Arg: res})
}

type stopCommand struct{}

func (s stopCommand) Execute(h Handler) {
  h.(*EventLoop).stop = true
}