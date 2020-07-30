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

func (c CustomerInteractor) EmailValidation(email *string) (domain.Validation, error) {
	return c.repository.EmailValidation(email)
}

func (c CustomerInteractor) PhoneNumberValidation(phoneNumber *string) (domain.Validation, error) {
	return c.repository.PhoneNumberValidation(phoneNumber)
}

func (c CustomerInteractor) Registration(customer *domain.Customer) error {
	return c.repository.Registration(customer)
}

func (c CustomerInteractor) Login(login domain.Login) (*domain.Customer, *domain.Error) {
	panic("implement me")
}

func (c CustomerInteractor) Modify(customer *domain.Customer) error {
	panic("implement me")
}
