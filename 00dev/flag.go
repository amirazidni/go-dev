package main

import (
	"flag"
	"fmt"
)

func Testmain5() {
	var name = flag.String("name", "anonymous", "type your name")
	var age = flag.Int64("age", 25, "type your age")

	flag.Parse()
	fmt.Printf("name\t: %s\n", *name)
	fmt.Printf("age\t: %d\n", *age)
}

// go run . -name="john wick" -age=28
