package main

import (
	"fmt"
	"time"
)

func main() {
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}
	}("Goroutine")

	fmt.Println("Hello")
	time.Sleep(500 * time.Millisecond)
}
