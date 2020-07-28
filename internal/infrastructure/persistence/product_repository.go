package persistence

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"pekabeta/internal/domain"
)

type ProductRepository struct {
	Conn *mongo.Database
}

func NewProductRepository(conn *mongo.Database) *ProductRepository {
	return &ProductRepository{Conn: conn}
}

func (p ProductRepository) RetrieveOne(id uuid.UUID) (*domain.Product, error) {
	panic("implement me")
}

func (p ProductRepository) RetrieveAll() (*[]domain.Product, error) {
	panic("implement me")
}
