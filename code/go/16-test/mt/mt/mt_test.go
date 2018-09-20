package mt

import (
	"testing"
	"strings"
	"reflect"
	"fmt"
)

var QA = map[string] string{
	"": "",
	"a cat": "一隻 貓",
	"a cat chase a dog":  "一隻 貓 追 一隻 狗",
}

/*
func TestACat(t *testing.T) {
	if !reflect.DeepEqual(MT([]string{"a", "cat"}), []string{"一隻", "貓"}) {
		t.Error(`MT(["a", "cat"]) != ["一隻", "貓"]`)
	}
}
*/
func TestMtACat(t *testing.T) {
	// e := []string{"a", "cat"}
	c := []string{"一隻", "貓"}
	e := strings.Split("a cat", " ")
	if !reflect.DeepEqual(E2C(e), c) {
		t.Error(`MT({{e}}) != {{c}}`)
	}
}

func TestQA(t *testing.T) {
	for e, c := range QA {
		ec := strings.Join(E2C(strings.Split(e, " ")), " ")
		// fmt.Printf("MT(%s) => %s == %s\n", e, ec, c)
		if !reflect.DeepEqual(ec, c) {
			t.Error("MT(e) != c")
		}
	}
}

func TestEE(t *testing.T) {
	for e, _ := range QA {
		e1 := strings.Split(e, " ")
		c1 := E2C(e1)
		e2 := C2E(c1)
		fmt.Printf("e1=%v c1=%v e2=%v\n", e1, c1, e2)
		if !reflect.DeepEqual(e1, e2) {
			t.Error("e1 != e2")
		}
	}
}