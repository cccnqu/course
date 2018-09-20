package main

import (
	"os"
	"fmt"
	"./mt"
)

func main() {
	e := os.Args[1:]
	c := mt.E2C(e)
	fmt.Printf("%v\n", c)
	e = []string {"a", "cat"}
	c = mt.E2C(e) 
	fmt.Printf("%v\n", c)
}
