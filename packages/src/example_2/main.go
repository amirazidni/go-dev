package main

import (
	"fmt"
	"math"
)

func main() {
	// This is OK
	fmt.Println("Call Exp: ", math.Exp(1.0))

	// This will give us an error. Uncomment to try it
	// fmt.Println("Call expmulti: ", math.expmulti(0.1, 0.1, 1))
}
