package main

import "fmt"

func sum(nums []int, ch chan int) {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	ch <- sum
}

func main() {
	numbers := []int{4, 2, 1, 5, 3, 10, 9, 8}
	ch := make(chan int)

	go sum(numbers[:len(numbers)/2], ch)
	go sum(numbers[len(numbers)/2:], ch)

	// Using loop:
	//
	// sum := 0
	// for i := 2; i > 0; i-- {
	// 	sum += <-ch
	// }

	// or, assign manually
	sum := <-ch + <-ch

	fmt.Println("Sum = ", sum)
}
