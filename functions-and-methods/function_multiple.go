// Multiple Returns

package main

import (
	"fmt"
)

func rectangle(x, y int) (int, int) {
	area := x * y
	perimeter := (x + y) * 2
	return area, perimeter
}

func squareAndName(x int) (int, string) {
	result := x * x
	name := "Felix"
	return result, name
}

func main() {
	a, p := rectangle(2, 5)
	fmt.Printf("Area = %d, Perimeter = %d\n", a, p)

	number, myName := squareAndName(7)
	fmt.Println(number, myName)
}
