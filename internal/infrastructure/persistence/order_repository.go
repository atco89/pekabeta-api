package persistence

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"pekabeta/internal/domain"
)

type OrderRepository struct {
	Conn *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) *OrderRepository {
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
