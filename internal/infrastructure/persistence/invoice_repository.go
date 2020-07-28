package persistence

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"pekabeta/internal/domain"
)

type InvoiceRepository struct {
	Conn *mongo.Database
}

func NewInvoiceRepository(conn *mongo.Database) *InvoiceRepository {
	return &InvoiceRepository{Conn: conn}
}

func (i InvoiceRepository) RetrieveOne(id uuid.UUID) (*domain.Invoice, error) {
	panic("implement me")
}

func (i InvoiceRepository) RetrieveAll() (*[]domain.Invoice, error) {
	panic("implement me")
}
