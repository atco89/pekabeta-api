//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"pekabeta/internal/application"
	"pekabeta/internal/domain/repository"
	"pekabeta/internal/infrastructure/persistence"
	"pekabeta/internal/interfaces/rest"
)

var repositorySet = wire.NewSet(
	persistence.NewCustomerRepository,
	persistence.NewProductRepository,
	persistence.NewOrderRepository,
	persistence.NewInvoiceRepository,
)

func CreateApi(db *gorm.DB) *rest.Api {
	panic(wire.Build(
		repositorySet,

		application.NewCustomerInteractor,
		wire.Bind(new(repository.CustomerRepository), new(*persistence.CustomerRepository)),

		application.NewProductInteractor,
		wire.Bind(new(repository.ProductRepository), new(*persistence.ProductRepository)),

		application.NewOrderInteractor,
		wire.Bind(new(repository.OrderRepository), new(*persistence.OrderRepository)),

		application.NewInvoiceInteractor,
		wire.Bind(new(repository.InvoiceRepository), new(*persistence.InvoiceRepository)),

		rest.NewApi,
	))
}
