package rest

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
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
	customer := &domain.Customer{}
	if err := ctx.Bind(customer); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return nil
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
	login := &api.Login{}
	if err := ctx.Bind(login); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	customer, err := a.customer.Login(domain.Login{
		Email:    login.Email,
		Password: []byte(login.Password),
	})

	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, customer)
}

func (a Api) ModifyCustomer(ctx echo.Context, id api.Id) error {
	customer := &domain.Customer{}
	if err := ctx.Bind(customer); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.MinCost)
	if err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return nil
	}

	modifiedCustomer := &domain.Customer{
		Base: domain.Base{
			ID: uuid.MustParse(string(id)),
		},
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Password:    string(password),
	}

	if err := a.customer.Modify(modifiedCustomer); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Response().WriteHeader(http.StatusNotFound)
			return err
		}

		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, modifiedCustomer)
}

func (a Api) RetrieveProduct(ctx echo.Context, id api.Id) error {
	customer, err := a.product.RetrieveOne(uuid.MustParse(string(id)))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Response().WriteHeader(http.StatusNotFound)
		}

		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, customer)
}

func (a Api) RetrieveProducts(ctx echo.Context) error {
	customer, err := a.product.RetrieveAll()
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Response().WriteHeader(http.StatusNotFound)
		}

		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	return ctx.JSON(http.StatusOK, customer)
}

func (a Api) CreateOrder(ctx echo.Context) error {
	order := &api.Order{}
	if err := ctx.Bind(order); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}

	orderDetails := &domain.Order{
		CustomerID:   uuid.MustParse(order.CustomerId),
		OrderStatus:  domain.OrderStatus(order.OrderStatus),
		ShippingType: domain.ShippingType(order.ShippingType),
	}

	if err := a.order.Store(orderDetails); err != nil {
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusCreated, orderDetails)
}

func (a Api) RetrieveOrder(ctx echo.Context, id api.Id) error {
	customerId := ctx.Request().Header.Get("Authorization")
	order, err := a.order.RetrieveOne(uuid.MustParse(string(id)), uuid.MustParse(customerId))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Response().WriteHeader(http.StatusNotFound)
		}
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, order)
}

func (a Api) RetrieveOrders(ctx echo.Context) error {
	customerId := ctx.Request().Header.Get("Authorization")
	orders, err := a.order.RetrieveAll(uuid.MustParse(customerId))
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			ctx.Response().WriteHeader(http.StatusNotFound)
		}
		ctx.Response().WriteHeader(http.StatusInternalServerError)
		return err
	}
	return ctx.JSON(http.StatusOK, orders)
}

func (a Api) RetrieveInvoice(ctx echo.Context, id api.Id) error {
	panic("implement me")
}

func (a Api) RetrieveInvoices(ctx echo.Context) error {
	panic("implement me")
}
