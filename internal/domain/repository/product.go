package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type ProductRepository interface {
	RetrieveOne(id uuid.UUID) (*domain.Product, error)
	RetrieveAll() (*[]domain.Product, error)
}
