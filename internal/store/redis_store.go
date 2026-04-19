package store

import (
	"rate-limiter-service/internal/model"
	"sync"
)
func NewRedisStore()  *RedisStore{
	rdb := redis.NewClient(&redisOptions{
		Addr: "localhost: 6379",
	}
	return &RedisStore{
		client: rdb,
		ctx: context.Background()
	}
}