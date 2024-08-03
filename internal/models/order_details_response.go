package models

type OrderDetailsResponse struct {
  OrderDate     time.Time     `json:"order_date"`
  TotalPrice    float64       `json:"total_price"`
  Products      []OrderDetailsProductResponse   `json:"products"`
}
