// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/dbsheet/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFooer() Fooer {
	bar := provideBar()
	return bar
}
