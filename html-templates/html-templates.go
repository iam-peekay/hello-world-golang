package main

import (
  "html/template"
  "net/http"
  "path"
)
type preethi struct {
  Name string
  Age int
  Description string
}

func main() {
  http.HandleFunc("/", showBooks)
  http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
  p := preethi{"Preethi Kasireddy", 25, "Passionate software engineer who loves building and learning new things <3"}

  fp := path.Join("templates", "index.html")
  tmpl, err := template.ParseFiles(fp)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  if err := tmpl.Execute(w, p); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}
