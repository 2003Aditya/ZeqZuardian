package ratelimiter

import (
	"time"

	"github.com/2003Aditya/zeqzuardian/internal/routes/debug"
)

func (b *RateLimiterBucket) Allow() bool {
	debug.Log("Allow() CALLED — before lock")

	b.Mutex.Lock()
	debug.Log("Allow() — after lock. Tokens=%v, LastRefill=%v", b.Tokens, b.LastRefill)
	defer func() {
		debug.Log("Allow() — unlocking. Tokens now=%v", b.Tokens)
		b.Mutex.Unlock()
	}()

	b.Refill()
	debug.Log("Allow() — after refill. Tokens=%v", b.Tokens)

	if b.Tokens >= 1 {
		b.Tokens -= 1
		debug.Log("Allow() — allowed. Tokens left=%v", b.Tokens)
		event := RateLimiterEvent{
			APIKey:     b.APIKey,
			TimeStamp:  time.Now(),
			Allowed:    true,
			TokensLeft: b.Tokens,
		}
		Events = append(Events, event)
		return true
	}

	debug.Log("Allow() — BLOCKED. Tokens=%v", b.Tokens)
	event := RateLimiterEvent{
		APIKey:     b.APIKey,
		TimeStamp:  time.Now(),
		Allowed:    false,
		TokensLeft: b.Tokens,
	}
	Events = append(Events, event)
	return false
}
