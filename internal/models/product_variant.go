package models

type ProductVariant struct {
	ProductId     int64       `json:"product_id"`
	ProductType   ProductType `json:"product_type"`
	ProductTypeId int64       `json:"product_type_id"`
	Price         float64     `json:"price"`
	OriginalPrice float64     `json:"original_price"`
}
