package main

import "fmt"

func main() {
	defer fmt.Println("Executed last")

	fmt.Println("First!")
}
