package models

type OrderDetailsProductResponse struct {
  ProductId     int64     `json:"product_id"`
  Title         string    `json:"title"`
  ProductType   string    `json:"product_type"`
  ImageUrl      string    `json:"image_url"`
  Quantity      int64     `json:"quantity"`
  TotalPrice    float64   `json:"total_price"`
}
