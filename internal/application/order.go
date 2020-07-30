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
	panic("implement me")
}

func (o OrderInteractor) RetrieveAll(customerId uuid.UUID) (*[]domain.Order, error) {
	panic("implement me")
}

func (o OrderInteractor) RetrieveOne(id uuid.UUID, customerId uuid.UUID) (*domain.Order, error) {
	panic("implement me")
}
