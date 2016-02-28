package testing

import (
  "net/http"
  "net/http/httptest"
  "testing"
)

func Test_HelloWorld(t *testing.T) {
  req, err := http.NewRequest("GET", "http://localhost:8080", nil)
  if err != nil {
    t.Fatal(err)
  }

  res := httptest.NewRecorder()
  helloWorld(res, req)

  expect := "Hello Testing World!"
  actual := res.Body.String()
  if expect != actual {
    t.Fatalf("Expected %s but got %s", expect, actual)
  }
}
