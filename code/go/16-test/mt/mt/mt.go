package mt

var e2c map[string]string = map[string]string {
	"a"    : "一隻",
	"the"  : "這隻",
	"cat"  : "貓",
	"dog"  : "狗",
	"chase": "追",
	"eat"  : "吃",
}

var c2e map[string]string = map[string]string {}

func Reverse(e2c map[string]string) map[string]string {
	for e, c := range e2c {
		c2e[c] = e
	}
	return c2e
}

func init() {
	c2e = Reverse(e2c)
}

func MT(e []string, e2c map[string]string) []string {
	size := len(e)
	c := make([]string, len(e))
	for i := 0; i < size; i++ {
		c[i] = e2c[e[i]]
		// fmt.Println("e=", e[i], "c=", c[i])
	}
	return c
}

func E2C(e []string) []string {
	return MT(e, e2c)
}

func C2E(c []string) []string {
	return MT(c, c2e)
}