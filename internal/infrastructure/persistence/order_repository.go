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
	tx := o.Conn.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&domain.Invoice{
		OrderId:       order.ID,
		Amount:        100.00,
		PaymentType:   domain.DEBIT_CARD,
		PaymentStatus: domain.COMPLETED,
	}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (o OrderRepository) RetrieveAll(customerId uuid.UUID) (*[]domain.Order, error) {
	orders := &[]domain.Order{}
	err := o.Conn.Find(orders, domain.Order{CustomerID: customerId}).Error
	return orders, err
}

func (o OrderRepository) RetrieveOne(id uuid.UUID, customerId uuid.UUID) (*domain.Order, error) {
	order := &domain.Order{}
	err := o.Conn.First(order, domain.Order{Base: domain.Base{ID: id}, CustomerID: customerId}).Error
	return order, err
}
