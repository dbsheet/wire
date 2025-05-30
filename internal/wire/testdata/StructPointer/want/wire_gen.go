// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/dbsheet/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooBar() *FooBar {
	foo := provideFoo()
	bar := provideBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

func injectEmptyStruct() *Empty {
	empty := &Empty{}
	return empty
}
