package domain

import "github.com/google/uuid"

type OrderItem struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  float32   `json:"quantity"`
}
