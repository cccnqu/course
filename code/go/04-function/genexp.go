// === BNF Grammar =====
// E = T [+-*/] E | T
// T = [0-9] | (E)

package main
import "fmt"
import "math/rand"
import "time"

// 用法:randInt(3,7) 會傳回 3,4,5,6,7 其中之一
func randInt(a, b int)int { // 隨機傳回一個介於 (a,b) 間的整數 (包含 a, b)
	return rand.Intn(b-a) + a
}

func randChoose(a []string)string {
	var len = len(a)
	var i = randInt(0, len-1)
	return a[i]
}

func E() {
	if randInt(1,10) <= 5 {
		T()
		fmt.Printf(randChoose([]string {"+", "-", "*", "/"}))
		E()
	} else {
		T()
	}
}

func T() {
	if (randInt(1,10) < 7) {
		fmt.Printf(randChoose([]string {"x", "y", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}))
	} else {
		fmt.Printf("(")
		E()
		fmt.Printf(")")
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	for i:=0; i<10; i++ {
		E()
		fmt.Printf("\n")
	}
}
