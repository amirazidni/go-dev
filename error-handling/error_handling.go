package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (result float64, err error) {
	if y == 0 {
		err = errors.New("Unable to divide by zero")
		return 0, err
	}
	return float64(x / y), nil
}

func main() {
	num, err := divide(22, 11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(num)
	}
}
