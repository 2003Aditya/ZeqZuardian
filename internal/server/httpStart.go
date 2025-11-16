package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/2003Aditya/zeqzuardian/internal/routes"
)


func MainServer() {

    routes.Routes()
    Port := ":6969"
    fmt.Printf("Server Starting on port %s...\n", Port)

    log.Fatal(http.ListenAndServe(Port, nil))

}
