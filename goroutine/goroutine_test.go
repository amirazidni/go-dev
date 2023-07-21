package golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello wor")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Hai")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGo(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(10 * time.Second)
}

// channel -> return nilai goroutine
