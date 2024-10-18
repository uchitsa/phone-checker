package main

import (
  "log"
  "net/http"
)

func main() {
  http.HandleFunc("/ping", pingHandler)
  http.HandleFunc("/validate", validateHandler)
  http.HandleFunc("/shutdown", shutdownHandler)

  srv := &http.Server{
    Addr: "127.0.0.1:8888"
  }

  log.Println("Server started on 127.0.0.1:8888")
  log.Fatal(srv.ListenAndServe())
}