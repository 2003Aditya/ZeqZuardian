package analytics

import (
	"encoding/json"
	"net/http"
)

type Blocked struct {
	Key  string `json:"key"`
	From string `json:"from"`
	To   string `json:"to"`
}

func BlockedHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)

	queryParam := r.URL.Query()

	key := queryParam.Get("key")
	from := queryParam.Get("from")
	to := queryParam.Get("to")

	response := Blocked{
		Key:  key,
		From: from,
		To:   to,
	}

	json.NewEncoder(w).Encode(response)
}
