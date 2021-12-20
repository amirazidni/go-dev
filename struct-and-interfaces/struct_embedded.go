// Embedded Struct

package main

import (
	"fmt"
)

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

func main() {
	var comp0 = company{}
	comp0.name = "Samsung"
	comp0.employee.email = "felix@example.com"
	comp0.id = 1000 // This will succeed because company inherits ID from employee even though company doesn't have ID.
	comp0.address = "Wisma GKBI, Jakarta"

	// If struct and substruct both contain variable of similar name, they will not share values. Empty variable will remain empty.
	fmt.Println("Company name  :", comp0.name)
	fmt.Println("Employee name  :", comp0.employee.name)

	// This will succeed because company inherits email from employee.
	fmt.Println("Company email  :", comp0.email)
	fmt.Println("Employee email :", comp0.employee.email)

	// This will succeed because company inherits ID from employee.
	fmt.Println("ID:   ", comp0.id)

	fmt.Println("Company Address:", comp0.address)
	// This will fail because employee doesn't have address and employee cannot inherit company's address.
	// fmt.Println("Employee Address:", comp0.employee.address)

}
