// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/dbsheet/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectBar() (*Bar, func()) {
	foo, cleanup := provideFoo()
	bar, cleanup2 := provideBar(foo)
	return bar, func() {
		cleanup2()
		cleanup()
	}
}
