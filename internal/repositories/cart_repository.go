package repositories

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (r *RepositoryPostgres) AddToCart(ctx context.Context, cartItems models.CartItem) (models.CartItem, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO cart (user_id, product_id, product_type_id, quantity) VALUES ($1, $2, $3, $4) RETURNING id, user_id, product_id, product_type_id, quantity`,
		cartItems.UserId,
		cartItems.ProductId,
		cartItems.ProductTypeId,
		cartItems.Quantity).Scan(&cartItems.Id, &cartItems.UserId, &cartItems.ProductId, &cartItems.ProductTypeId, &cartItems.Quantity)
	if err != nil {
		return models.CartItem{}, err
	}

	return cartItems, nil
}

func (r *RepositoryPostgres) GetCartItems(ctx context.Context, cartItems models.CartItem) ([]models.CartItem, error) {
	var product = models.Product
	var pt = models.ProductType
	err := r.Conn.QueryRow(
		ctx,
		`SELECT p.id,
		pt.id,
		ci.quantity
		FROM products AS p
		JOIN product_type 
		
		
		`

	)
}