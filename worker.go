package workerpool

import (
	"sync"
)

type worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

// newWorker creates takes a numeric id and a channel w/ worker pool.
func newWorker(id int, workerPool chan chan Job) worker {
	return worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

func (w worker) start(wg *sync.WaitGroup) {
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				job()
				wg.Done()
			case <-w.quitChan:
				// We have been asked to stop.
				return
			}
		}
	}()
}

func (w worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
