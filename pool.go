package workerpool

// Job is the function which should be executed in worker.
type Job func()

// Pool contains all information for the pool instance.
// Exports JobQueue used to send jobs to the pool
type Pool struct {
	JobQueue chan Job
	d        *dispatcher
}

// New creates pool of workers.
// numWorkers - how many workers will be created for this pool
// queueLen - how many jobs can we accept until we block
//
// Returned object contains JobQueue reference, which you can use to send job to pool.
func New(numWorkers int, queueLen int) *Pool {
	jobQueue := make(chan Job, queueLen)

	pool := &Pool{
		JobQueue: jobQueue,
		d:        newDispatcher(numWorkers, jobQueue),
	}

	return pool
}

// Stop will stop the workers and release resources used by pool
func (p *Pool) Stop() {
	p.d.stop()
}
