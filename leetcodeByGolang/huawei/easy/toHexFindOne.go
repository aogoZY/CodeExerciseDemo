package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scanln(&num)
	_, countNum := hex(num)
	fmt.Println(countNum)
}

func hex(num int) (res string, count int) {
	for i := num; i > 0; i /= 2 {
		lsa := i % 2
		if lsa == 1 {
			count++
		}
		lsaSr := strconv.Itoa(lsa)
		res = lsaSr + res
	}
	return res, count
}
