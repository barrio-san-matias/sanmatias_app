package handler

import (
    "fmt"
    "net/http"
    "time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("incoming request...")
    currentTime := time.Now().Format(time.RFC850)
    fmt.Fprintf(w, currentTime)
}
