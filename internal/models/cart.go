package models

type CartItem struct {
	Id            int64 `json:"id"`
	UserId        int64 `json:"user_id"`
	ProductId     int64 `json:"product_id"`
	ProductTypeId int64 `json:"product_type_id"`
	Quantity      int64 `json:"quantity"`
}
