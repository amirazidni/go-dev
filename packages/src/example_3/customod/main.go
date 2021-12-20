package main

import (
	"fmt"
	gomath "math"

	"local.host/example/customod/math"
	"local.host/example/customod/math/list"
	"local.host/example/customod/text"
)

func main() {
	fmt.Println("Add number: 1.5 + 2.5 = ", math.Add(1.5, 2.5))
	fmt.Println("Multiply number: 1.5 + 2.5 = ", math.Multiply(1.5, 2.5))
	fmt.Println("Constant: PI: ", gomath.Pi)

	fmt.Println("Add all of the number in array: sum([3.0, 4.0, 1.0, 2.5]) = ", list.ListSum([]float64{3.0, 4.0, 1.0, 2.5}))
	fmt.Println("Multiply all of the number in array: mul([3.0, 4.0, 1.0, 2.5]) = ", list.ListMul([]float64{3.0, 4.0, 1.0, 2.5}))

	fmt.Println("Reverse: rev(Hello World) = ", text.ReverseText("Hello World"))
}
