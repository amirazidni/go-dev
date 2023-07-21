package golangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
}

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

// direct write from reader (direct decoder)
func TestStreamDecoder(t *testing.T) {
	reader, _ := os.Open("Sample.json")
	decoder := json.NewDecoder(reader)

	customer := &Customer{}
	decoder.Decode(customer)

	fmt.Println(customer)
}

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("SampleOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "FirstName",
		MiddleName: "FirstName2",
		LastName:   "FirstName3",
	}

	encoder.Encode(customer)

	fmt.Println(customer)
}
