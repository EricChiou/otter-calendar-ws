package jobqueue

import "github.com/EricChiou/jobqueue"

type event struct {
	queue jobqueue.Queue
}

func (e *event) NewEventQueueJob(run func() interface{}) error {
	wait := make(chan interface{})
	e.queue.Add(worker{run: run, wait: &wait})

	result := <-wait
	switch r := result.(type) {
	case error:
		return r
	default:
		return nil
	}
}

var Event event = event{
	queue: jobqueue.New(1024),
}
