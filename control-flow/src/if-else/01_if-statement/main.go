package main

import "fmt"

func multipleOfTen(x int) bool {
	if x%10 == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("20 is multiple of 10? ", multipleOfTen(20))
	fmt.Println("12 is multiple of 10? ", multipleOfTen(12))
}
