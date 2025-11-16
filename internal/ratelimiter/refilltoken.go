package ratelimiter

import "time"

func(b *RateLimiterBucket) ReFillToken() {
    b.Mutex.Lock()
    defer b.Mutex.Unlock()

    now := time.Now()

    elapsed := now.Sub(b.LastRefill).Seconds()
    tokenToAdd := elapsed * b.RefillRate

    b.Tokens += tokenToAdd

    if b.Tokens > b.Capacity {
        b.Tokens = b.Capacity

    }

    b.LastRefill = now




}
