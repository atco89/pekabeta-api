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
	mail := *email
	validation := domain.Validation{}

	validation.IsValid = false

	if err := c.Conn.First(&domain.Customer{}, domain.Customer{Email: mail}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			validation.IsValid = true
			return validation, nil
		}
		return validation, err
	}

	return validation, nil
}

func (c CustomerRepository) PhoneNumberValidation(phoneNumber *string) (domain.Validation, error) {
	phone := *phoneNumber
	validation := domain.Validation{}

	validation.IsValid = false

	if err := c.Conn.First(&domain.Customer{}, domain.Customer{PhoneNumber: phone}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			validation.IsValid = true
			return validation, nil
		}
		return validation, err
	}

	return validation, nil
}

func (c CustomerRepository) Registration(customer *domain.Customer) error {
	return c.Conn.Create(customer).Error
}

func (c CustomerRepository) Login(login domain.Login) (*domain.Customer, *domain.Error) {
	panic("implement me")
}

func (c CustomerRepository) Modify(customer *domain.Customer) error {
	panic("implement me")
}
