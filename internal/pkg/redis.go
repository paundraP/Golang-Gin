package pkg

import (
	"context"
	"log"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type RedisLimiter struct {
	*redis_rate.Limiter
}

func SetupRedisLimiter(connStr string) *RedisLimiter {
	client := redis.NewClient(&redis.Options{
		Addr: connStr,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	}

	return &RedisLimiter{redis_rate.NewLimiter(client)}
}
