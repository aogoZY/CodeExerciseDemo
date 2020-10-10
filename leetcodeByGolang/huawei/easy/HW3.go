package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int
	fmt.Scan(&num)
	var str string
	var strList []string
	for i := 0; i < 2; i++ {
		fmt.Scan(&str)
		strList = append(strList, str)
	}
	//fmt.Println(strList)
	_ = GetOrNum(strList)
	fmt.Println(1)
}

func GetOrNum(strList []string) (res int) {
	var count int
	for i := 0; i < len(strList[0]); i++ {
		x,_ := strconv.Atoi(string(strList[0][i]))
		y,_ := strconv.Atoi(string(strList[1][i]))
		if x == 1 && y == 1 {
			count++
		}
	}
	return  len(strList[0])-count
}
