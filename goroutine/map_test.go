package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, val int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	data.Store(val, val)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go AddToMap(data, i, wg)
	}
	wg.Wait()
	defer wg.Done()

	data.Range(func(k, v interface{}) bool {
		fmt.Println(k, ":", v)
		return true
	})

	fmt.Println("Done")
}
