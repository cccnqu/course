package main

import "fmt"
import "math/cmplx"

func Cbrt(x complex128) complex128 {
  z := 1+0i
  for i:=0; i < 10; i++ {
	z = z - (cmplx.Pow(z, 3+0i) - x)/(3*cmplx.Pow(z, 2+0i))
  }
  return z
}

func main() {
	fmt.Println(Cbrt(2))
}
