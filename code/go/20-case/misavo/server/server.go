package main

import (
  "fmt"
  "encoding/json"
  "net/http"
  "gopkg.in/mgo.v2/bson"
  "../lib/misavo"
)

func main() {
  _, err := misavo.OpenDb()
	if err != nil {
		panic(err)
  }
	defer misavo.CloseDb()
  http.HandleFunc("/shop", shopQuery)
  http.ListenAndServe(":3000", nil)
}

func shopQuery(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
  // q := r.URL.Query()
	q := bson.M{"name":"茶舖子"}
	shops, err := misavo.QueryShop(q)
  // shops := []Shop { {"茶舖子", "323133", Location{24.45, 118.32} }, {"狗狗花園", "323152", Location{24.45, 118.33} } }
  js, err := json.Marshal(shops)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  fmt.Printf("js", js)
  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
}