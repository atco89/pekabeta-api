package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type OrderRepository interface {
	Store(order *domain.Order) error
	RetrieveOne(id uuid.UUID) (*domain.Order, error)
	RetrieveAll() (*[]domain.Order, error)
}
