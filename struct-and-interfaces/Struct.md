# Golang Structs

Struct in Golang is a group of variables and/or functions/methods grouped as a new data type.

## Declaration
Struct in Golang is declared by the keyword `type`

```go
type employee struct {
    name string
    id int
    email string
}
```


## Filling The Variables
You can fill the variables after struct initialization:

```go
var emp0 = employee{} // struct initialization
emp0.name = "Felix" // Fill in the values
emp0.id = 1234
emp0.email = "felix@example.com"
fmt.Println("employee 0:", emp0) // Output: {Felix, 1234, felix@example.com}
```

You can also fill the variables at initialization. If variable names is not mentioned, filling must be done in order.
```go
var emp1 = employee{"Ismail", 5555, "ismail@example.com"}
fmt.Println("employee 1:", emp1) // Output: {Ismail, 5555, ismail@example.com}
```

Same as above but since the variable name is mentioned, you can fill out of order.
```go
var emp2 = employee{email: "aryu@example.com", id: 9876, name: "Aryu"}
fmt.Println("employee 2:", emp2) // Output: {Aryu, 9876, aryu@example.com}
```

## Struct Pointers
Struct types have pointer values which can be stored as struct pointers.

```go
var emp3 *employee = &emp0 //emp3 is a pointer to emp0 (Felix)
fmt.Println("employee 3 name:", emp3.name) // Output: Felix

emp3.name = "Vikra"
fmt.Println("employee 0 name:", emp0.name)
fmt.Println("employee 3 name:", emp3.name)  //Both will output Vikra
```

## Embedded Structs

It is possible to put struct inside another struct. For example, here is a company struct with employee struct inside (employee become substruct):

```go
type employee struct {
	name  string
	id    int
	email string
}

type company struct {
	name    string
	address string
	employee
}
```

Struct can inherit variable from substruct but substruct cannot inherit from struct. ([Demonstration](./struct_embedded.go))