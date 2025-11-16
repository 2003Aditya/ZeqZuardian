package routes

import (
	"encoding/json"
	"net/http"

	"github.com/2003Aditya/zeqzuardian/internal/ratelimiter"
	"github.com/2003Aditya/zeqzuardian/internal/routes/debug"
)

type Limited struct {
    Status string `json:"status"`
    HTTPCode int `json:"HttpCode"`
}

func LimitedHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")
    debug.Log("LimitedHandler CALLED: %s", r.URL.String())

    key := r.URL.Query().Get("key")
    debug.Log("Key=%s", key)

    bucket := ratelimiter.GetBucket(key)
    debug.Log("Got bucket: %+v", bucket)

    allowed := bucket.Allow()
    debug.Log("Allow() returned: %v", allowed)


    if allowed {
        debug.Log("Responding ALLOWED")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "status": "allowed",
            "tokenleft": bucket.Tokens,
        })
        return
    }

    debug.Log("Responding BLOCKED")
    w.WriteHeader(http.StatusTooManyRequests)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "blocked",
        "tokenleft": bucket.Tokens,
    })
}

