package domain

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
)

type OrderItems []OrderItem

type OrderItem struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  float32   `json:"quantity"`
}

func (o *OrderItems) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), &o)
}

func (o OrderItems) Value() (driver.Value, error) {
	val, err := json.Marshal(o)
	return string(val), err
}
