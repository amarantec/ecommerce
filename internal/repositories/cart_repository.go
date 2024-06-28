package repositories

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (r *RepositoryPostgres) AddToCart(ctx context.Context, cartItem models.CartItem) (models.CartItem, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO cart_items (product_id, product_type_id) VALUES ($1, $2) RETURNING product_id, product_type_id`, cartItem.ProductId, cartItem.ProductTypeId).Scan(&cartItem.ProductId, &cartItem.ProductTypeId)
	if err != nil {
		return models.CartItem{}, err
	}
	return cartItem, err
}

func (r *RepositoryPostgres) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT product_id, product_type_id FROM cart_items`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var cartItems []models.CartItem
	for rows.Next() {
		var cartItem models.CartItem
		if err := rows.Scan(
			&cartItem.ProductId,
			&cartItem.ProductTypeId); err != nil {
			return nil, err
		}
		cartItems = append(cartItems, cartItem)
	}
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
