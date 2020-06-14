package lib

import "sync"

type Task struct {
	finished bool
	started bool
	err error
	handler TaskHandler
	mux sync.Mutex
	runningLock sync.Mutex
}

func NewTask(handler TaskHandler) *Task {
	return &Task{
		finished: false,
		err:      nil,
		handler:  handler,
	}
}

func (t *Task) Handler() TaskHandler {
	return t.handler
}

func (t *Task) Error() error {
	return t.err
}

func (t *Task) Finished() bool {
	return t.finished
}

func (t *Task) Progress() float32 {
	return t.handler.Progress()
}

func (t *Task) Start() {
	t.mux.Lock()
	defer t.mux.Unlock()
	if t.started {
		return
	}
	t.started = true
	t.runningLock.Lock()

	go func() {
		err := t.handler.Run()
		t.err = err
		t.finished = true
		t.runningLock.Unlock()
	}()
}

func (t *Task) Wait() error {
	t.runningLock.Lock()
	t.runningLock.Unlock()

	return t.err
}

func (t *Task) finish(err error) {
	t.finished = true
	t.err = err
}
