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

func (o OrderRepository) Store(order domain.Order) error {
	return o.Conn.Create(order).Error
}

func (o OrderRepository) RetrieveOne(id uuid.UUID, customerId uuid.UUID) (domain.Order, error) {
	var order domain.Order
	err := o.Conn.First(&order, domain.Order{
		Base: domain.Base{
			ID: id,
		},
		CustomerID: customerId,
	}).Error
	return order, err
}

func (o OrderRepository) RetrieveAll(customerId uuid.UUID) ([]domain.Order, error) {
	var orders []domain.Order
	err := o.Conn.Find(&orders, domain.Order{
		CustomerID: customerId,
	}).Error
	return orders, err
}
