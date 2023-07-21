package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var wg = sync.WaitGroup{}
var cond = sync.NewCond(&locker)

func WaitCondition(val int) {
	defer wg.Done()
	wg.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", val)
	cond.L.Unlock()
}

func TestCondition(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Done", i)

			time.Sleep(1 * time.Second)
			cond.Signal() // send signal to tell goroutine to run
		}
	}()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	// wg.Wait()

	fmt.Println("===> Done ")
}
