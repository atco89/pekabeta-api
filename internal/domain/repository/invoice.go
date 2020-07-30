package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type InvoiceRepository interface {
	Store(invoice *domain.Invoice) error
	RetrieveAll(customerId uuid.UUID) (*[]domain.Invoice, error)
	RetrieveOne(id uuid.UUID, customerId uuid.UUID) (*domain.Invoice, error)
}
