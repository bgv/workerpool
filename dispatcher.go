package workerpool

import (
	"sync"
)

type dispatcher struct {
	wg         *sync.WaitGroup
	workers    chan worker
	workerPool chan chan Job
	numWorkers int
	jobPool    chan Job
	quitChan   chan bool
}

// newDispatcher creates, and starts new dispatcher object.
func newDispatcher(numWorkers int, jobQueue chan Job) *dispatcher {
	workerPool := make(chan chan Job, numWorkers)
	workers := make(chan worker, numWorkers)

	d := &dispatcher{
		wg:         &sync.WaitGroup{},
		workers:    workers,
		numWorkers: numWorkers,
		workerPool: workerPool,
		jobPool:    jobQueue,
		quitChan:   make(chan bool),
	}

	for i := 0; i < d.numWorkers; i++ {
		w := newWorker(i+1, d.wg, d.workerPool)
		w.start()
		d.workers <- w
	}

	go d.dispatch()

	return d
}

// stop waits for any jobs to finish and stops dispatcher
func (d *dispatcher) stop() {
	defer func() {
		// clear WorkerPool
		for range d.workerPool {
			if len(d.workerPool) == 0 {
				return
			}
		}
	}()

	for w := range d.workers {
		w.stop()
		if len(d.workers) == 0 {
			return
		}
	}
	d.wg.Wait()
}

// dispatch starts the dispatcher
func (d *dispatcher) dispatch() {
	defer func() {
		d.stop()
	}()

	for {
		select {
		case job := <-d.jobPool:
			// New job added, assign worker
			// from the pool and send the job
			go func(job Job) {
				workerQueue := <-d.workerPool
				workerQueue <- job
			}(job)
		case <-d.quitChan:
			return
		}
	}
}
