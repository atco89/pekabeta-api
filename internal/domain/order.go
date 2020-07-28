package domain

import "github.com/google/uuid"

type Order struct {
	CustomerID   uuid.UUID    `bson:"customer_id" json:"customer_id"`
	OrderItems   []OrderItem  `bson:"order_items" json:"order_items"`
	OrderStatus  OrderStatus  `bson:"order_status" json:"order_status"`
	ShippingType ShippingType `bson:"shipping_type" json:"shipping_type"`
}
