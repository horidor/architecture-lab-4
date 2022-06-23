package engine

import (
	"sync"
	"fmt"
)

type command interface {
	Execute(handler handler)
}

type handler interface {
	Post(cmd command) (error)
}

type queue struct {
	sync.Mutex
	cmdArray      []command
	signalToAwait chan struct{}
	awaiting      bool
}

func (q *queue) push(Command command) {
	q.Lock()
	defer q.Unlock()

	q.cmdArray = append(q.cmdArray, Command)
	if q.awaiting {
		q.awaiting = false
		q.signalToAwait <- struct{}{}
	}

}

func (q *queue) pull() command {
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

func (q *queue) empty() bool {
	return len(q.cmdArray) == 0
}

type EventLoop struct {
	cmdQ          *queue
	stopSignal  chan struct{}
	stop bool
  }

  func (l *EventLoop) Start() {
	l.cmdQ = &queue{signalToAwait: make(chan struct{})}
	l.stopSignal = make(chan struct{})
	go func() {
	  for !l.stop || !l.cmdQ.empty() {
		cmd := l.cmdQ.pull()
		cmd.Execute(l)
	  }
	  l.stopSignal <- struct{}{}
	}()
  }
  
  func (l *EventLoop) Post(cmd command) (error) {
	if l.stop == true {
		return fmt.Errorf("Unexpected command after eventloop finish")
	}

	l.cmdQ.push(cmd)
	return nil
  }
  
  func (l *EventLoop) AwaitFinish() {
	l.Post(stopCommand{})
    <-l.stopSignal
  }