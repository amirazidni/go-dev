package main

import "fmt"

func main() {
	for n := 0; n < 100; n++ {
		if n == 10 {
			break
		}

		if n%2 == 0 {
			continue
		}

		fmt.Printf("%d ", n)
	}
}
