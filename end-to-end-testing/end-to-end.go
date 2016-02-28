package main

import (
    "fmt"
    "net/http"

    "github.com/codegangsta/negroni"
    "github.com/julienschmidt/httprouter"
)

func main() {
  http.ListenAndServe(":3000", app())
}

func app() http.Handler {
  n := negroni.Classic()

  m := func(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    fmt.Fprint(res, "Before...")
    next(res, req)
    fmt.Fprint(res, "...After")
    }

  n.Use(negroni.HandlerFunc(m))

  r := httprouter.New()

  r.GET("/", helloWorld)
  n.UseHandler(r)
  return n
}

func helloWorld(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
  fmt.Fprint(res, "hello world")
}
