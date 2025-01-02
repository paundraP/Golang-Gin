package middleware

import (
	"fmt"
	"net/http"
	"rest-api-go/internal/pkg"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
)

type RateLimitMiddleware struct {
	RedisLimiter *pkg.RedisLimiter
}

func NewRateLimit(connStr string) *RateLimitMiddleware {
	return &RateLimitMiddleware{pkg.SetupRedisLimiter(connStr)}
}

const RateRequest = "rate_request_%s"

func (r *RateLimitMiddleware) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		res, _ := r.RedisLimiter.Allow(ctx, fmt.Sprintf(RateRequest, "userName"), redis_rate.Limit{
			Rate:   10,
			Burst:  10,
			Period: time.Minute,
		})
		if res.Allowed <= 0 {
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		c.Next()
	}
}
