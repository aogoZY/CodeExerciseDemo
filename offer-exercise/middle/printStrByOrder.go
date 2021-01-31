package main

import (
	"fmt"
)

func main() {
	sortString("A Famous Saying: Much Ado About Nothing (2012/8).")
}

func printStrByOrder(str string) (res string) {
	for i := 1; i < len(str); i++ {
		for j := i - 1; j > 0; j-- {
			if string(str[i]) < "z" && string(str[i]) > "a" || string(str[i]) < "Z" && string(str[i]) > "A" {
				if str[i] < str[j] {

				}

			}
		}
	}
	return
}


func sortString(s string) {
	str := []byte(s)
	n := len(str)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if isChar(str[j]) {
				next := j + 1
				for next < n && !isChar(str[next]) {
					next++
				}

				if next < n && isChar(str[next]) && toInt(str[j]) > toInt(str[next]) {
					str[j], str[next] = str[next], str[j]
				}
			}
		}
	}

	fmt.Println(string(str))
}

func isChar(char uint8) bool {
	return (char >= 'a' && char <= 'z') || char >= 'A' && char <= 'Z'
}

func toInt(char uint8) uint8 {
	if char >= 'A' && char <= 'Z' {
		return char + 32
	} else {
		return char
	}
}

