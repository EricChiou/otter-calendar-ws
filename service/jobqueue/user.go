package jobqueue

import "github.com/EricChiou/jobqueue"

type user struct {
	queue jobqueue.Queue
}

func (u *user) NewUserQueueJob(run func() interface{}) error {
	wait := make(chan interface{})
	u.queue.Add(worker{run: run, wait: &wait})

	result := <-wait
	switch r := result.(type) {
	case error:
		return r
	default:
		return nil
	}
}

var User user = user{
	queue: jobqueue.New(1024),
}
