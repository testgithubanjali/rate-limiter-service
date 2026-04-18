package structs

import (
	"sync"
)

type User struct {
	tokens         int
	maxTokens      int
	refillRate     int
	lastRefillTime int64
	mu             sync.Mutex
}
