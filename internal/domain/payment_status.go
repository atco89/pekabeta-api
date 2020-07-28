package domain

type PaymentStatus string

const (
	CANCELED  PaymentStatus = "canceled"
	COMPLETED PaymentStatus = "completed"
	PENDING   PaymentStatus = "pending"
	REFUNDED  PaymentStatus = "refunded"
)
