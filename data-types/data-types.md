# Data Types

```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

## Zero values

 - Variables declared without an explicit initial value are given their zero value.

 - The zero value is:

    + 0 for numeric types,
    + false for the boolean type, and
    + "" (the empty string) for strings.

```
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
```

output :

```
0 0 false ""
```

url : [Zero values](https://play.golang.org/p/r9pa_UdbpkC)

## Type conversions

The expression T(v) converts the value v to the type T.

Some numeric conversions:
```
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```
Or, put more simply:
```
i := 42
f := float64(i)
u := uint(f)
```

url : [Type conversions](https://play.golang.org/p/h8cR9UhupHR)

Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.

## Types - Array

- Uninitialized array is prefilled with zero, or “” for string array
- Example array declaration:

```
package main

import "fmt"

func main() {
	var triple [3]int
	var arr [2]string = [2]string{"java","golang"}

	fmt.Println(triple)
	fmt.Println(arr)
}
```

output :
```
[0 0 0]
[java golang]
```

url : [Type Array](https://play.golang.org/p/z1Ffn1vwDbw)

## Types - Slice

- Similar like array but doesn’t have limits in declaration
- Uninitialized slice will result nil
- Slice component :
	+ pointer
	+ lenght
	+ capacity
- length <= capacity
- Before slice is ready to used, it is need to initialized using literal value or made using make built in function. This function will create and initialized slice

```
package main

import "fmt"

func main() {
	var triple [3]int
	var arr [5]string = [...]string{"java","golang","php","python","C"}

	fmt.Println(triple)
	fmt.Println(arr)

	var slice1 = arr[2:]
	fmt.Println("slice 1 : ", slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = append(slice1, "C++")
	fmt.Println("slice 2 : ", slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))
	
	slice3 := make([]string,0)
	fmt.Println("slice 3 : ", slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]string, 2, 5)
	slice4[0] = "item 1"
	slice4[1] = "item 2"
	fmt.Println("slice 4 : ", slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice5 := []int{1,2,3,4,5}
	fmt.Println("slice 5 : ", slice5)
	fmt.Println(len(slice5))
	fmt.Println(cap(slice5))
}
```

output :

```
[0 0 0]
[java golang php python C]
slice 1 :  [php python C]
3
3
slice 2 :  [php python C C++]
4
6
slice 3 :  []
0
0
slice 4 :  [item 1 item 2]
2
5
slice 5 :  [1 2 3 4 5]
5
5
```

url : [Type Slice](https://play.golang.org/p/6fA4emQkaQK)

## Types - Map

- Uninitialized map has value nil data and will cause a runtime panic
- Map value initialized using composite literal or using built-in function make

```
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
```

output :

```
map[key1:value1 key2:value2 key3:value3]
value1
value3
map[key19:value19 key21:value21]
map[key21:value21]
```

url : [Type Slice](https://play.golang.org/p/sIJSvn67lhr)