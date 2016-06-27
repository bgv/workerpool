package workerpool

import (
	"sync"
)

type worker struct {
	id         int
	wg         *sync.WaitGroup
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
}

// newWorker creates takes a numeric id and a channel w/ worker pool.
func newWorker(id int, wg *sync.WaitGroup, workerPool chan chan Job) worker {
	return worker{
		id:         id,
		wg:         wg,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		quitChan:   make(chan bool),
	}
}

func (w worker) start() {
	w.wg.Add(1)
	go func() {
		for {
			// Add my jobQueue to the worker pool.
			w.workerPool <- w.jobQueue

			select {
			case job := <-w.jobQueue:
				// Dispatcher has added a job to my jobQueue.
				job()
			case <-w.quitChan:
				// We have been asked to stop.
				w.wg.Done()
				return
			}
		}
	}()
}

func (w worker) stop() {
	w.quitChan <- true
}
