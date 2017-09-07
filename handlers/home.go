package handlers

import (
    "net/http"
    "fmt"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprintf(w, "OK");
}
