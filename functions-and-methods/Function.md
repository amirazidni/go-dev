# Function

Function is a block of code with callable name. The simplest one is a function without parameter and return.
```go
func meow() {
	fmt.Println("meow")
}
```

## Parameters

Parameters are put inside parentheses.

```go
func square(x int) {
	fmt.Println(x*x)
}
```

Multiple parameters can be the same or different types.

```go
func introduce(name string, age int) {
    fmt.Printf("Hi, my name is %s, I am %d years old.\n", name, age)
}
```

Multiple parameters of the same type can be grouped together.

```go
func volume(x int, y int, z int)
func volume(x, y, z int)

func volumeAndName(x int, y int, z int, firstName string, lastName string)
func volumeAndName(x, y, z int, firstName, lastName string)
```

## Return Values

Return value type is declared at the start of the function before the curl brackets.

```go
func square(x int) int {
    result := x*x
	return result
}
```

You can declare the variable name from the start too.

```go
func square(x int) (result int) {
	result = x*x
	return
}
```

Return can also be used to end a void function early.

```go
func divide(x, y int) {
	if y == 0 {
        fmt.Println("Unable to divide by zero.")
        return
    }
    fmt.Println(x/y)
}
```

## Multiple Return Values

Golang can have function that returns multiple values.

```go
func rectangle(x, y int) (int, int) {
    area := x*y
    perimeter := (x+y)*2
	return area, perimeter
}
```
The return values can be of different type too

```go
func squareAndName(x int) (int, string) {
    result := x*x
    name := "Felix"
    return result, name
}
```

When declaring variable names from the start, you can group together return values of the same type. (just like parameters)

```go
func rectangle(x, y int) (area, perimeter int) {
    area = x*y
    perimeter = (x+y)*2
	return
}
```

## Variadic Function

Variadic function accepts any number of parameters of the same data type. To declare one, add leading `...` at the parameter type.

```go
func sum(x ...int) (result int) {
    for _, i := range x {
        result += x
    }
    return
}
```

You can feed a variadic function either with multiple parameters or with a slice (add trailing `...` at the parameter name)

```go
total1 := sum(1,2,3,4,5)

var numbers = []int{1, 2, 3, 4, 5}
total2 := sum(numbers...)
```

You can combine regular and variadic parameter in a function, but variadic parameter must be the last.

```go
func nameAndSum(name string, x ...int) (result int) {
    fmt.Println("Hello, ", name, "the result is being calculated...")

    for _, i := range x {
        result += x
    }
    return
}
```