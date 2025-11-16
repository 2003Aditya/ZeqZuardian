package routes

import (
	"net/http"

	"github.com/2003Aditya/zeqzuardian/internal/routes/analytics"
)


func Routes() {

    http.HandleFunc("/", Home)
    http.HandleFunc("/health", HealthHandler)
    http.HandleFunc("/limited", LimitedHandler)
    http.HandleFunc("/analytics/requests", analytics.RequestHandler)
    http.HandleFunc("/analytics/top-clients", analytics.TopClientHandler)
    http.HandleFunc("/analytics/blocked", analytics.BlockedHandler)
    http.HandleFunc("/analytics/user-graph", analytics.UsageGraphHandler)

}
