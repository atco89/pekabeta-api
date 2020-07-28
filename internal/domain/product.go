package domain

type Product struct {
	Base
	Name            string         `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description     *string        `json:"description" gorm:"column:description;type:text"`
	ProductGroup    ProductGroup   `json:"product_group" gorm:"column:product_group;type:varchar(255);not null"`
	MeasurementUnit MeasurmentUnit `json:"measurment_unit" gorm:"column:measurment_unit;type:varchar(255);not null"`
	Price           float32        `json:"price" gorm:"column:price;type:decimal(10,2);not null"`
	ImagePath       *string        `json:"image_path" gorm:"column:image_path;type:text"`
	InStock         float32        `json:"in_stock" gorm:"column:in_stock;type:decimal(10,2);not null"`
}
