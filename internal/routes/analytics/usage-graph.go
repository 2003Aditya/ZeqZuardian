package analytics

import (
	"encoding/json"
	"net/http"
)

type UsageGraph struct {
	Interval string `json:"interval"`
	From     string `json:"from"`
	To       string `json:"to"`
}

func UsageGraphHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)

	queryParam := r.URL.Query()

	interval := queryParam.Get("interval")
	from := queryParam.Get("from")
	to := queryParam.Get("to")

	response := UsageGraph{
		Interval: interval,
		From:     from,
		To:       to,
	}

	json.NewEncoder(w).Encode(response)
}
