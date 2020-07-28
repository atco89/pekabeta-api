package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type CustomerRepository interface {
	Registration(customer *domain.Customer) error
	Login(login domain.Login) (*domain.Customer, error)
	Update(id uuid.UUID, customer *domain.Customer) error
}
