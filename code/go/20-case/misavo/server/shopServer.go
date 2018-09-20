package main

import (
  "fmt"
  "encoding/json"
  "net/http"
)

type Location struct {
  Lat float64
  Lng float64
}

type Shop struct {
  Name string
  Tel  string
  At   Location
}

func main() {
  http.HandleFunc("/", shopQuery)
  http.ListenAndServe(":3000", nil)
}

func shopQuery(w http.ResponseWriter, r *http.Request) {
  shops := []Shop { {"茶舖子", "323133", Location{24.45, 118.32} }, {"狗狗花園", "323152", Location{24.45, 118.33} } }
  // shops := Shop{"茶舖子", "323133", Location{24.45, 118.32}}
  js, err := json.Marshal(shops)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Printf("js", js)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}