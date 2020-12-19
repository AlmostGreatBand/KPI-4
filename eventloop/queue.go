package eventloop

import (
	"sync"
)

type Queue struct {
	sync.Mutex

	commands []Command

	signal chan struct{}
	waited bool
}

func (q *Queue) push(command Command) {
	q.Lock()
	defer q.Unlock()

	q.commands = append(q.commands, command)
	if q.waited {
		q.waited = false
		q.signal <- struct{}{}
	}
}

func (q *Queue) pop() Command {
	q.Lock()
	defer q.Unlock()

	if q.empty() {
		q.waited = true
		q.Unlock()
		<- q.signal
		q.Lock()
	}

	command := q.commands[0]
	q.commands[0] = nil
	q.commands = q.commands[1:]
	return command
}

func (q *Queue) empty() bool {
	return len(q.commands) == 0
}