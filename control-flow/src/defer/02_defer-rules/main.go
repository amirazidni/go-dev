package main

import "fmt"

func deferEvaluation() {
	i := 0
	// 0 will be printed instead of 2, despite being executed at the end.
	defer fmt.Println(i)
	i += 2
}

func deferOrder() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Third defer")
}

func deferReturn() (i int) {
	i = 3
	defer func() { i += 2 }()
	return i
}

func main() {
	// This will output: 0
	deferEvaluation()

	// This will execute from the last defer (LIFO)
	deferOrder()

	// This will output: 2
	fmt.Println(deferReturn())
}
