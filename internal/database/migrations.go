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
		domain.Invoice{},
	)

	db.Model(domain.Customer{}).
		AddUniqueIndex("uq_customers_email", "email").
		AddUniqueIndex("uq_customers_phone_number", "phone_number")

	db.Model(domain.Order{}).
		AddForeignKey("customer_id", "customers(id)", "NO ACTION", "CASCADE")

	db.Model(domain.Invoice{}).
		AddForeignKey("order_id", "orders(id)", "NO ACTION", "CASCADE")
}
