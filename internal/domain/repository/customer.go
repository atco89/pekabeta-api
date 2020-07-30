package repository

import "pekabeta/internal/domain"

type CustomerRepository interface {
	EmailValidation(email *string) (domain.Validation, error)
	PhoneNumberValidation(phoneNumber *string) (domain.Validation, error)
	Registration(customer *domain.Customer) error
	Login(login domain.Login) (*domain.Customer, error)
	Modify(customer *domain.Customer) error
}
