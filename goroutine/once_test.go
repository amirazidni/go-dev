package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := &sync.Once{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			once.Do(OnlyOnce)
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Done", counter)
}
