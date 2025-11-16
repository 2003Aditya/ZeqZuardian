package analytics

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Message string `json:"message"`
	Route   string `json:"route"`
	Params  Param `json:"Param"`
}

type Param struct {
	Key  string `json:"key"`
	From string `json:"from"`
	To   string `json:"to"`
}

func RequestHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(http.StatusOK)

	queryParam := r.URL.Query()

	key := queryParam.Get("key")
	from := queryParam.Get("from")
	to := queryParam.Get("to")

    response := Request {
        Message: "Not implemented completely",
        Route: "/analytics/requests.go",
        Params: Param{
            Key: key,
            From: from,
            To: to,

        },

    }

    json.NewEncoder(w).Encode(response)

}
