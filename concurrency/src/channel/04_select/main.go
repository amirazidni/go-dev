package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- 10
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch2 <- 20
	}()

	for received := 0; ; {
		select {
		case v := <-ch1:
			fmt.Println("Receive from ch1 = ", v)
			received++
		case <-ch2:
			fmt.Println("Receive from ch2, don't need the value")
			received++
		default:
			if received == 2 {
				return
			}
		}
	}
}
