package services

import (
	"context"

	"github.com/amarantec/e-commerce/internal/models"
)

func (s Service) AddToCart(ctx context.Context, cartItem models.CartItem) (models.CartItem, error) {
	if cartItem.ProductId == 0 {
		return models.CartItem{}, ErrCartItemProductIdEmpty
	}

	if cartItem.ProductTypeId == 0 {
		return models.CartItem{}, ErrCartItemProductTypeIdEmpty
	}

	return s.Repository.AddToCart(ctx, cartItem)
}

func (s Service) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
	cartItems, err := s.Repository.GetCartItems(ctx)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
