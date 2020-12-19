package eventloop

type EventLoop struct {
	queue *Queue

	stop bool
	stopSignal chan struct{}
}

func (e *EventLoop) Start() {
	e.queue = &Queue{ signal: make(chan struct{}) }
	e.stopSignal = make(chan struct{}, 10)

	go func() {
		for !e.stop || !e.queue.empty() {
			command := e.queue.pop()
			command.Execute(e)
		}
		e.stopSignal <- struct{}{}
	}()
}

func (e *EventLoop) Post(command Command, isInner bool) {
	if !e.stop || isInner {
		e.queue.push(command)
	}
}

func (e *EventLoop) AwaitFinish() {
	e.Post(CommandFunc(func(h Handler) {
		e.stop = true
	}), false)
	<- e.stopSignal
}
