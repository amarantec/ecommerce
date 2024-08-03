package models

type OrderOverviewResponse struct {
  Id                int64       `json:"id"`
  OrderDate         time.Time   `json:"order_date"`
  TotalPrice        float64     `json:"total_price"`
  Product           string      `json:"product"`
  ProductImageUrl   string      `json:"product_image_url"`
}
