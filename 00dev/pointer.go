package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountry(address *Address) {
	address.Country = "Indonesia"
}

func Testmain8() {
	address1 := Address{"Subang", "Jabar", "Indo"}
	address2 := address1

	fmt.Println(address2)
	address2.City = "Bandung"

	fmt.Println(address2)
	fmt.Println(address1) //not changed because pass by value

	fmt.Println("=====================")
	//lets try pointer
	address3 := &address1 //pass by reference
	address3.City = "Tasik"
	fmt.Println(address1)
	fmt.Println(address3) //will show &, kasih tau itu adalah pointer

	fmt.Println("=====================")
	*address3 = Address{"Bnyms", "Central java", "Indo"} // pointing *
	fmt.Println(address1)
	fmt.Println(address3)

	fmt.Println("=====================")
	var address4 *Address = new(Address) // empty pointer struct
	fmt.Println(address4)

	fmt.Println("=====================")
	var alamat = Address{
		City:     "Pwt",
		Province: "Jateng",
		Country:  "",
	}
	ChangeCountry(&alamat)
	fmt.Println(alamat)

}
