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

func (c CustomerRepository) Login(login domain.Login) (*domain.Customer, error) {
	customer := &domain.Customer{}
	if err := c.Conn.First(customer, domain.Customer{Email: login.Email}).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(customer.Password), login.Password); err != nil {
		return nil, err
	}
	return customer, nil
}

func (c CustomerRepository) Modify(customer *domain.Customer) error {
	modifiedCustomer := domain.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		Email:       customer.Email,
		PhoneNumber: customer.PhoneNumber,
	}

	if customer.Password != "" {
		modifiedCustomer.Password = customer.Password
	}
	return c.Conn.Model(customer).UpdateColumns(modifiedCustomer).Error
}
