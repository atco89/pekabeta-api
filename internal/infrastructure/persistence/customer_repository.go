package persistence

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"pekabeta/internal/domain"
)

type CustomerRepository struct {
	Conn *gorm.DB
}

func NewCustomerRepository(conn *gorm.DB) *CustomerRepository {
	return &CustomerRepository{Conn: conn}
}

func (c CustomerRepository) Registration(customer *domain.Customer) error {
	return c.Conn.Create(customer).Error
}

func (c CustomerRepository) Login(login domain.Login) (*domain.Customer, error) {
	panic("implement me")
}

func (c CustomerRepository) Update(id uuid.UUID, customer *domain.Customer) error {
	panic("implement me")
}
