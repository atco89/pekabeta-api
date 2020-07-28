package repository

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
)

type InvoiceRepository interface {
	RetrieveOne(id uuid.UUID) (*domain.Invoice, error)
	RetrieveAll() (*[]domain.Invoice, error)
}
