package repositories

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (r *RepositoryPostgres) AddToCart(ctx context.Context, cartItems models.CartItem) (models.CartItem, error) {
	err := r.Conn.QueryRow(
		ctx,
		`INSERT INTO cart (user_id, product_id, product_type_id, quantity) VALUES ($1, $2, $3, $4) RETURNING id`,
		cartItems.UserId,
		cartItems.ProductId,
		cartItems.ProductTypeId,
		cartItems.Quantity).Scan(&cartItems.Id)
	if err != nil {
		return models.CartItem{}, err
	}

	return cartItems, nil
}


func (r *RepositoryPostgres) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
	rows, err := r.Conn.Query(
		ctx,
		`SELECT product_id, product_type_id, quantity FROM cart;`)

  if err != nil {
    return nil, err
  }
  
  defer rows.Close()

  var cartItems []models.CartItem
  for rows.Next() {
    var cartItem models.CartItem
    if err := rows.Scan(
      &cartItem.ProductId,
      &cartItem.ProductTypeId,
      &cartItem.Quantity); err != nil {
        return nil, err
    }
    
    cartItems = append(cartItems, cartItem)
  }

  if err := rows.Err(); err != nil {
    return nil, err
  }
  return cartItems, nil
} 

func (r *RepositoryPostgres) UpdateCartItems(ctx context.Context, cartItem models.CartItem, id int64) error {
  _, err := r.Conn.Exec(
    ctx,
    `UPDATE carts SET product_id = $2, product_type_id = $3, quantity = $4 WHERE id = $1;`, id, cartItem.ProductId, cartItem.ProductTypeId, cartItem.Quantity)
  
  if err != nil {
    return err
  }
  return nil
}
