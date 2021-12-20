// Variadic Functions

package main

import (
	"fmt"
)

func sum(x ...int) (result int) {
	for _, i := range x {
		result += i
	}
	return
}

func nameAndSum(name string, x ...int) (result int) {
	fmt.Println("Hello,", name, ", the result is being calculated...")

	for _, i := range x {
		result += i
	}
	return
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))

	var num = []int{1, 2, 3, 4, 6}
	total2 := sum(num...)
	fmt.Println(total2)

	var num2 = []int{100, 20, 3, 4}
	total3 := nameAndSum("Felix", num2...)
	fmt.Println(total3)
}
