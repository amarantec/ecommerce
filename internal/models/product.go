package models

type Product struct {
	ID          int64            `json:"id"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ImageURL    string           `json:"image_url"`
	Category    Category         `json:"category"`
	CategoryId  int64            `json:"category_id"`
	Featured    bool             `json:"featured"`
	Variants    ProductVariant   `json:"variants"`
}
