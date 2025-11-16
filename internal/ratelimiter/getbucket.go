package ratelimiter

import "time"

var manager = &RateLimiterManager{
    Bucket: make(map[string]*RateLimiterBucket),
}

// GetBucket func will return the bucket by using the Key
// if that key doesn't exist yet
// it will create a new bucket with that key
func(m *RateLimiterManager) GetBucket(key string) *RateLimiterBucket {

    m.Mutex.Lock()

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


     m.Mutex.Unlock()
    return bucket
}


func GetBucket(key string) *RateLimiterBucket {
    return manager.GetBucket(key)
}
