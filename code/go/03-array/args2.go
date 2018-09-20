package main
import "os"
import "fmt"
func main() {
	size := len(os.Args)
	for i := 0; i < size; i++ {
		fmt.Println(os.Args[i])
	}
}