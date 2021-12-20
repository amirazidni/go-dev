package main

import (
	"fmt"
	"sync"
	"time"
)

func task(wg *sync.WaitGroup, id int) {
	defer wg.Done()

	fmt.Println("Start task ", id)
	time.Sleep(time.Second)
	fmt.Println("End task ", id)
}

func main() {
	var wg sync.WaitGroup

	// wg.Add(5)
	for t := 0; t < 5; t++ {
		wg.Add(1)
		go task(&wg, t)
	}

	wg.Wait()
	fmt.Println("All tasks done")
}
