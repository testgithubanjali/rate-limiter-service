package store

import (
	"rate-limiter-service/internal/model"
)

type store interface {
	Get(userID string) model.TokenBucket
}
