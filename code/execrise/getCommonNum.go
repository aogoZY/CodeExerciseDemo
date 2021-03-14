package main

import (
	"fmt"
	sort2 "sort"
)

//[1,2,2,1,1]
//return 1

func main() {
	input := []int{1, 2, 2, 1, 1, 2, 2, 2, 33, 44}
	//orderInput := GetInputSort(input)
	//fmt.Println(orderInput)
	countMap := GetCountMap(input)
	fmt.Println(countMap)
	length := len(input) / 2
	//fmt.Println(length)
	res := CheckMapValue(countMap, length)
	fmt.Println(res)
}

func GetCountMap(input []int) map[int]int {
	countMap := make(map[int]int)
	for _, item := range input {
		_, ok := countMap[item]
		if ok {
			countMap[item] = countMap[item] + 1
		} else {
			countMap[item] = 1
		}
	}
	fmt.Println(countMap)
	return countMap
}

func CheckMapValue(countMap map[int]int, length int) int {
	for i, item := range countMap {
		if item > length {
			return i
		}
	}
	return -1
}

//将列表元素排序
func GetInputSort(input []int) []int {
	sort2.Ints(input)
	return input
}
