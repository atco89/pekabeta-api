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

func (p ProductRepository) RetrieveOne(id uuid.UUID) (*domain.Product, error) {
	panic("implement me")
}

func (p ProductRepository) RetrieveAll() (*[]domain.Product, error) {
	products := &[]domain.Product{}
	err := p.Conn.Find(domain.Product{}).Error
	return products, err
}
