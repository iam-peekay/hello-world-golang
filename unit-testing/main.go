package testing

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", helloWorld)
    http.ListenAndServe(":8080", nil)
}

func helloWorld(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "Hello Testing World!")
}
