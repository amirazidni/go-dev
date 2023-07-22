package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type NumberTypeSet interface {
	~int | int8 | int16 | int32 | int64 |
		float32 | float64
}

func Min[T NumberTypeSet](one, two T) T {
	if one < two {
		return one
	}
	return two
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min[int64](100, 200))
	assert.Equal(t, float32(100), Min[float32](100, 200))
	assert.Equal(t, Age(100), Min[Age](100, 200)) // error constraint if not use tilde
}

// tidak bisa menggunakan type declaration
type Age int

func TestTypeInference(t *testing.T) {
	assert.Equal(t, 100, Min(100, 200))
	assert.Equal(t, int64(100), Min(int64(100), int64(200)))
	assert.Equal(t, float64(100.0), Min(100.0, 200.0))
	assert.Equal(t, Age(100), Min(Age(100), Age(200)))
}

// generic interface

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(val T)
}

func ChangeValue[T any](param GetterSetter[T], val T) T {
	param.SetValue(val)
	return param.GetValue()
}

type MyData[T any] struct {
	Val T
}

func (d *MyData[T]) GetValue() T {
	return d.Val
}

func (d *MyData[T]) SetValue(val T) {
	d.Val = val
}

func TestGenericAny(t *testing.T) {
	myDat := MyData[string]{}

	res := ChangeValue[string](&myDat, "Zidni")
	assert.Equal(t, "Zidni", res)
	assert.Equal(t, "Zidni", myDat.Val)
}
