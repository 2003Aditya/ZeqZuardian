package ratelimiter

import "time"

func(b *RateLimiterBucket) Refill() {

    now := time.Now()

    elapsed := now.Sub(b.LastRefill).Seconds()
    tokenToAdd := elapsed * b.RefillRate

    b.Tokens += tokenToAdd

    if b.Tokens > b.Capacity {
        b.Tokens = b.Capacity

    }

    b.LastRefill = now




}
