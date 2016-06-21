package workerpool

// Job is the function which should be executed in worker.
type Job func()

type Pool struct {
	JobQueue chan Job
	d        *dispatcher
}

// New creates pool of workers.
// numWorkers - how many workers will be created for this pool
// queLen - how many jobs can we accept until we block
//
// Returned object contains JobQueue reference, which you can use to send job to pool.
func New(numWorkers int, queLen int) *Pool {
	jobQueue := make(chan Job, queLen)

	pool := &Pool{
		JobQueue: jobQueue,
		d:        newDispatcher(numWorkers, jobQueue),
	}

	return pool
}

// Will release resources used by pool
func (p *Pool) Stop() {
	p.d.stop()
}
