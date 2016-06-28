package workerpool

import (
	"sync"
)

// Job is the function which should be executed in worker.
type Job func()

// Pool contains all information for the pool instance.
// Exports JobQueue used to send jobs to the pool
type Pool struct {
	JobQueue chan Job
	wg       *sync.WaitGroup
	d        *dispatcher
}

// New creates pool of workers.
//
// numWorkers - how many workers will be created for this pool
//
// queueLen - how many jobs can we accept until we block
//
// Returned object contains JobQueue reference, which you can use to send job to pool.
func New(numWorkers int, queueLen int) *Pool {
	jobQueue := make(chan Job, queueLen)

	return &Pool{
		JobQueue: jobQueue,
		wg:       &sync.WaitGroup{},
		d:        newDispatcher(numWorkers, jobQueue),
	}
}

// AddJob adds job to the que and increases WaitGroup counter
func (p *Pool) AddJob(job Job) {
	p.wg.Add(1)
	p.JobQueue <- job
}

// JobDone calls WaitGroup.Done() to remove the job
// after it is finished from the WaitGroup counter
func (p *Pool) JobDone() {
	p.wg.Done()
}

// Stop will wait for all the jobs to finish if any added with workerpool.AddJob()
// and then will stop the workers. If jobs are added directly to the que,
// Stop() will just shutdown the workers without waiting for any jobs already in the que.
func (p *Pool) Stop() {
	p.wg.Wait()
	p.d.quitChan <- true
}
