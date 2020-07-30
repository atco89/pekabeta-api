package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type OrderRepository interface {
	Store(order *domain.Order) error
	RetrieveAll(customerId uuid.UUID) (*[]domain.Order, error)
	RetrieveOne(id uuid.UUID, customerId uuid.UUID) (*domain.Order, error)
}
