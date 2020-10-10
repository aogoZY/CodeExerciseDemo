package main

import "fmt"

func ResevieStr(str string) {
	for i := len(str) - 1; i >= 0; i-- {
		fmt.Print(string(str[i]))
	}
	fmt.Println()

}

func main() {
	var str string
	fmt.Scanln(&str)
	//str = "abcd"
	ResevieStr(str)

}
