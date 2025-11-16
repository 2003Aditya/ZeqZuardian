package ratelimiter



func (b *RateLimiterBucket) Allow() bool {

    b.Mutex.Lock()
    defer b.Mutex.Unlock()

    b.ReFillToken()

    if b.Tokens >= 1 {
        b.Tokens -= 1
        return true
    }

    return false

}
