package main

import (
	"fmt"
	"strings"
)

//判断输入的字符是否可以相互抵消
//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//用栈，判断当前值若为）是否和栈顶元素相等，相等则消消乐，不想等代表不合法直接return false
func isValid(s string) bool {
	var stack []byte
	paren_map := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	if len(s) < 2 {
		return false
	}
	for i, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || paren_map[s[i]] != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

func main() {
	res := isValidByStack("([]()[]{})")
	//printInt([]int{3,2,1})
	fmt.Println(res)
}

func isValid2(s string) bool {
	for {
		old := s
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
		if s == "" {
			return true
		}
		if len(s) == len(old) {
			return false
		}
	}
	return false

}

func isValidByStack(s string) bool {
	var stack []byte
	converMap := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i, v := range s {
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, s[i])
		} else if len(stack) == 0 || converMap[s[i]] != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack)==0{
		return true
	}
	return false
}

func printInt(input []int) {
	new := input[:2]
	fmt.Println(new)
}
