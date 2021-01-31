package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for i, v := range nums {
		for k := i + 1; k <= len(nums)-1; k++ {
			if nums[k] == target-v {
				result = append(result, i, k)
			}
		}
	}
	return  result
}

func main() {
	res := twoSum([]int{2, 7, 11, 15}, 26)
	fmt.Println(res)
}
