package ratelimiter

import "time"

// GetBucket func will return the bucket by using the Key
// if that key doesn't exist yet
// it will create a new bucket with that key
func(m *RateLimiterManager) GetBucket(key string) *RateLimiterBucket {

    m.Mutex.Lock()
    defer m.Mutex.Unlock()

    bucket, exists := m.Bucket[key]
    if !exists {
        bucket = &RateLimiterBucket{
            Tokens: 1,
            Capacity: 10,
            RefillRate: 1,
            LastRefill: time.Now(),
        }

        m.Bucket[key] = bucket

    }


    return bucket
}
