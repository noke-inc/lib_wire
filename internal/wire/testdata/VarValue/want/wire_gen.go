// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/pantsmann/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func injectedMessage() string {
	string2 := _wireStringValue
	return string2
}

var (
	_wireStringValue = msg
)
