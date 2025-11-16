package routes

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from the Go Application to all :6969  user ")

}
