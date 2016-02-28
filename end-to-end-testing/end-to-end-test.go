package main

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func Test_App(t *testing.T) {
  ts := httptest.NewServer(app())
  defer ts.Close()

  res, err := http.Get(ts.URL)
  if err != nil {
    t.Fatal(err)
  }

  body, err := ioutil.ReadAll(res.Body)
  res.Body.Close()

  if err != nil {
    t.Fatal(err)
  }

  exp := "Before...hello world...After"

  if exp != string(body) {
    t.Fatalf("Expected %s got %s", exp, body)
  }
}
