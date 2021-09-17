package string

import (
	"fmt"
	"strings"
)

// [] () ([]

func main() {
	str := "((])"
	res := IsValid(str)
	fmt.Println(res)

}

func IsValid(str string) bool {
	for len(str) > 0 {
		old := str
		str = strings.ReplaceAll(str, "()", "")
		str = strings.ReplaceAll(str, "[]", "")
		if len(str)%2 != 0 {
			return false
		}
		if len(str) == 0 {
			return true
		}
		if old == str {
			return false
		}
	}
	return false
}
