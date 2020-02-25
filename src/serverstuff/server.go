package serverstuff

import (
  "net/http"
  "io"
)

func StartServer() {
  http.HandleFunc("/name-list", NameListHandler)
  http.ListenAndServe(":9000", nil)
}

func NameListHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Work in progress\n")
}
