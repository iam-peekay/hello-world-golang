package main

import (
    "html/template"
    "net/http"
    "path"
)

type book struct {
    Title  string
    Author string
}

func main() {
    http.HandleFunc("/", showBooks)
    http.ListenAndServe(":8080", nil)
}

func showBooks(w http.ResponseWriter, r *http.Request) {
    b := book{"Building Web Apps with Go", "Jeremy Saenz"}

    fp := path.Join("templates", "index.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if err := tmpl.Execute(w, b); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
