package domain

type Product struct {
	Name            string         `bson:"name" json:"name"`
	Description     *string        `bson:"description" json:"description"`
	ProductGroup    ProductGroup   `bson:"product_group" json:"product_group"`
	MeasurementUnit MeasurmentUnit `bson:"measurment_unit" json:"measurment_unit"`
	Price           float32        `bson:"price" json:"price"`
	ImagePath       *string        `bson:"image_path" json:"image_path"`
	InStock         float32        `bson:"in_stock" json:"in_stock"`
}
