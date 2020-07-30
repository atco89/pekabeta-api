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

func (p ProductRepository) RetrieveAll() (*[]domain.Product, error) {
	products := &[]domain.Product{}
	err := p.Conn.Find(products).Error
	return products, err
}

func (p ProductRepository) RetrieveProductGroup(productGroup domain.ProductGroup) (*[]domain.Product, error) {
	products := &[]domain.Product{}
	err := p.Conn.Find(products, domain.Product{ProductGroup: productGroup}).Error
	return products, err
}

func (p ProductRepository) RetrieveOne(id uuid.UUID) (*domain.Product, error) {
	products := &domain.Product{}
	err := p.Conn.First(products, domain.Product{Base: domain.Base{ID: id}}).Error
	return products, err
}
