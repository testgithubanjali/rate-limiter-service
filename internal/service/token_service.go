package service

import (
	"rate-limiter-service/internal/model"
	"rate-limiter-service/internal/pkg/utils"
	"rate-limiter-service/store"
	"sync"
	"time"
)

type RateLimiterService struct {
	store store.store
}

func NewRateLimiterService(s store.store) *RateLimiterService {
	return &RateLimiterService{store: s}
}

func (rl *RateLimiterService) Allow(userID string) bool {
	bucket := rl.store.Get(userID)
	bucket.mu.Lock()
	defer bucket.mu.Unlock()
	now := time.Now()
	elapsed := now.Sub(bucket.lastRefillTime).Seconds()
	newTokens := int(elapsed * float64(bucket.refillRate))
	if newTokens > 0 {
		bucket.Tokens = utils.Min(bucket.maxTokens, bucket.tokens+newTokens)
		bucket.lastRefill = now
	}
	allowed := false
	if bucket.tokens > 0 {
		bucket.tokens--
		return true
	}
	if redisStore, ok := rl.store.(*store.RedisStore); ok {
		redisStore.Save(userID, bucket)
	}
	return allowed

}
