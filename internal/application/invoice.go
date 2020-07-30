package application

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
	"pekabeta/internal/domain/repository"
)

type InvoiceInteractor struct {
	repository repository.InvoiceRepository
}

func NewInvoiceInteractor(repo repository.InvoiceRepository) *InvoiceInteractor {
	return &InvoiceInteractor{
		repository: repo,
	}
}

func (i InvoiceInteractor) RetrieveAll(customerId uuid.UUID) ([]domain.Invoice, error) {
	panic("implement me")
}

func (i InvoiceInteractor) RetrieveOne(id uuid.UUID, customerId uuid.UUID) (domain.Invoice, error) {
	panic("implement me")
}
