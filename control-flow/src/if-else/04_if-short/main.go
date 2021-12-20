package main

import "fmt"

func main() {
	if number := 21; number%10 == 0 {
		fmt.Printf("%d is multiple of 10\n", number)
	} else {
		fmt.Printf("%d is not multiple of 10\n", number)
	}

	// var number is no longer valid here
}
