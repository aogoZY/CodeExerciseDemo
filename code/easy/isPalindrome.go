package main

import "fmt"

func isPalindrome(s string) bool {
	var str string
	res := true
	if s == "" {
		return true
	}
	for i := 0; i < len(s); i++ {
		if Flag(s[i]) {
			str += string(IsNormal(s[i]))
		}
	}
	//fmt.Println(str)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-1-i] {
			res = false
			break
		}
		if i == len(str)-1-i {
			break
		}
	}
	return res
}

func main() {
	res := isPalindrome("0P")
	fmt.Println(res)
}

func IsNormal(s byte) (res byte) {
	if s >= '0' && s <= '9' {
		return s
	}
	if s >= 'a' && s <= 'z' {
		return s
	}
	if s >= 'A' && s <= 'Z' {
		return s + 32
	}
	return
}

func Flag(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	if s >= 'a' && s <= 'z' {
		return true
	}
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}
