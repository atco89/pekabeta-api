package database

import (
	"github.com/jinzhu/gorm"
	"pekabeta/internal/domain"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(
		domain.Customer{},
		domain.Product{},
		domain.Order{},
		domain.OrderItem{},
		domain.Invoice{},
	)
}
