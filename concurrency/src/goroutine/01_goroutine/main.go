package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d %s\n", i, s)
	}
}

func main() {
	go say("goroutine")
	say("hello")
}
