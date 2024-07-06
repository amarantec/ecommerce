package models

type CartItem struct {
	ProductId     int64 `json:"product_id"`
	ProductTypeId int64 `json:"product_type_id"`
}
