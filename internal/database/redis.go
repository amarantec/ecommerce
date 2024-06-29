package database

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var LocalStorage *redis.Client

func InitRedis() (*redis.Client, error) {
	LocalStorage = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "091021",
		DB:       0,
	})

	_, err := LocalStorage.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return LocalStorage, nil
}
