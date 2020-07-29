package application

import (
	"pekabeta/internal/domain"
	"pekabeta/internal/domain/repository"
)

type CustomerInteractor struct {
	repository repository.CustomerRepository
}

func NewCustomerInteractor(repo repository.CustomerRepository) *CustomerInteractor {
	return &CustomerInteractor{
		repository: repo,
	}
}

func (c CustomerInteractor) Registration(customer *domain.Customer) error {
	return c.repository.Registration(customer)
}

func (c CustomerInteractor) Login(login domain.Login) (domain.Customer, error) {
	return c.repository.Login(login)
}

func (c CustomerInteractor) Modify(customer domain.Customer) error {
	return c.repository.Modify(customer)
}
