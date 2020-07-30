package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type ProductRepository interface {
	RetrieveAll() (*[]domain.Product, error)
	RetrieveProductGroup(productGroup domain.ProductGroup) (*[]domain.Product, error)
	RetrieveOne(id uuid.UUID) (*domain.Product, error)
}
