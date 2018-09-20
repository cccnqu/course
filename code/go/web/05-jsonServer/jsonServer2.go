package main

import (
  "encoding/json"
  "net/http"
)

type Profile struct {
  Name    string // 注意，Name 的第一個字母一定要大寫，否則無法 Marshal
  Tel     string
}

func main() {
  http.HandleFunc("/", foo)
  http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
  profile := Profile{"Alex", "331234"}

  js, err := json.Marshal(profile)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}