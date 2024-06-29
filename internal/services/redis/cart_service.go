package redis

import (
	"context"
	"errors"

	"github.com/amarantec/e-commerce/internal/models"
)

type RedisService struct {
	CacheStorage RedisRepository
}

func (service RedisService) AddToCart(ctx context.Context, cartItem models.CartItem) (models.CartItem, error) {
	if cartItem.ProductId == 0 {
		return models.CartItem{}, errors.New("product id empty")
	}
	if cartItem.ProductTypeId == 0 {
		return models.CartItem{}, errors.New("product type id empty")
	}
	return service.AddToCart(ctx, cartItem)
}

func (service RedisService) GetCartItems(ctx context.Context) ([]models.CartItem, error) {
	cartItems, err := service.GetCartItems(ctx)
	if err != nil {
		return nil, err
	}
	return cartItems, nil
}
