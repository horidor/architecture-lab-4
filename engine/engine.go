package engine

import (
	"sync"
)

type Command interface {
	Execute(handler Handler)
}

type Handler interface {
	Post(cmd Command)
}

type Queue struct {
	sync.Mutex
	cmdArray      []Command
	signalToAwait chan struct{}
	awaiting      bool
}
