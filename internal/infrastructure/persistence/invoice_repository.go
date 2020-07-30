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

func (i InvoiceRepository) Store(invoice *domain.Invoice) error {
	return i.Conn.Create(invoice).Error
}

func (i InvoiceRepository) RetrieveAll(customerId uuid.UUID) (*[]domain.Invoice, error) {
	invoices := &[]domain.Invoice{}
	err := i.Conn.Find(invoices).Error
	return invoices, err
}

func (i InvoiceRepository) RetrieveOne(id uuid.UUID, customerId uuid.UUID) (*domain.Invoice, error) {
	invoice := &domain.Invoice{}
	err := i.Conn.First(invoice, domain.Invoice{Base: domain.Base{ID: id}}).Error
	return invoice, err
}
