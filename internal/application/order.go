package application

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
	"pekabeta/internal/domain/repository"
)

type OrderInteractor struct {
	repository repository.OrderRepository
}

func NewOrderInteractor(repo repository.OrderRepository) *OrderInteractor {
	return &OrderInteractor{
		repository: repo,
	}
}

func (o OrderInteractor) Store(order *domain.Order) error {
	return o.repository.Store(order)
}

func (o OrderInteractor) RetrieveOne(id uuid.UUID) (*domain.Order, error) {
	return o.repository.RetrieveOne(id)
}

func (o OrderInteractor) RetrieveAll() (*[]domain.Order, error) {
	return o.repository.RetrieveAll()
}
