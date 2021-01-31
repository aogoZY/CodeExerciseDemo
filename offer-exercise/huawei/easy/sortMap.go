package main

import (
	"fmt"
	"sort"
)

//给定n个字符串，请对n个字符串按照字典序排列。

func main() {
	var num int
	var str string
	fmt.Scanln(&num)
	var strList []string
	for i := 0; i < num; i++ {
		fmt.Scanln(&str)
		strList = append(strList, str)
	}
	fmt.Println(strList)
	sort.Strings(strList)
	for _,item:=range strList{
		fmt.Println(item)
	}
}
