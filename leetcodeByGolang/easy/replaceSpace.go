package main

import "fmt"

func replaceSpace(s string) string {
	var result string
	for _, item := range s {
		if string(item) == " " {
			result += "%20"
		} else {
			result += string(item)
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	s := replaceSpace("we are happy")
	fmt.Println(s)
}
