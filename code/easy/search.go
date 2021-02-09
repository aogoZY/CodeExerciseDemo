package main

import (
	"fmt"
)

//二分查找 有序序列 中间值与目标值比较 不断缩小范围
func search(nums []int, target int) int {
	var left, right, mid int
	right = len(nums) - 1
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			return mid
		}
	}
	return -1

}

func main() {
	//input := []int{7,1,5,3,6,4}
	//res := maxProfit(input)
	input := []int{1, 2, 4, 5, 7}
	res := search(input, 4)
	fmt.Println(res)
}

func maxProfit(prices []int) int {
	var list []int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > 0 {
				list = append(list, prices[j]-prices[i])
			}
		}
	}
	fmt.Println(list)
	sortedList := sortList(list)
	return sortedList[len(list)-1]
}

func sortList(input []int) []int {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	fmt.Println(input)
	return input
}
