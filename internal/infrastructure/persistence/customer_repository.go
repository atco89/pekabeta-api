package persistence

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"pekabeta/internal/domain"
)

type CustomerRepository struct {
	Conn *mongo.Database
}

func NewCustomerRepository(conn *mongo.Database) *CustomerRepository {
	return &CustomerRepository{Conn: conn}
}

func (c CustomerRepository) Registration(customer *domain.Customer) error {
	panic("implement me")
}

func (c CustomerRepository) Login(login domain.Login) (*domain.Customer, error) {
	panic("implement me")
}

func (c CustomerRepository) Update(id uuid.UUID, customer *domain.Customer) error {
	panic("implement me")
}
