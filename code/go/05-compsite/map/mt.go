package main
import "os"
import "fmt"

var e2c = map[string]string {
	"a"    : "一隻",
	"the"  : "這隻",
	"cat"  : "貓",
	"dog"  : "狗",
	"chase": "追",
	"eat"  : "吃",
}

func mt(e []string)[]string {
	size := len(e)
	c := make([]string, len(e))
	for i := 0; i < size; i++ {
		c[i] = e2c[e[i]]
		fmt.Println("e=", e[i], "c=", c[i])
	}
	return c
}

func main() {
	e := os.Args[1:] 
	c := mt(e)
	fmt.Printf("%v\n", c)
}
