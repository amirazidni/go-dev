package main

import (
	"container/list"
	"fmt"
)

func Testmain1() {
	data := list.New()

	data.PushBack("Zidni")
	data.PushBack("iman")
	data.PushBack("S")
	data.PushFront("Setiadi")

	fmt.Println(data.Front().Value)
	fmt.Println(data.Back().Value)
}
