package persistence

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"pekabeta/internal/domain"
)

type InvoiceRepository struct {
	Conn *gorm.DB
}

func NewInvoiceRepository(conn *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{Conn: conn}
}

func (i InvoiceRepository) RetrieveOne(id uuid.UUID) (*domain.Invoice, error) {
	panic("implement me")
}

func (i InvoiceRepository) RetrieveAll() (*[]domain.Invoice, error) {
	panic("implement me")
}
