package generic_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// single type parameter
func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

// multiple type parameter
func MultiParam[T1, T2 any](param1 T1, param2 T2) T2 {
	fmt.Println(param1, param2)
	return param2
}

// comparable type param
func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestSample(t *testing.T) {
	var result string = Length[string]("Amira")
	assert.Equal(t, "Amira", result)

	// can use this
	result = Length("Amira 2")
	assert.Equal(t, "Amira 2", result)

	var resultNumber int = Length[int](100)
	assert.Equal(t, 100, resultNumber)
}

func TestMultiParam(t *testing.T) {
	var result int = MultiParam[string, int]("Amira", 100)
	assert.Equal(t, 100, result)

	// can use this
	result = MultiParam("Amira 2", 100)
	assert.Equal(t, 100, result)

	MultiParam(100, "Amira")
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame("Eko", "Eko"))
	assert.True(t, IsSame(100, 100))
	assert.True(t, IsSame(true, true))
}
