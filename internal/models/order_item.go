package models

type OrderItem struct {
  Order         Order       `json:"order"`
  OrderId       int64       `json:"order_id"`
  Product       Product     `json:"product"`
  ProductId     int64       `json:"product_id"`
  ProductType   ProductType `json:"product_type"`
  ProductTypeId int64       `json:"product_type_id"`
  Quantity      int64       `json:"quantity"`
  TotalPrice    float64     `json:"total_price"`
}
