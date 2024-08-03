package models

import "time"

type Order struct {
  Id          int64       `json:"id"`
  UserId      int64       `json:"user_id"`
  OrderDate   time.Now()  `json:"order_date"`
  TotalPrice  float64     `json:"total_price"`
  OrderItems  []OrderItem `json:"order_items"`
}
