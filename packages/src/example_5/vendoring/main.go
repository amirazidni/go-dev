package main

import (
	"fmt"

	"github.com/pioz/faker"
)

func main() {
	faker.SetSeed(623)

	fmt.Println(faker.Username())
	fmt.Println(faker.String())
	fmt.Println(faker.IntInRange(1, 10))
	fmt.Println(faker.CurrencyName())
}
