// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/osousa/drato/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"fmt"
)

// Injectors from wire.go:

func injectStringer(s MyString) fmt.Stringer {
	return s
}
