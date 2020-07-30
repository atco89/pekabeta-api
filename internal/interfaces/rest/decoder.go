package rest

import (
	"github.com/google/uuid"
	"pekabeta/internal/domain"
	"pekabeta/internal/interfaces/rest/api"
)

func DecodeOrderItems(items []api.OrderItem) domain.OrderItems {
	var domainOrderItems domain.OrderItems
	for i := 0; i < len(items); i++ {
		domainOrderItems = append(domainOrderItems, domain.OrderItem{
			ProductID: uuid.MustParse(items[i].ProductId),
			Quantity:  items[i].Quantity,
		})
	}
	return domainOrderItems
}
