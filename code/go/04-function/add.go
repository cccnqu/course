package main

import "fmt"

func add(x, y int) int { // 也可以寫成 add := func (x, y int) int {
	return x + y
}


func main() {
	sub:= func(x, y float64) float64 {
		return x - y
	}
	
	fmt.Println(add(42, 13))
	fmt.Println(sub(42.0, 13.0))
}