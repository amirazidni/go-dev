package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)

		channel <- "Zidni aja"
		fmt.Println("Selesai kirim ke channel")
	}()
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "This is your response"

}
func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveResponse(channel)
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// send only
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Zidni Iman Sholihati In"
}

// get/received only
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestIOChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// Buffered channel = penampungan data
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Zidni"

	fmt.Println(<-channel)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		// jika ga ditutup maka deadlock karena baris 87 akan cek terus
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Receive", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	chan1 := make(chan string)
	chan2 := make(chan string)
	defer close(chan1)
	defer close(chan2)

	go GiveResponse(chan1)
	go GiveResponse(chan2)

	counter := 0
	for {
		select {
		case data := <-chan1:
			fmt.Println("From channel 1", data)
			counter++
		case data := <-chan2:
			fmt.Println("From channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}
