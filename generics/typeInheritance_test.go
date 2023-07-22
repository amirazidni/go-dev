package generic_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Todo interface {
	GetTodo() string
}

// Type parameter inheritance
func GetTodo[T Todo](param T) string {
	return param.GetTodo()
}

type Note interface {
	GetTodo() string
	GetNote() string
}

type MyNote struct {
	Title string
}

func (m *MyNote) GetTodo() string {
	return m.Title
}

func (m *MyNote) GetNote() string {
	return m.Title
}

type Book interface {
	GetTodo() string
	GetBook() string
}

type MyBook struct {
	Cover string
}

func (m *MyBook) GetTodo() string {
	return m.Cover
}

func (m *MyBook) GetBook() string {
	return m.Cover
}

func TestGetTodo(t *testing.T) {
	assert.Equal(t, "Amira", GetTodo[Note](&MyNote{"Amira"}))
	assert.Equal(t, "Amira", GetTodo[Book](&MyBook{"Amira"}))
}
