// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/osousa/drato/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

/* blockComment returns Foo and has a /*- style doc comment */
func blockComment() *Foo {
	foo := &Foo{}
	return foo
}

// lineComment returns Bar and has a //- style doc comment
func lineComment() *Bar {
	bar := &Bar{}
	return bar
}
