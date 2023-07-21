package main

import (
	"fmt"
)

// type assertion adalah mengubah tipedata atau casting
func random() interface{} {
	return false
}
func main() {
	result := random()
	// resultString := result.(string) //convert here from empty interface{}
	// fmt.Println(resultString)

	// use switch for avoid panic
	switch value := result.(type) {
	case string:
		fmt.Println("Value ", value, "is string")
	case int:
		fmt.Println("Value ", value, "is int")
	default:

		fmt.Println("Value ", value)
	}
}
