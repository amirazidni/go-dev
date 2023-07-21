package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// go test -v -run=TestTimer

// menggunakan timer untuk menunggu, return timer.channel
func TestTimer(t *testing.T) {
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(time.Now())

	time := <-timer.C
	fmt.Println(time)

}

// menggunakan .After untuk menunggu, return channel
func TestAfter(t *testing.T) {
	channel := time.After(3 * time.Second)
	fmt.Println(time.Now())

	time := <-channel
	fmt.Println(time)

}

// menggunakan .AfterFunc untuk menunggu dan menjalankan sebuah func
func TestAfterFunc(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)
	// will run with goroutine, so use waitgroup
	time.AfterFunc(3*time.Second, func() {
		fmt.Println(time.Now())
		wg.Done()
	})

	fmt.Println(time.Now())

	wg.Wait()

}
