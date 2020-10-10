package main

import (
	"fmt"
	"strconv"
)

func printNumbers(n int) []int {
	var big string
	var res []int
	for i := 0; i < n; i++ {
		big = big + "9"
	}
	fmt.Println(big)
	intNumI,_ :=strconv.Atoi(big)
	for i:=1;i<=intNumI;i++{
		res= append(res, i)
	}
	return res
}

func main() {
	res := printNumbers(2)
	fmt.Println(res)
}
