package models

type CartProductResponse struct {
  ProductId     int64     `json:"product_id"`
  Title         string    `json:"title"`
  ProductTypeId int64     `json:"product_type_id"`
  ProductType   string    `json:"product_type"`
  ImageUrl      string    `json:"image_url"`
  Price         float64   `json:"price"`
  Quantity      int64     `json:"quantity"`
}
