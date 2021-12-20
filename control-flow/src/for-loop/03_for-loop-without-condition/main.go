package main

import "fmt"

func main() {
	for j := 0; ; j++ {
		fmt.Printf("%d ", j)
		if j == 10 {
			break
		}
	}
}
