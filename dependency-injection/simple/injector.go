//go:build wireinject
// +build wireinject

// up there is to tell go, this is wire injection and ignored by build tool

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

// injector
func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

// how to run wire
// wire gen "go-dev/dependency-injection/simple"
// or run "wire" at package folder with injector

// multi binding injector
func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

// provider set or grouping
var fooSet = wire.NewSet(NewFooRepository, NewFooService)

// provider set or grouping
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// Salah
//func InitializedHelloService() *HelloService {
//	wire.Build(NewHelloService, NewSayHelloImpl)
//	return nil
//}

// binding interface
var helloSet = wire.NewSet(
	NewSayHelloImpl, //provider set
	wire.Bind(new(SayHello), new(*SayHelloImpl)), // kalau perlu pertama, kirim kedua
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

// struct provider
func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

// binding value
func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

// binding interface value
func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin)) // kalau ada yg butuh reader, return stdin
	return nil
}

// struct field provider
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		// buat provider dari field data Application, menggunakan field Configuration
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// clean up function
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
