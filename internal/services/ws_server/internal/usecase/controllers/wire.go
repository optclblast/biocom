//go:build wireinject
// +build wireinject

package controllers

import "github.com/google/wire"

func InitializeRootController() (*RootController, error) {
	wire.Build(
		NewRootController,
		NewAuthController,
	)
	return &RootController{}, nil
}
