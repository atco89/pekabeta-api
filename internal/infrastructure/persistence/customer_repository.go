package persistence

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
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
	customer := &domain.Customer{}
	if err := c.Conn.First(&customer, domain.Customer{Email: login.Email}).Error; err != nil {
		return nil, err
	}

	err := bcrypt.CompareHashAndPassword([]byte(customer.Password), login.Password)
	if err != nil {
		return nil, err
	}

	return customer, err
}

func (c CustomerRepository) Modify(customer *domain.Customer) error {
	return c.Conn.Model(&customer).UpdateColumns(domain.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
		Password:    customer.Password,
	}).Error
}
