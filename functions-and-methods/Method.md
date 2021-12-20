# Method

Methods are functions tied to an object type (usually structs). The object who owns the method is placed between declaration `func` and the function itself.

```go
type employee struct {
	name  string
	id    int
}

func (e.employee) sayHello() {
    fmt.Println("Hello,", e.name)
}

func (e.employee) paySalary(amount int) {
    fmt.Printf("You paid %d to %s", amount, e.name)
}

func (e.employee) getID() int {
    return e.id
}
```

Different objects can have same method name. In fact, this will be important for making interfaces.

```go
func (e employee) sayHello() {
	fmt.Println("Hello,", e.name)
}

func (c cat) sayHello() {
	fmt.Println("Meow meow,", c.name)
}
```

## Method Pointers

Method can be declared as pointer. It is used for changing struct values, since if you try to change struct value using regular method the change will only persist inside that method. By using method pointer, the change is done at the reference variable.

```go
func (e employee) changeNameWrong(name string) {
    e.name = name
}

func (e *employee) changeNameCorrect(name string) {
    e.name = name
}
```