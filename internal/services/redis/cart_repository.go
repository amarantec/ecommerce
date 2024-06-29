package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/amarantec/e-commerce/internal/models"
	"github.com/redis/go-redis/v9"
)

type IRedisRepository interface {
	AddToCart(ctx context.Context, cartItem models.CartItem) (models.CartItem, error)
	GetCartItems(ctx context.Context) ([]models.CartItem, error)
}

type RedisRepository struct {
	LocalStorage *redis.Client
}

func (r *RedisRepository) AddToCart(ctx context.Context, userId int64, cartItem models.CartItem) (models.CartItem, error) {
	cartKey := fmt.Sprintf("cart:%d", userId)
	sProductId := strconv.Itoa(int(cartItem.ProductId))
	sProductTypeId := strconv.Itoa(int(cartItem.ProductTypeId))

	err := r.LocalStorage.HSet(ctx, sProductId, sProductTypeId).Err()
	if err != nil {
		log.Fatalf("Could not add item to cart: %v", err)
	}

	err = r.LocalStorage.Expire(ctx, cartKey, 72*time.Hour).Err()
	if err != nil {
		log.Fatalf("Could not set expiration: %v", err)
	}

	return cartItem, nil
}

func (r *RedisRepository) GetCartItems(ctx context.Context, userId string) ([]models.CartItem, error) {
	cartKey := fmt.Sprintf("cart: %s", userId)

	cartItemsMap, err := r.LocalStorage.HGetAll(ctx, cartKey).Result()
	if err != nil {
		log.Fatalf("Could not get cart items: %v", err)
	}

	var cartItems []models.CartItem
	for productIdStr, productTypeIdStr := range cartItemsMap {
		productId, err := strconv.ParseInt(productIdStr, 10, 64)
		if err != nil {
			log.Printf("Invalid product ID for product %s: %v", productIdStr, err)
			continue
		}

		productTypeId, err := strconv.ParseInt(productTypeIdStr, 10, 64)
		if err != nil {
			log.Printf("Invalid product type ID for product %s: %v", productTypeIdStr, err)
		}
		cartItem := models.CartItem{
			ProductId:     productId,
			ProductTypeId: productTypeId,
		}
		cartItems = append(cartItems, cartItem)
	}
	return cartItems, nil
}
