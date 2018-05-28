package main

import (
  "net/http"
  "encoding/json"
  "log"
)

func pingHandler(writer http.ResponseWriter, req *http.Request) {
    js, err := json.Marshal("Pong!")
    if err != nil {
      http.Error(writer, err.Error(), http.StatusInternalServerError)
      return
    }

    writer.Header().Set("Content-Type", "application/json")
    writer.Write(js)
}

func main() {
  http.HandleFunc("/ping", pingHandler)

  log.Fatal(http.ListenAndServe(":8084", nil))
}
