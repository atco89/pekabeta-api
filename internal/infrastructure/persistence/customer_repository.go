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

func (c CustomerRepository) EmailValidation(email *string) (domain.Validation, error) {
	panic("implement me")
}

func (c CustomerRepository) PhoneNumberValidation(phoneNumber *string) (domain.Validation, error) {
	panic("implement me")
}

func (c CustomerRepository) Registration(customer *domain.Customer) error {
	panic("implement me")
}

func (c CustomerRepository) Login(login domain.Login) (*domain.Customer, *domain.Error) {
	panic("implement me")
}

func (c CustomerRepository) Modify(customer *domain.Customer) error {
	panic("implement me")
}
