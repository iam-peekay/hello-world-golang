package main

import (
  "fmt"
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"
  "github.com/codegangsta/negroni"
)

func main() {
  // Middleware stack
  n := negroni.New(
    negroni.NewRecovery(),
    negroni.HandlerFunc(myMiddleware),
    negroni.NewLogger(),
    negroni.NewStatic(http.Dir("public")),
  )

  // Define out router
  r := httprouter.New()
  r.GET("/", homeHandler)

  n.UseHandler(r)

  n.Run(":8080")
}

func myMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
  log.Println("Logging on the way there")

  if r.URL.Query().Get("password") == "secret123" {
    next(rw, r)
  } else {
    http.Error(rw, "Not authorized", 401)
  }

  log.Println("Logging on the way back")
}

func homeHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "Home")
}
