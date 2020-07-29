package rest

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"pekabeta/internal/application"
	"pekabeta/internal/domain"
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

func (a Api) CustomerRegistration(ctx echo.Context) error {
	customer := &domain.Customer{}
	if err := ctx.Bind(customer); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	if err := a.customer.Registration(&domain.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Password:    string(password),
	}); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}

	ctx.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (a Api) CustomerLogin(ctx echo.Context) error {
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
