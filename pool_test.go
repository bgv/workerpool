package workerpool

import (
	"sync/atomic"
	"testing"
)

func Test_NewPool(t *testing.T) {
	pool := New(100, 1000)

	iterations := 1000000
	var counter uint64 = 0

	for i := 0; i < iterations; i++ {
		arg := uint64(1)

		job := func() {
			atomic.AddUint64(&counter, arg)
			if arg != uint64(1) {
    			t.Errorf("Expected: %d, got: %d", uint64(1), arg)
			}
		}

		pool.JobQueue <- job
	}

    pool.Stop()

	counterFinal := atomic.LoadUint64(&counter)
	if counterFinal != uint64(iterations){
    	t.Errorf("Expected: %d, got: %d", uint64(iterations), counterFinal)
	}
}
