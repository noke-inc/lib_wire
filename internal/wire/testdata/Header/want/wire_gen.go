// This is a sample header file.
//
// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/github.com/noke-inc/lib_wiremann/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectFoo() Foo {
	foo := provideFoo()
	return foo
}
