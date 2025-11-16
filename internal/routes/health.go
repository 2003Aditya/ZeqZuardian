package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Health struct {
	Status      string `json:"status"`
	UpTime      string `json:"uptime"`
	CurrentTime string `json:"currenttime"`
}

var StartTime = time.Now()

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello From the Health Route")

    w.Header().Set("Content-Type", "application/json")

    response := Health {
        Status: "ok",
        UpTime: time.Since(StartTime).String(),
        CurrentTime: time.Now().Format(time.RFC3339),
    }

    json.NewEncoder(w).Encode(response)

}
