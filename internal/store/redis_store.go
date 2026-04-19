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
		ctx: context.Background(),
	}
}
func (r *RedisStore) Get(userID string) *model.TokenBucket{
	key := "rate_limit:" + userID
	data, err := r.client.Get(r.ctx, key) Result()
    //if not found create new bucket
	if err == redis.Nil{
		bucket := &model.TokenBucket{
			Tokens: 10,
			maxTokens: 10,
			RefillRate: 5,
			LastRefill: time.Now(),

		}
		r.save(key, bucket)
		return bucket
	}
	if err != nil {
		return &model.TokenBucket{
			Tokens:     10,
			MaxTokens:  10,
			RefillRate: 5,
			LastRefill: time.Now(),
		}
	}
	var bucket model.TokenBucket
	json.Unmarshal([]byte(data), &bucket)
	return &bucket
}

func (r *RedisStore) save(userID string, bucket *model.TokenBucket){
	key := "rate_limit:" + userID
	r.save(key, bucket)
}

func (r *RedisStore) save(key string, bucket *model.TokenBucket){
	bytes, _ := json.Marshal(bucket)

	r.client.Set(r.ctx, key, bytes, time.minute)
}