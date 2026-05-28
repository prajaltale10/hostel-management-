package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	ips map[string][]time.Time
	mu  sync.Mutex
	rate int
	burst int
}

var limiter = &rateLimiter{
	ips:   make(map[string][]time.Time),
	rate:  100, // max 100 requests
	burst: 60,  // per 60 seconds
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter.mu.Lock()
		defer limiter.mu.Unlock()

		now := time.Now()
		validTimes := []time.Time{}

		if times, exists := limiter.ips[ip]; exists {
			for _, t := range times {
				if now.Sub(t).Seconds() < float64(limiter.burst) {
					validTimes = append(validTimes, t)
				}
			}
		}

		if len(validTimes) >= limiter.rate {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			return
		}

		validTimes = append(validTimes, now)
		limiter.ips[ip] = validTimes
		c.Next()
	}
}
