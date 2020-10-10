package main

import (
	"fmt"
	"strings"
)

//思路：先遍历最短的子串，看子串是否在长子串中
func main() {
	strA := "abcdefghijklmnop"
	strB := "bcsafjklmnopqrstuvw"
	res := printFirstLongSubstr(strA, strB)
	fmt.Println(res)
}

func printFirstLongSubstr(strA, strB string) (res string) {
	if len(strA) > len(strB) {
		strA, strB = strB, strA
	}
	length := 0
	for i := 0; i < len(strA); i++ {
		for j := i + 1; j < len(strA); j++ {
			subStr := strA[i:j]
			if strings.Contains(strB, subStr) && (j-i) > length {
				res = subStr
				length = j - i
			}
		}
	}
	return res
}
