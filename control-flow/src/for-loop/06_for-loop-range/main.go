package main

import "fmt"

func main() {
	numbers := []int{2, 5, 7, 19, 20, 30}

	for idx, val := range numbers {
		fmt.Printf("%d = %d\n", idx, val)
	}

	// Ignore index
	for _, val := range numbers {
		fmt.Printf("%d\n", val)
	}

	// Map iteration
	keyvalue := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	for key, val := range keyvalue {
		fmt.Printf("key %s = %d\n", key, val)
	}
}
