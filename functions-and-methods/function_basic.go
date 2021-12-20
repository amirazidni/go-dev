// Basic Function

package main

import "fmt"

func meow() {
	fmt.Println("meow!!")
}

func square(x int) {
	fmt.Println(x * x)
}

func introduce(name string, age int) {
	fmt.Printf("Hi, my name is %s, I am %d years old.\n", name, age)
}

func volume(x, y, z int, fn, ln string) {
	fmt.Println(x * y * z)
	fmt.Println(fn, ln)
}

func main() {
	square(4)
	meow()
	introduce("Felix", 23)
	volume(2, 3, 4, "Felix", "Andrean")
}
