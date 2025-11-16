package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiterEvent struct {
	APIKey     string    `json:"apikey"`
	TimeStamp  time.Time `json:"timestamp"`
	Allowed    bool      `json:"allowed"`
	TokensLeft float64    `json:"tokensleft"`
}

type RateLimiterBucket struct {
	Tokens     float64    `json:"tokens"`
	Capacity   float64    `json:"capacity"`
	LastRefill time.Time  `json:"lastrefill"`
	Mutex      sync.Mutex `json:"_"`
	RefillRate float64    `json:"refillrate"`
}

type RateLimiterManager struct {
    Bucket map[string]*RateLimiterBucket
    Mutex sync.Mutex

}
