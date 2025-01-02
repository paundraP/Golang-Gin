package pkg

import (
	"context"
	"log"

	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Allow(ctx context.Context, key string, limit redis_rate.Limit) (*redis_rate.Result, error)
}

type RedisLimiter struct {
	*redis_rate.Limiter
}

func SetupRedisLimiter(connStr string) Redis {
	client := redis.NewClient(&redis.Options{
		Addr: connStr,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	}

	return &RedisLimiter{Limiter: redis_rate.NewLimiter(client)}
}
