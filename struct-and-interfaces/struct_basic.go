// Struct Basic

package main

import (
	"fmt"
)

type employee struct {
	name  string
	id    int
	email string
}

func main() {

	var emp0 = employee{} // struct initialization
	emp0.name = "Felix"   // Fill in the values
	emp0.id = 1234
	emp0.email = "felix@example.com"
	fmt.Println("employee 0:", emp0)

	// Filling Variables
	var emp1 = employee{"Ismail", 5555, "ismail@example.com"}
	fmt.Println("employee 1:", emp1)

	var emp2 = employee{email: "aryu@example.com", id: 9876, name: "Aryu"}
	fmt.Println("employee 2:", emp2)

	// Struct Pointer
	var emp3 *employee = &emp0                 //emp3 is a pointer to emp0 (Felix)
	fmt.Println("employee 3 name:", emp3.name) // Output: Felix

	emp3.name = "Vikra"
	fmt.Println("employee 0 name:", emp0.name)
	fmt.Println("employee 3 name:", emp3.name) //Both will output Vikra

}
