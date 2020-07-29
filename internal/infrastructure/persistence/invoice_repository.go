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

func (i InvoiceRepository) RetrieveOne(id uuid.UUID, orderId uuid.UUID) (domain.Invoice, error) {
	var invoice domain.Invoice
	err := i.Conn.First(invoice, domain.Invoice{
		Base: domain.Base{
			ID: id,
		},
		OrderId: orderId,
	}).Error
	return invoice, err
}

func (i InvoiceRepository) RetrieveAll(customerId uuid.UUID) ([]domain.Invoice, error) {
	panic("implement me")
}
