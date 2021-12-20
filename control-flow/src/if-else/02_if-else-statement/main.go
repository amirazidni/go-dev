package main

import "fmt"

func main() {
	number := 21
	if number%10 == 0 {
		fmt.Printf("%d is multiple of 10\n", number)
	} else {
		fmt.Printf("%d is not multiple of 10\n", number)
	}
}
