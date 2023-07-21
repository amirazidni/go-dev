package golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := sync.WaitGroup{}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)

			}
			wg.Done()
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter =", x)
}
