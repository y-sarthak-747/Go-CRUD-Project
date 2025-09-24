package config

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var (
	Ctx   = context.Background()
	Redis *redis.Client
)

func ConnectRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // docker mapped port
		Password: "",               // no password by default
		DB:       0,                // use default DB
	})

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("✅ Connected to Redis")
}
