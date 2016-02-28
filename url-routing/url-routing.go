package main

import (
  "fmt"
  "net/http"

  "github.com/julienschmidt/httprouter"
)

func main() {
  // Define out router
  r := httprouter.New()
  r.GET("/", homeHandler)

  // Post collection
  r.GET("/posts", postsIndexHandler)
  r.POST("/posts", postsCreateHandler)

  // Post Singular
  r.GET("/posts/:id", postShowHandler)
  r.PUT("/posts/:id", postUpdateHandler)
  r.GET("/posts/:id/edit", postEditHandler)

  fmt.Println("Starting server on :8080")
  http.ListenAndServe(":8080", r)
}


func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "Home")
}

func postsIndexHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "posts index")
}

func postsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "posts create")
}

func postShowHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    id := p.ByName("id")
    fmt.Fprintln(rw, "showing post", id)
}

func postUpdateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post update")
}

func postDeleteHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post delete")
}

func postEditHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "post edit")
}
