//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/fingo-martpedia/fingo-wallet/helpers"
	"github.com/fingo-martpedia/fingo-wallet/internal/controller"
	"github.com/fingo-martpedia/fingo-wallet/internal/interfaces"
	"github.com/fingo-martpedia/fingo-wallet/internal/repository"
	"github.com/fingo-martpedia/fingo-wallet/internal/services"
	"github.com/google/wire"
	"gorm.io/gorm"
)

type Dependency struct {
	WalletController interfaces.IWalletController
}

func provideDB() *gorm.DB {
	return helpers.DB
}

func InitDependency() Dependency {
	wire.Build(
		provideDB,

		repository.NewWalletRepository,
		wire.Bind(new(interfaces.IWalletRepository), new(*repository.WalletRepository)),

		services.NewWalletService,
		wire.Bind(new(interfaces.IWalletService), new(*services.WalletService)),

		controller.NewWalletController,
		wire.Bind(new(interfaces.IWalletController), new(*controller.WalletController)),

		wire.Struct(new(Dependency), "*"),
	)
	return Dependency{}
}
