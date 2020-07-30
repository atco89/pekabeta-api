package domain

import "github.com/google/uuid"

type Invoice struct {
	Base
	OrderId       uuid.UUID     `json:"order_id" gorm:"column:order_id;type:char(36);not null"`
	Amount        float32       `json:"amount" gorm:"column:amount;type:decimal(10,2);not null"`
	PaymentType   PaymentType   `json:"payment_type" gorm:"column:payment_type;type:varchar(15);not null"`
	PaymentStatus PaymentStatus `json:"payment_status" gorm:"column:payment_status;type:varchar(10);not null"`
}
