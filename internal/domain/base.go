package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID  `json:"id" gorm:"column:id;type:char(36);primary_key;not null"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:current_timestamp"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:current_timestamp"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:datetime;null"`
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}
