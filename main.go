package main

import (
    "fmt"
    "net/http"
    "os"
)

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
}

func main() {
    port := os.Getenv("PORT")
    if len(port) <= 0 {
        port = "8888"
    }
    http.HandleFunc("/", hello)
    http.ListenAndServe(":"+port, nil)
}
