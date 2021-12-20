// Empty Interface

package main

import "fmt"

func main() {
	var i interface{}
	print(i)

	i = 123
	print(i)

	i = 3.14
	print(i)

	i = "Felix"
	print(i)

	var j int
	i = &j
	print(i)
}

func print(i interface{}) {
	fmt.Println(i)
}
