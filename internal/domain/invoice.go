package domain

import "github.com/google/uuid"

type Invoice struct {
	OrderId       uuid.UUID     `bson:"order_id" json:"order_id"`
	Amount        float32       `bson:"amount" json:"amount"`
	PaymentStatus PaymentStatus `bson:"payment_status" json:"payment_status"`
	PaymentType   PaymentType   `bson:"payment_type" json:"payment_type"`
}
