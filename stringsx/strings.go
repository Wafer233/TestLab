package stringsx

import (
	"fmt"
	"strings"
)

func Field() {
	ss := strings.Fields("  foo bar  baz   ")
	fmt.Println("[]slice =", ss)
}

func Contains() {
	str := "abcdfeg"
	cons1 := strings.Contains(str, "a")
	cons2 := strings.Contains(str, "h")
	fmt.Println("a? ", cons1)
	fmt.Println("b? ", cons2)

}

func Joins() {
	s := []string{"foo", "bar", "baz"}
	prt := strings.Join(s, ", ")
	fmt.Println(prt)
}
