package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Limited struct {
    Status string `json:"status"`
    HTTPCode int `json:"HttpCode"`
}

func LimitedHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from the Limited route")

    w.Header().Set("Content-type", "application/json")

    response := Limited {
        Status: "ok",
        HTTPCode: http.StatusAccepted,
    }

    json.NewEncoder(w).Encode(response)

}
