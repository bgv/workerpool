package main

import (
	"fmt"
	"time"

	"github.com/bgv/workerpool"
)

func main() {
	// Number of workers, and Size of the job queue
	pool := workerpool.New(10, 50)

	// create and submit 10 jobs to the pool
	for i := 0; i < 100; i++ {
		count := i

		pool.JobQueue <- func() {
			fmt.Printf("I am job! Number %d\n", count)
		}
	}

	// dummy wait until jobs are finished
	time.Sleep(1 * time.Second)

	// release resources used by workerpool
	pool.Stop()
}
