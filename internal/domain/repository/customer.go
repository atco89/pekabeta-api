package repository

import "pekabeta/internal/domain"

type CustomerRepository interface {
	Registration(customer *domain.Customer) error
	Login(login domain.Login) (domain.Customer, error)
	Modify(customer domain.Customer) error
}
