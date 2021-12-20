package main

import "fmt"

func main() {
	var triple [3]int
	var arr [5]string = [...]string{"java", "golang", "php", "python", "C"}

	fmt.Println(triple)
	fmt.Println(arr)

	var slice1 = arr[2:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = append(slice1, "C++")
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := make([]string, 0)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]string, 2, 5)
	slice4[0] = "item 1"
	slice4[1] = "item 2"
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
