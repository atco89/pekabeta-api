package domain

type OrderStatus string

//noinspection ALL
const (
	APPROVED      OrderStatus = "approved"
	CANCELLED     OrderStatus = "cancelled"
	INVOICED      OrderStatus = "invoiced"
	ON_HOLD       OrderStatus = "on_hold"
	PACKED        OrderStatus = "packed"
	READY_TO_SHIP OrderStatus = "ready_to_ship"
)
