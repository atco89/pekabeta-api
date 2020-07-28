package rest

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	customer := &api.Customer{}
	if err := ctx.Bind(customer); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	if err := a.customer.Registration(&domain.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Password:    string(password),
	}); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	ctx.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (a Api) CustomerLogin(ctx echo.Context) error {
	login := &api.Login{}
	if err := ctx.Bind(login); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	customer, err := a.customer.Login(domain.Login{
		Email:    login.Email,
		Password: login.Password,
	})

	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, customer)
}

func (a Api) ModifyCustomer(ctx echo.Context, id api.Id) error {
	customer := &api.Customer{}
	if err := ctx.Bind(customer); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	if err = a.customer.Update(uuid.MustParse(string(id)), &domain.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Password:    string(password),
	}); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	ctx.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (a Api) RetrieveInvoices(ctx echo.Context) error {
	invoices, err := a.invoice.RetrieveAll()
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, invoices)
}

func (a Api) RetrieveInvoice(ctx echo.Context, id api.Id) error {
	invoice, err := a.invoice.RetrieveOne(uuid.MustParse(string(id)))
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, invoice)
}

func (a Api) CreateOrder(ctx echo.Context) error {
	order := &api.Order{}
	if err := ctx.Bind(order); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	if err := a.order.Store(&domain.Order{}); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	ctx.Response().WriteHeader(http.StatusCreated)
	return nil
}

func (a Api) RetrieveOrders(ctx echo.Context) error {
	orders, err := a.order.RetrieveAll()
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (a Api) RetrieveOrder(ctx echo.Context, id api.Id) error {
	order, err := a.order.RetrieveOne(uuid.MustParse(string(id)))
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, order)
}

func (a Api) RetrieveProducts(ctx echo.Context) error {
	products, err := a.product.RetrieveAll()
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, products)
}

func (a Api) RetrieveProduct(ctx echo.Context, id api.Id) error {
	product, err := a.product.RetrieveOne(uuid.MustParse(string(id)))
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, product)
}
