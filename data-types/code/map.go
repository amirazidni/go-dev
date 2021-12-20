package main

import "fmt"

func main() {
	map1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	map1["key3"] = "value3"

	fmt.Println(map1)
	fmt.Println(map1["key1"])
	fmt.Println(map1["key3"])

	var map2 map[string]string = make(map[string]string)
	map2["key21"] = "value21"
	map2["key19"] = "value19"

	fmt.Println(map2)
	delete(map2, "key19")

	fmt.Println(map2)
}
