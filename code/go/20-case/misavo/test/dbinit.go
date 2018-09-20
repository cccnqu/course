package main

import (
	"fmt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"log"
	// "time"
	"encoding/json"
	"io/ioutil"
	"gopkg.in/mgo.v2/bson"
	"../lib/misavo"
)

func main() {
	var shops []misavo.Shop
	jsonBuffer, err := ioutil.ReadFile("someShops.json")
	if err != nil {
		panic(err)
	}
	jsonStr := string(jsonBuffer)
	fmt.Println(jsonStr) // print the content as a 'string'
	json.Unmarshal([]byte(jsonStr), &shops)
	fmt.Println(shops)

	db, err := misavo.OpenDb()
	if err != nil {
		panic(err)
	}
	defer misavo.CloseDb()
	c := db.C("shop")
	c.RemoveAll(nil)
	for _, shop := range shops {
		shop.Id = misavo.Guid()
		shop.Fulltext = ""
		shopB, _ := json.Marshal(shop)
		shop.Fulltext = string(shopB)
		fmt.Printf("shop.FullText=%s\n", shop.Fulltext)
		err = c.Insert(&shop)
	}
	if err != nil {
		log.Fatal(err)
	}

	q := bson.M{"name":"茶舖子"}
	queryShops, err := misavo.QueryShop(q)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("queryShops=", queryShops)
	}
}
