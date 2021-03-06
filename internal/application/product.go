package application

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
	"pekabeta/internal/domain/repository"
)

type ProductInteractor struct {
	repository repository.ProductRepository
}

func NewProductInteractor(repo repository.ProductRepository) *ProductInteractor {
	return &ProductInteractor{
		repository: repo,
	}
}

func (p ProductInteractor) RetrieveAll() (*[]domain.Product, error) {
	return p.repository.RetrieveAll()
}

func (p ProductInteractor) RetrieveProductGroup(productGroup domain.ProductGroup) (*[]domain.Product, error) {
	return p.repository.RetrieveProductGroup(productGroup)
}

func (p ProductInteractor) RetrieveOne(id uuid.UUID) (*domain.Product, error) {
	return p.repository.RetrieveOne(id)
}
