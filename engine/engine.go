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

func (q *Queue) push(command Command) {
	q.Lock()
	defer q.Unlock()

	q.cmdArray = append(q.cmdArray, command)
	if q.awaiting {
		q.awaiting = false
		q.signalToAwait <- struct{}{}
	}

}

func (q *Queue) pull() Command {
	q.Lock()
	defer q.Unlock()

	if q.empty() {
		q.awaiting = true
		q.Unlock()
		<-q.signalToAwait
		q.Lock()
	}

	res := q.cmdArray[0]
	q.cmdArray[0] = nil
	q.cmdArray = q.cmdArray[1:]
	return res
}

func (q *Queue) empty() bool {
	return len(q.cmdArray) == 0
}
