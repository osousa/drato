// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/osousa/drato/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func inject(foo *Foo) *Bar {
	bar := NewBar(foo)
	return bar
}
