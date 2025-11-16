package wrapper

import "net/http"


func GET(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodGet {
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
        h(w, r)
    }
}

