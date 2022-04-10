package jobqueue

import (
	"errors"
	"otter-calendar/http/response"

	"github.com/EricChiou/jobqueue"
)

type worker struct {
	run  func() interface{}
	wait *chan interface{}
}

func Init() {
	run(&User.queue)
	run(&Event.queue)
}

func Wait() {
	User.queue.Wait()
}

func run(queue *jobqueue.Queue) {
	queue.SetWorker(func(w interface{}) {
		if w, ok := w.(worker); ok {
			*w.wait <- w.run()
		} else {
			*w.wait <- errors.New(string(response.Error))
		}
	})
	queue.Run()
}
