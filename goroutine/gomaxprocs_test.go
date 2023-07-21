package golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	wg.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	runtime.GOMAXPROCS(12)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	wg.Wait()
}
