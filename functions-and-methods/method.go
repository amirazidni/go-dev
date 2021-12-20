// Method

package main

import (
	"fmt"
)

type employee struct {
	name string
	id   int
}

type cat struct {
	name string
}

func (e employee) sayHello() {
	fmt.Println("Hello,", e.name)
}

func (c cat) sayHello() {
	fmt.Println("Meow meow,", c.name)
}

func (e employee) paySalary(amount int) {
	fmt.Printf("You paid %d to %s\n", amount, e.name)
}

func (e employee) getID() int {
	return e.id
}

// Method Pointers

func (e employee) changeNameWrong(name string) {
	e.name = name
	fmt.Println("Employee Name inside method:", e.name)
}

func (e *employee) changeNameCorrect(name string) {
	e.name = name
	fmt.Println("Employee Name inside method:", e.name)
}

func main() {
	var emp0 = employee{"Felix", 1234}
	var cat0 = cat{"Kitty"}

	emp0.sayHello()
	cat0.sayHello()

	emp0.paySalary(500)
	empID := emp0.getID()
	fmt.Println("Employee ID:", empID)

	emp0.changeNameWrong("Vikra")
	fmt.Println("Employee Name:", emp0.name)

	emp0.changeNameCorrect("Gary")
	fmt.Println("Employee Name:", emp0.name)
}
