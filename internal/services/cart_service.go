package services

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (s Service) AddToCart(ctx context.Context, cartItems models.CartItem) (models.CartItem, error) {
	if cartItems.ProductId == 0 {
		return models.CartItem{}, ErrCartProductIdEmpty
	}

	if cartItems.ProductTypeId == 0 {
		return models.CartItem{}, ErrCartProductTypeIdEmpty
	}

	if cartItems.Quantity == 0 {
		return models.CartItem{}, ErrCartQuantityEmpty
	}

	return s.Repository.AddToCart(ctx, cartItems)
}

func (s Service) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
  cartItems, err := s.Repository.GetCartItems(ctx)
  if err != nil {
    return nil, err
  }

  return cartItems, nil
}

func (s Service) UpdateCartItems(ctx context.Context, cartItem models.CartItem) error {
  err := s.Repository.UpdateCartItems(ctx, cartItem)
  if err != nil {
    return err
  }
  return nil
}
