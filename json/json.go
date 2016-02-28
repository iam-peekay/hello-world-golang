package main

import (
  "encoding/json"
  "net/http"
)

type book struct {
  Title string `json:"title"`
  Author string `json:"author"`
}

func main() {
  http.HandleFunc("/", showBooks)
  http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
  b := book{"Build Web Apps With Go", "Jeremy Saenz"}

  js, err := json.Marshal(b)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}
