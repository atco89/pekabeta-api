package persistence

import (
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

func (c CustomerRepository) Login(login domain.Login) (domain.Customer, error) {
	var customer domain.Customer
	err := c.Conn.First(&customer, domain.Customer{
		Email:    login.Email,
		Password: login.Password,
	}).Error
	return customer, err
}

func (c CustomerRepository) Modify(customer domain.Customer) error {
	return c.Conn.Update(customer).Error
}
