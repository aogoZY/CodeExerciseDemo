package main

import (
	"fmt"
	"strings"
)

func main() {
	res := GetLowerDictStr("de", "df")
	fmt.Println(res)
}

func GetLowerDictStr(str1, str2 string) string {
	strA := str1 + str2
	strB := str2 + str1
	compareRes := strings.Compare(strA, strB)
	if compareRes == 0 {
		return strA
	}
	if compareRes == 1 {
		return strB
	} else {
		return strA
	}
}

