// +build ignore
package main

import (
	"fmt"

	"github.com/bgv/workerpool"
)

func main() {
	// Number of workers, and Size of the job queue
	simplepool := workerpool.New(10, 50)

	// create and submit 100 jobs to the pool
	for i := 0; i < 100; i++ {
		count := i

		simplepool.JobQueue <- func() {
			fmt.Printf("I am job! Number %d\n", count)
		}
	}

	// Wait for all jobs to finish and stop the workers
	simplepool.Stop()
}
