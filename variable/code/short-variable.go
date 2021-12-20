package main

import (
	"fmt"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
}
