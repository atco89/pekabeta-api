package rest

import (
	"github.com/labstack/echo/v4"
	"pekabeta/internal/application"
	"pekabeta/internal/interfaces/rest/api"
)

type Api struct {
	customer *application.CustomerInteractor
	product  *application.ProductInteractor
	order    *application.OrderInteractor
	invoice  *application.InvoiceInteractor
}

func NewApi(customer *application.CustomerInteractor,
	product *application.ProductInteractor,
	order *application.OrderInteractor,
	invoice *application.InvoiceInteractor) *Api {
	return &Api{
		customer: customer,
		product:  product,
		order:    order,
		invoice:  invoice,
	}
}

func (a Api) CustomerLogin(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) CustomerRegistration(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) ModifyCustomer(ctx echo.Context, id api.Id) error {
	panic("implement me")
}

func (a Api) RetrieveInvoices(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) RetrieveInvoice(ctx echo.Context, id api.Id) error {
	panic("implement me")
}

func (a Api) RetrieveOrders(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) CreateOrder(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) RetrieveOrder(ctx echo.Context, id api.Id) error {
	panic("implement me")
}

func (a Api) RetrieveProducts(ctx echo.Context) error {
	panic("implement me")
}

func (a Api) RetrieveProduct(ctx echo.Context, id api.Id) error {
	panic("implement me")
}
