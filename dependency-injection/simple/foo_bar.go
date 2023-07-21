package simple

type Foo struct {
}

func NewFoo() *Foo {
	return &Foo{}
}

type Bar struct {
}

func NewBar() *Bar {
	return &Bar{}
}

// struct provider
type FooBar struct {
	*Foo
	*Bar
}
