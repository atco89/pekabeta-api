package persistence

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"pekabeta/internal/domain"
)

type ProductRepository struct {
	Conn *gorm.DB
}

func NewProductRepository(conn *gorm.DB) *ProductRepository {
	return &ProductRepository{Conn: conn}
}

func (p ProductRepository) RetrieveOne(id uuid.UUID) (domain.Product, error) {
	var product domain.Product
	err := p.Conn.First(&product, domain.Product{
		Base: domain.Base{
			ID: id,
		},
	}).Error
	return product, err
}

func (p ProductRepository) RetrieveAll() ([]domain.Product, error) {
	var products []domain.Product
	err := p.Conn.First(&products).Error
	return products, err
}
