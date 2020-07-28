package domain

import "github.com/google/uuid"

type Order struct {
	Base
	CustomerID   uuid.UUID    `json:"customer_id" gorm:"column:customer_id;type:char(36);not null"`
	OrderStatus  OrderStatus  `json:"order_status" gorm:"column:order_status;type:varchar(255);not null"`
	ShippingType ShippingType `json:"shipping_type" gorm:"column:shipping_type;type:varchar(255);not null"`
}
