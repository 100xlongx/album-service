package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", Hello)
    fmt.Println("Server is running on port 8080...")
    http.ListenAndServe(":8080", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, Golang!")
}
