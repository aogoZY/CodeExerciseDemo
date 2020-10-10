package main

import (
	"fmt"
)

func singleNumber(nums []int) []int {
	var res []int
	var testMap map[int]int
	testMap = make(map[int]int)
	for _, item := range nums {
		_, ok := testMap[item]
		if ok {
			testMap[item] += 1
		} else {
			testMap[item] = 1
		}

	}
	fmt.Println(testMap)
	for item := range testMap {
		if testMap[item]==1{
			res= append(res, item)
		}
	}
	return res
}

func main() {
	res := singleNumber([]int{1, 2, 1, 3, 2, 5})
	fmt.Println(res)
}
