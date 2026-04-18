package store

import (
	"rate-limiter-service/internal/model"
	"sync"
)

type store interface {
	Get(userID string) model.tokenBucket
}
