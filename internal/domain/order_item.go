package domain

import "github.com/google/uuid"

type OrderItem struct {
	Base
	OrderID   uuid.UUID `json:"order_id" gorm:"column:order_id;type:char(36);not null"`
	ProductID uuid.UUID `json:"product_id" gorm:"column:product_id;type:char(36);not null"`
	Quantity  float32   `json:"quantity" gorm:"column:quantity;type:decimal(10,2);not null"`
}
