package main

import (
	"fmt"
)

//字符中所有出现的数字前后加上符号“*”，其他字符保持不变
//Jkdi234klowe90a3
//Jkdi*234*klowe*90*a*3*

func main() {
	var str string
	fmt.Scan(&str)
	res := pritnStr(str)
	fmt.Println(res)
}

func pritnStr(str string) (res string) {
	for i := 0; i < len(str)-1; i++ {
		a := str[i]
		fmt.Println(a)
		boolBefore := IsNum(str[i])
		boolAfter := IsNum(str[i+1])
		if !boolBefore && boolAfter {
			res += string(str[i]) + "*"
		} else if boolBefore && !boolAfter {
			res += string(str[i]) + "*"
		} else {
			res += string(str[i])
		}
	}
	if IsNum(str[len(str)-1]) {
		res += string(str[len(str)-1]) + "*"
	}
	return res

}

func IsNum(value byte) bool {
	fmt.Println(value)
	if value >= '0' && value <= '9' {
		return true
	}
	return false
}
