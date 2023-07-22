package generic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Bag[T any] []T

func PrintBag[T any](b Bag[T]) {
	for _, v := range b {
		fmt.Printf("v: %v\n", v)
	}
}

func TestBStr(t *testing.T) {
	names := Bag[string]{"Zidni", "amira"}
	PrintBag(names)
}

func TestBInt(t *testing.T) {
	names := Bag[int]{42, 12}
	PrintBag(names)
}

type Data[T any] struct {
	First T
	Last  Bag[string]
}

func TestData(t *testing.T) {
	names := Bag[string]{"Zidni", "amira"}
	data := Data[string]{
		First: "hai",
		Last:  names,
	}
	fmt.Printf("data: %v\n", data)
	fmt.Printf("sayhello: %v \n", data.SayHello(names[len(names)-2]))
}

func (r *Data[T]) SayHello(name string) string {
	fmt.Printf("name: %v\n", name)
	return "Hello " + name
}

// type parameter hanya untuk function, tidak di method

// In Line type declaration
func FindMin[T interface{ int | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 100, FindMin(100, 200))
	assert.Equal(t, int64(100), FindMin(int64(100), int64(200)))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Zidni", "Iman", "Sholihati",
	}
	f := GetFirst(names)
	fmt.Printf("f: %v\n", f)
	assert.Equal(t, "Zidni", f)
}
