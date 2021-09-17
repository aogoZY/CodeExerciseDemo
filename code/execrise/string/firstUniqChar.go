package main

import "fmt"

func main() {
	res := firstUniqChar("leetcode")
	fmt.Println(string(res))
}

func firstUniqChar(s string) byte {
	list := make([]int, 26)
	for _, v := range s {
		list[v-'a'] += 1
	}
	fmt.Println(list)
	for _, value := range s {
		if list[value-'a'] == 1 {
			return byte(value)
		}
	}
	return ' '
}
