package workerpool

import (
	"log"
	"sync"
)

type dispatcher struct {
	wg         *sync.WaitGroup
	workers    chan worker
	workerPool chan chan Job
	numWorkers int
	jobPool    chan Job
}

// New creates, and starts new dispatcher object.
func newDispatcher(numWorkers int, jobQueue chan Job) *dispatcher {
	workerPool := make(chan chan Job, numWorkers)
	workers := make(chan worker, numWorkers)

	d := &dispatcher{
		wg:         &sync.WaitGroup{},
		workers:    workers,
		numWorkers: numWorkers,
		workerPool: workerPool,
		jobPool:    jobQueue,
	}

	log.Printf("Spawning %d background workers", d.numWorkers)
	for i := 0; i < d.numWorkers; i++ {
		w := newWorker(i+1, d.workerPool)
		w.start(d.wg)
		d.workers <- w
	}

	log.Printf("Started %d background workers", d.numWorkers)

	go d.dispatch()

	return d
}

// Stop dispatcher
func (d *dispatcher) stop() {
	log.Printf("Stopping background workers")
	defer func() {
		// clear WorkerPool
		for _ = range d.workerPool {
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

func (d *dispatcher) dispatch() {
	defer func() {
		d.stop()
	}()

	for {
		select {
		case job := <-d.jobPool:
			go func() {
				workerQueue := <-d.workerPool
				workerQueue <- job
			}()
		}
	}
}
