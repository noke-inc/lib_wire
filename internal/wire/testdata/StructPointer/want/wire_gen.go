// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/github.com/noke-inc/lib_wiremann/wire/cmd/wire
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
