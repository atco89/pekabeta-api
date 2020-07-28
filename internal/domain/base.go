package domain

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type Base struct {
	ID        uuid.UUID  `gorm:"column:id;type:char(36);primary_key;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;not null;default:current_timestamp"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;not null;default:current_timestamp"`
}

func (b *Base) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("ID", uuid.New().String())
}
