package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)

	fmt.Println("Hi")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(wg)
	}

	wg.Wait()
	fmt.Println("Done")
}
