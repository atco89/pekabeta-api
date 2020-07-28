package persistence

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"pekabeta/internal/domain"
)

type OrderRepository struct {
	Conn *mongo.Database
}

func NewOrderRepository(conn *mongo.Database) *OrderRepository {
	return &OrderRepository{Conn: conn}
}

func (o OrderRepository) Store(order *domain.Order) error {
	panic("implement me")
}

func (o OrderRepository) RetrieveOne(id uuid.UUID) (*domain.Order, error) {
	panic("implement me")
}

func (o OrderRepository) RetrieveAll() (*[]domain.Order, error) {
	panic("implement me")
}
