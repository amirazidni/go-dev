package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	total := 0
	for i := 0; i < 1000; i++ {
		total += i
	}

	fmt.Println("Without goroutine: total = ", total)

	var mutex = &sync.Mutex{}

	total = 0
	for i := 0; i < 1000; i++ {
		go func(n int) {
			mutex.Lock()
			total += n
			mutex.Unlock()
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("With goroutine: total = ", total)
}
