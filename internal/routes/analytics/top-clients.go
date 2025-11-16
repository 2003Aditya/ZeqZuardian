package analytics

import (
	"encoding/json"
	"net/http"
)

type TopClient struct {
	Limit string `json:"limit"`
	From  string `json:"from"`
	To    string `json:"to"`
}

func TopClientHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)

    queryParam := r.URL.Query()

    limit := queryParam.Get("limit")
    from := queryParam.Get("from")
    to := queryParam.Get("to")

	response := TopClient{
        Limit: limit,
        From: from,
        To: to,
    }

    json.NewEncoder(w).Encode(response)
}
