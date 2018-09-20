package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	var z float64 = 1.0
  for diff:=1.0; math.Abs(diff) > 0.01; {
		fmt.Println("z=", z, "diff=", diff)
    diff = z*z-x
    z -= diff / (2*z)
  }
  return z
}

func main() {
	fmt.Println(Sqrt(2))
}